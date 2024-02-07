//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package cache

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/common"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/hnsw/distancer"
)

// A PaginatedCache is a cache that stores vectors in pages of a fixed size.
// Pages are created on demand and stored in a slice.
// Each index in the slice can be either nil or point to a page.
// The combination of slice and pages represent a virtual one-dimensional array of vectors.
// Each vector is identified by its index in the virtual array.
//
// # Access and locking
//
// Reading and writing a single vector to the cache is a O(1) operation.
// The vector id is used to calculate the page and index.
// The cache uses a sharded lock to allow concurrent reads and writes,
// however it is not possible to use the sharded lock to lock a single page.
// This means that creating a page or increasing the size of the slice
// requires a complete lock on the cache. This is not a problem in practice,
// because these are relatively rare operations.
//
// # Memory usage and pruning
//
// When the cache slice is grown, pages are not allocated immediately.
// Instead, pages are allocated when they are needed.
// A pruning operation is run periodically to remove empty pages.
type paginatedCache[T float32 | byte | uint64] struct {
	shardedLocks     *common.ShardedRWLocks
	cache            [][][]T
	vectorForID      common.VectorForID[T]
	maxSize          int64
	count            int64
	cancel           chan bool
	logger           logrus.FieldLogger
	deletionInterval time.Duration
	pruneInterval    time.Duration

	// The maintenanceLock makes sure that only one maintenance operation, such
	// as growing the cache or clearing the cache happens at the same time.
	maintenanceLock sync.RWMutex

	// The pageLock makes sure that only one goroutine can create a page at the
	// same time.
	pageLock sync.Mutex
}

const (
	// Size of a page in the cache
	PageSize = 1024
	// Number of pages to grow the cache by
	GrowthDelta          = 6 * PageSize
	DefaultPruneInterval = 5 * time.Minute
)

func NewPaginatedFloat32Cache(vecForID common.VectorForID[float32], maxSize int,
	logger logrus.FieldLogger, normalizeOnRead bool, deletionInterval time.Duration,
) Cache[float32] {
	vc := &paginatedCache[float32]{
		vectorForID: func(ctx context.Context, id uint64) ([]float32, error) {
			vec, err := vecForID(ctx, id)
			if err != nil {
				return nil, err
			}
			if normalizeOnRead {
				vec = distancer.Normalize(vec)
			}
			return vec, nil
		},
		cache:            make([][][]float32, 1),
		count:            0,
		maxSize:          int64(maxSize),
		cancel:           make(chan bool),
		logger:           logger,
		shardedLocks:     common.NewDefaultShardedRWLocks(),
		maintenanceLock:  sync.RWMutex{},
		deletionInterval: deletionInterval,
		pruneInterval:    DefaultPruneInterval,
	}

	vc.cache[0] = make([][]float32, PageSize)

	vc.watchForDeletion()
	vc.watchForPrune()
	return vc
}

func NewPaginatedByteCache(vecForID common.VectorForID[byte], maxSize int,
	logger logrus.FieldLogger, deletionInterval time.Duration,
) Cache[byte] {
	vc := &paginatedCache[byte]{
		vectorForID:      vecForID,
		cache:            make([][][]byte, 1),
		count:            0,
		maxSize:          int64(maxSize),
		cancel:           make(chan bool),
		logger:           logger,
		shardedLocks:     common.NewDefaultShardedRWLocks(),
		maintenanceLock:  sync.RWMutex{},
		deletionInterval: deletionInterval,
		pruneInterval:    DefaultPruneInterval,
	}

	vc.cache[0] = make([][]byte, PageSize)

	vc.watchForDeletion()
	vc.watchForPrune()
	return vc
}

func NewPaginatedUInt64Cache(vecForID common.VectorForID[uint64], maxSize int,
	logger logrus.FieldLogger, deletionInterval time.Duration,
) Cache[uint64] {
	vc := &paginatedCache[uint64]{
		vectorForID:      vecForID,
		cache:            make([][][]uint64, 1),
		count:            0,
		maxSize:          int64(maxSize),
		cancel:           make(chan bool),
		logger:           logger,
		shardedLocks:     common.NewDefaultShardedRWLocks(),
		maintenanceLock:  sync.RWMutex{},
		deletionInterval: deletionInterval,
		pruneInterval:    DefaultPruneInterval,
	}

	vc.cache = make([][][]uint64, 1)
	vc.cache[0] = make([][]uint64, PageSize)

	vc.watchForDeletion()
	vc.watchForPrune()
	return vc
}

