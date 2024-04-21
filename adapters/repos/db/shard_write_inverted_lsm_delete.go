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

package db

import (
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/weaviate/weaviate/adapters/repos/db/inverted"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv"

	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
)

func (s *Shard) deleteFromInvertedIndicesLSM(props []inverted.Property, nilProps []inverted.NilProperty,
	docID uint64,
) error {
	for _, prop := range props {
		if prop.HasFilterableIndex {
			bucket, err := lsmkv.FetchMeABucket(s.store, "filterable_properties", helpers.BucketFromPropertyNameLSM(prop.Name), prop.Name, s.propIds)
			if err != nil {
				return fmt.Errorf("no bucket for prop '%s' found: %v", prop.Name, err)
			}

			for _, item := range prop.Items {
				if err := s.deleteFromPropertySetBucket(bucket, docID, item.Data); err != nil {
					return errors.Wrapf(err, "delete item '%s' from index",
						string(item.Data))
				}
			}
		}

		if prop.HasSearchableIndex {
			bucket, err := lsmkv.FetchMeABucket(s.store, "searchable_properties", helpers.BucketSearchableFromPropertyNameLSM(prop.Name), prop.Name, s.propIds)
			if err != nil {
				return fmt.Errorf("no bucket searchable for prop '%s: %v' found", prop.Name, err)
			}

			for _, item := range prop.Items {
				if err := s.deleteInvertedIndexItemWithFrequencyLSM(bucket, item,
					docID); err != nil {
					return errors.Wrapf(err, "delete item '%s' from index",
						string(item.Data))
				}
			}
		}

		// add non-nil properties to the null-state inverted index, but skip internal properties (__meta_count, _id etc)
		if isMetaCountProperty(prop) || isInternalProperty(prop) {
			continue
		}

		// properties where defining a length does not make sense (floats etc.) have a negative entry as length
		if s.index.invertedIndexConfig.IndexPropertyLength && prop.Length >= 0 {
			if err := s.deleteFromPropertyLengthIndex(prop.Name, docID, prop.Length); err != nil {
				return errors.Wrap(err, "add indexed property length")
			}
		}

		if s.index.invertedIndexConfig.IndexNullState {
			if err := s.deleteFromPropertyNullIndex(prop.Name, docID, prop.Length == 0); err != nil {
				return errors.Wrap(err, "add indexed null state")
			}
		}
	}

	// remove nil properties from the nullstate and property length inverted index
	for _, nilProperty := range nilProps {
		if s.index.invertedIndexConfig.IndexPropertyLength && nilProperty.AddToPropertyLength {
			if err := s.deleteFromPropertyLengthIndex(nilProperty.Name, docID, 0); err != nil {
				return errors.Wrap(err, "add indexed property length")
			}
		}

		if s.index.invertedIndexConfig.IndexNullState {
			if err := s.deleteFromPropertyNullIndex(nilProperty.Name, docID, true); err != nil {
				return errors.Wrap(err, "add indexed null state")
			}
		}
	}

	return nil
}

func (s *Shard) deleteInvertedIndexItemWithFrequencyLSM(bucket lsmkv.BucketInterface,
	item inverted.Countable, docID uint64,
) error {
	lsmkv.CheckExpectedStrategy(bucket.GetRegisteredName(), bucket.Strategy(), lsmkv.StrategyMapCollection)

	docIDBytes := make([]byte, 8)
	// Shard Index version 2 requires BigEndian for sorting, if the shard was
	// built prior assume it uses LittleEndian
	if s.versioner.Version() < 2 {
		binary.LittleEndian.PutUint64(docIDBytes, docID)
	} else {
		binary.BigEndian.PutUint64(docIDBytes, docID)
	}

	return bucket.MapDeleteKey(item.Data, docIDBytes)
}

func (s *Shard) deleteFromPropertyLengthIndex(propName string, docID uint64, length int) error {
	bucketLength, err := lsmkv.FetchMeABucket(s.store, "filterable_properties", helpers.BucketFromPropertyNameLengthLSM(propName), propName, s.propIds)
	if err != nil {
		return errors.Errorf("no bucket for prop '%s' length found: %v", propName, err)
	}
	bucketLength.CheckBucket()

	lsmkv.CheckExpectedStrategy(bucketLength.GetRegisteredName(), bucketLength.Strategy(), lsmkv.StrategySetCollection, lsmkv.StrategyRoaringSet)

	key, err := bucketKeyPropertyLength(length)
	if err != nil {
		return errors.Wrapf(err, "failed creating key for prop '%s' length", propName)
	}
	if err := s.deleteFromPropertySetBucket(bucketLength, docID, key); err != nil {
		return errors.Wrapf(err, "failed adding to prop '%s' length bucket", propName)
	}
	return nil
}

func (s *Shard) deleteFromPropertyNullIndex(propName string, docID uint64, isNull bool) error {
	bucketNull, err := lsmkv.FetchMeABucket(s.store, "filterable_properties", helpers.BucketFromPropertyNameNullLSM(propName), propName, s.propIds)
	if err != nil {
		return errors.Errorf("no bucket for prop '%s' null found", propName)
	}

	key, err := bucketKeyPropertyNull(isNull)
	if err != nil {
		return errors.Wrapf(err, "failed creating key for prop '%s' null", propName)
	}
	if err := s.deleteFromPropertySetBucket(bucketNull, docID, key); err != nil {
		return errors.Wrapf(err, "failed adding to prop '%s' null bucket", propName)
	}
	return nil
}

func (s *Shard) deleteFromPropertySetBucket(bucket lsmkv.BucketInterface, docID uint64, key []byte) error {
	lsmkv.CheckExpectedStrategy(bucket.Strategy(), lsmkv.StrategySetCollection, lsmkv.StrategyRoaringSet)

	if bucket.Strategy() == lsmkv.StrategySetCollection {
		docIDBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(docIDBytes, docID)

		return bucket.SetDeleteSingle(key, docIDBytes)
	}

	return bucket.RoaringSetRemoveOne(key, docID)
}
