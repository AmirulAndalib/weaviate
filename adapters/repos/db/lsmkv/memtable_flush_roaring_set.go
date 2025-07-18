//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package lsmkv

import (
	"fmt"

	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv/segmentindex"
	"github.com/weaviate/weaviate/adapters/repos/db/roaringset"
)

func (m *Memtable) flushDataRoaringSet(f *segmentindex.SegmentFile) ([]segmentindex.Key, error) {
	flat := m.roaringSet.FlattenInOrder()

	totalDataLength := totalPayloadSizeRoaringSet(flat)
	header := &segmentindex.Header{
		IndexStart:       uint64(totalDataLength + segmentindex.HeaderSize),
		Level:            0, // always level zero on a new one
		Version:          segmentindex.ChooseHeaderVersion(m.enableChecksumValidation),
		SecondaryIndices: 0,
		Strategy:         segmentindex.StrategyRoaringSet,
	}

	n, err := f.WriteHeader(header)
	if err != nil {
		return nil, err
	}
	headerSize := int(n)
	keys := make([]segmentindex.Key, len(flat))

	totalWritten := headerSize
	for i, node := range flat {
		sn, err := roaringset.NewSegmentNode(node.Key, node.Value.Additions,
			node.Value.Deletions)
		if err != nil {
			return nil, fmt.Errorf("create segment node: %w", err)
		}

		ki, err := sn.KeyIndexAndWriteTo(f.BodyWriter(), totalWritten)
		if err != nil {
			return nil, fmt.Errorf("write node %d: %w", i, err)
		}

		keys[i] = ki
		totalWritten = ki.ValueEnd
	}

	return keys, nil
}

func totalPayloadSizeRoaringSet(in []*roaringset.BinarySearchNode) int {
	var sum int
	for _, n := range in {
		sum += 8 // uint64 to segment length
		sum += 8 // uint64 to indicate length of additions bitmap
		sum += len(n.Value.Additions.ToBuffer())
		sum += 8 // uint64 to indicate length of deletions bitmap
		sum += len(n.Value.Deletions.ToBuffer())
		sum += 4 // uint32 to indicate key size
		sum += len(n.Key)
	}

	return sum
}