func (s *paginatedCache[T]) All() [][]T {
	s.shardedLocks.RLockAll()
	defer s.shardedLocks.RUnlockAll()

	out := make([][]T, 0, s.Len())
	for _, page := range s.cache {
		for _, vec := range page {
			if vec != nil {
				out = append(out, vec)
			}
		}
	}

	return out
}

// getPageForID returns the page and index of the vector for the given id.
// This function assumes that the caller holds a read lock on the page.
func (s *paginatedCache[T]) getPageForID(id uint64) (page [][]T, idx int) {
	pageIdx := id / PageSize
	page = s.cache[pageIdx]
	if page == nil {
		return nil, -1
	}

	return page, int(id % PageSize)
}

// upsert gets or creates a page and applies an update to it.
// No locks must be held when calling this function.
func (s *paginatedCache[T]) upsert(id uint64, update func(page [][]T, idx int) error) error {
	// fast path
	s.shardedLocks.Lock(id)

	page, idx := s.getPageForID(id)
	if page != nil {
		err := update(page, idx)
		s.shardedLocks.Unlock(id)
		return err
	}
	s.shardedLocks.Unlock(id)

	s.pageLock.Lock()
	defer s.pageLock.Unlock()

	// try again in case the page was created in the meantime.
	// this is to avoid the waiting routines to all acquire a LockAll
	// lock when multiple goroutines try to create the same page.
	s.shardedLocks.Lock(id)
	page, idx = s.getPageForID(id)
	if page != nil {
		err := update(page, idx)
		s.shardedLocks.Unlock(id)
		return err
	}
	s.shardedLocks.Unlock(id)

	// try again, this time with all locks held
	s.shardedLocks.LockAll()
	page, idx = s.getPageForID(id)
	if page != nil {
		err := update(page, idx)
		s.shardedLocks.UnlockAll()
		return err
	}

	// page doesn't exist yet, create it
	pageIdx := id / PageSize
	page = make([][]T, PageSize)
	s.cache[pageIdx] = page
	idx = int(id % PageSize)
	err := update(page, idx)

	// opportunistically create the next page
	if len(s.cache) > int(pageIdx)+1 && s.cache[pageIdx+1] == nil {
		page = make([][]T, PageSize)
		s.cache[pageIdx+1] = page
	}

	s.shardedLocks.UnlockAll()
	return err
}

func (s *paginatedCache[T]) Get(ctx context.Context, id uint64) ([]T, error) {
	var vec []T

	s.shardedLocks.RLock(id)
	page, idx := s.getPageForID(id)
	if page != nil {
		vec = page[idx]
	}
	s.shardedLocks.RUnlock(id)

	if vec != nil {
		return vec, nil
	}

	return s.handleCacheMiss(ctx, id)
}

func (s *paginatedCache[T]) Delete(ctx context.Context, id uint64) {
	s.shardedLocks.Lock(id)
	defer s.shardedLocks.Unlock(id)

	if int(id) >= len(s.cache)*PageSize {
		return
	}

	page, idx := s.getPageForID(id)
	if page == nil || page[idx] == nil {
		return
	}

	page[idx] = nil
	atomic.AddInt64(&s.count, -1)

	// check if page is empty
	for _, v := range page {
		if v != nil {
			return
		}
	}

	// page is empty, delete it
	s.cache[id/PageSize] = nil
}

func (s *paginatedCache[T]) handleCacheMiss(ctx context.Context, id uint64) ([]T, error) {
	vec, err := s.vectorForID(ctx, id)
	if err != nil {
		return nil, err
	}

	atomic.AddInt64(&s.count, 1)

	err = s.upsert(id, func(page [][]T, idx int) error {
		page[idx] = vec
		return nil
	})

	return vec, err
}

func (s *paginatedCache[T]) MultiGet(ctx context.Context, ids []uint64) ([][]T, []error) {
	out := make([][]T, len(ids))
	errs := make([]error, len(ids))

	for i, id := range ids {
		out[i], errs[i] = s.Get(ctx, id)
	}

	return out, errs
}

var prefetchFunc func(in uintptr) = func(in uintptr) {
	// do nothing on default arch
	// this function will be overridden for amd64
}

func (s *paginatedCache[T]) Prefetch(id uint64) {
	s.shardedLocks.RLock(id)
	page, idx := s.getPageForID(id)
	if page != nil {
		prefetchFunc(uintptr(unsafe.Pointer(&page[idx])))
	}
	s.shardedLocks.RUnlock(id)
}

func (s *paginatedCache[T]) Preload(id uint64, vec []T) {
	s.shardedLocks.Lock(id)

	page, idx := s.getPageForID(id)
	if page == nil {
		s.shardedLocks.Unlock(id)
		return
	}

	page[idx] = vec
	s.shardedLocks.Unlock(id)

	atomic.AddInt64(&s.count, 1)

	s.upsert(id, func(page [][]T, idx int) error {
		page[idx] = vec
		return nil
	})
}

func (s *paginatedCache[T]) Grow(node uint64) {
	s.maintenanceLock.RLock()
	if node < uint64(len(s.cache)*PageSize) {
		s.maintenanceLock.RUnlock()
		return
	}
	s.maintenanceLock.RUnlock()

	s.maintenanceLock.Lock()
	defer s.maintenanceLock.Unlock()

	// make sure cache still needs growing
	// (it could have grown while waiting for maintenance lock)
	if node < uint64(len(s.cache)*PageSize) {
		return
	}

	s.shardedLocks.LockAll()
	defer s.shardedLocks.UnlockAll()

	pages := int((node + GrowthDelta) / PageSize)
	newCache := make([][][]T, pages)
	copy(newCache, s.cache)
	s.cache = newCache
}

func (s *paginatedCache[T]) Len() int32 {
	s.maintenanceLock.RLock()
	defer s.maintenanceLock.RUnlock()

	return int32(len(s.cache) * PageSize)
}

func (s *paginatedCache[T]) CountVectors() int64 {
	return atomic.LoadInt64(&s.count)
}

func (s *paginatedCache[T]) Drop() {
	s.deleteAllVectors()
	close(s.cancel)
}

func (s *paginatedCache[T]) deleteAllVectors() {
	s.shardedLocks.LockAll()
	defer s.shardedLocks.UnlockAll()

	s.logger.WithField("action", "hnsw_delete_vector_cache").
		Debug("deleting full vector cache")
	for i := range s.cache {
		s.cache[i] = nil
	}

	atomic.StoreInt64(&s.count, 0)
}

func (s *paginatedCache[T]) watchForDeletion() {
	if s.deletionInterval != 0 {
		go func() {
			t := time.NewTicker(s.deletionInterval)
			defer t.Stop()
			for {
				select {
				case <-s.cancel:
					return
				case <-t.C:
					s.replaceIfFull()
				}
			}
		}()
	}
}

func (s *paginatedCache[T]) watchForPrune() {
	if s.pruneInterval != 0 {
		go func() {
			t := time.NewTicker(s.pruneInterval)
			defer t.Stop()
			for {
				select {
				case <-s.cancel:
					return
				case <-t.C:
					s.prune()
				}
			}
		}()
	}
}

func (s *paginatedCache[T]) replaceIfFull() {
	if atomic.LoadInt64(&s.count) >= atomic.LoadInt64(&s.maxSize) {
		s.deleteAllVectors()
	}
}

func (s *paginatedCache[T]) prune() {
	var total, deleted, used, unallocated int

	start := time.Now()
	defer func() {
		s.logger.WithField("action", "hnsw_prune_vector_cache").
			WithField("took", time.Since(start)).
			WithField("deleted_pages", deleted).
			WithField("used_pages", used).
			WithField("unallocated_pages", unallocated).
			WithField("total_pages", total).
			Debugf("pruned vector cache, took %s", time.Since(start))
	}()

	var i int

	for {
		s.maintenanceLock.RLock()
		total = len(s.cache)
		s.maintenanceLock.RUnlock()

		if i >= total {
			return
		}

		// any lock will do, we just need to make sure that the cache doesn't
		// change while we're checking the page.
		s.shardedLocks.RLock(uint64(i))
		page := s.cache[i]
		if page == nil {
			s.shardedLocks.RUnlock(uint64(i))
			i++
			unallocated++
			continue
		}
		s.shardedLocks.RUnlock(uint64(i))

		s.shardedLocks.LockAll()
		page = s.cache[i]

		empty := true
		for j := range page {
			if page[j] != nil {
				empty = false
				break
			}
		}

		if !empty {
			s.shardedLocks.UnlockAll()
			used++
			i++
			continue
		}

		s.cache[i] = nil
		s.shardedLocks.UnlockAll()
		deleted++
	}
}

func (s *paginatedCache[T]) UpdateMaxSize(size int64) {
	atomic.StoreInt64(&s.maxSize, size)
}

func (s *paginatedCache[T]) CopyMaxSize() int64 {
	sizeCopy := atomic.LoadInt64(&s.maxSize)
	return sizeCopy
}

// noopCache can be helpful in debugging situations, where we want to
// explicitly pass through each vectorForID call to the underlying vectorForID
// function without caching in between.
type noopCache struct {
	vectorForID common.VectorForID[float32]
}

func NewNoopCache(vecForID common.VectorForID[float32], maxSize int,
	logger logrus.FieldLogger,
) *noopCache {
	return &noopCache{vectorForID: vecForID}
}
