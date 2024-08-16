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

package roaringsetrange

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/weaviate/sroar"
	"github.com/weaviate/weaviate/adapters/repos/db/roaringset"
	"github.com/weaviate/weaviate/entities/errors"
	"github.com/weaviate/weaviate/entities/filters"
)

type SegmentReader struct {
	cursor *GaplessSegmentCursor
}

func NewSegmentReader(cursor *GaplessSegmentCursor) *SegmentReader {
	return &SegmentReader{cursor: cursor}
}

func (r *SegmentReader) Read(ctx context.Context, value uint64, operator filters.Operator,
) (roaringset.BitmapLayer, error) {
	if err := ctx.Err(); err != nil {
		return roaringset.BitmapLayer{}, err
	}

	switch operator {
	case filters.OperatorEqual:
		return r.readEqual(ctx, value)

	case filters.OperatorNotEqual:
		return r.readNotEqual(ctx, value)

	case filters.OperatorLessThan:
		return r.readLessThan(ctx, value)

	case filters.OperatorLessThanEqual:
		return r.readLessThanEqual(ctx, value)

	case filters.OperatorGreaterThan:
		return r.readGreaterThan(ctx, value)

	case filters.OperatorGreaterThanEqual:
		return r.readGreaterThanEqual(ctx, value)

	default:
		return roaringset.BitmapLayer{}, fmt.Errorf("operator %v not supported for segments of strategy %q",
			operator.Name(), "roaringsetrange") // TODO move strategies to separate package?
	}
}

func (r *SegmentReader) firstLayer() (roaringset.BitmapLayer, bool) {
	_, layer, ok := r.cursor.First()
	if !ok {
		return roaringset.BitmapLayer{
			Additions: sroar.NewBitmap(),
			Deletions: sroar.NewBitmap(),
		}, false
	}

	// additions := layer.Additions
	// if additions == nil {
	// 	additions = sroar.NewBitmap()
	// }
	// deletions := layer.Deletions
	// if deletions == nil {
	// 	deletions = sroar.NewBitmap()
	// }

	// return roaringset.BitmapLayer{
	// 	Additions: additions,
	// 	Deletions: deletions,
	// }, !additions.IsEmpty()

	var deletions *sroar.Bitmap
	if layer.Deletions == nil {
		deletions = sroar.NewBitmap()
	} else {
		deletions = layer.Deletions.Clone()
	}

	if layer.Additions.IsEmpty() {
		return roaringset.BitmapLayer{
			Additions: sroar.NewBitmap(),
			Deletions: deletions,
		}, false
	}
	// additions will be cloned when needed by merging methods
	return roaringset.BitmapLayer{
		Additions: layer.Additions,
		Deletions: deletions,
	}, true
}

func (r *SegmentReader) readEqual(ctx context.Context, value uint64,
) (roaringset.BitmapLayer, error) {
	if value == 0 {
		return r.readLessThanEqual(ctx, value)
	}
	if value == math.MaxUint64 {
		return r.readGreaterThanEqual(ctx, value)
	}

	firstLayer, ok := r.firstLayer()
	if !ok {
		return firstLayer, nil
	}

	eq, err := r.mergeBetween(ctx, value, value+1, firstLayer.Additions)
	if err != nil {
		return roaringset.BitmapLayer{}, err
	}

	return roaringset.BitmapLayer{
		Additions: eq,
		Deletions: firstLayer.Deletions,
	}, nil
}

func (r *SegmentReader) readNotEqual(ctx context.Context, value uint64,
) (roaringset.BitmapLayer, error) {
	if value == 0 {
		return r.readGreaterThan(ctx, value)
	}
	if value == math.MaxUint64 {
		return r.readLessThan(ctx, value)
	}

	firstLayer, ok := r.firstLayer()
	if !ok {
		return firstLayer, nil
	}

	neq := firstLayer.Additions.Clone()
	eq, err := r.mergeBetween(ctx, value, value+1, firstLayer.Additions)
	if err != nil {
		return roaringset.BitmapLayer{}, err
	}

	neq.AndNot(eq)
	return roaringset.BitmapLayer{
		Additions: neq,
		Deletions: firstLayer.Deletions,
	}, nil
}

func (r *SegmentReader) readLessThan(ctx context.Context, value uint64,
) (roaringset.BitmapLayer, error) {
	firstLayer, ok := r.firstLayer()
	if !ok {
		return firstLayer, nil
	}

	if value == 0 {
		// no value is < 0
		return roaringset.BitmapLayer{
			Additions: sroar.NewBitmap(),
			Deletions: firstLayer.Deletions,
		}, nil
	}

	lt := firstLayer.Additions.Clone()
	gte, err := r.mergeGreaterThanEqual(ctx, value, firstLayer.Additions)
	if err != nil {
		return roaringset.BitmapLayer{}, err
	}

	lt.AndNot(gte)
	return roaringset.BitmapLayer{
		Additions: lt,
		Deletions: firstLayer.Deletions,
	}, nil
}

func (r *SegmentReader) readLessThanEqual(ctx context.Context, value uint64,
) (roaringset.BitmapLayer, error) {
	firstLayer, ok := r.firstLayer()
	if !ok {
		return firstLayer, nil
	}

	if value == math.MaxUint64 {
		// all values are <= max uint64
		return firstLayer, nil
	}

	lte := firstLayer.Additions.Clone()
	gte1, err := r.mergeGreaterThanEqual(ctx, value+1, firstLayer.Additions)
	if err != nil {
		return roaringset.BitmapLayer{}, err
	}

	lte.AndNot(gte1)
	return roaringset.BitmapLayer{
		Additions: lte,
		Deletions: firstLayer.Deletions,
	}, nil
}

func (r *SegmentReader) readGreaterThan(ctx context.Context, value uint64,
) (roaringset.BitmapLayer, error) {
	firstLayer, ok := r.firstLayer()
	if !ok {
		return firstLayer, nil
	}

	if value == math.MaxUint64 {
		// no value is > max uint64
		return roaringset.BitmapLayer{
			Additions: sroar.NewBitmap(),
			Deletions: firstLayer.Deletions,
		}, nil
	}

	gte1, err := r.mergeGreaterThanEqual(ctx, value+1, firstLayer.Additions)
	if err != nil {
		return roaringset.BitmapLayer{}, err
	}

	return roaringset.BitmapLayer{
		Additions: gte1,
		Deletions: firstLayer.Deletions,
	}, nil
}

func (r *SegmentReader) readGreaterThanEqual(ctx context.Context, value uint64,
) (roaringset.BitmapLayer, error) {
	firstLayer, ok := r.firstLayer()
	if !ok {
		return firstLayer, nil
	}

	gte, err := r.mergeGreaterThanEqual(ctx, value, firstLayer.Additions)
	if err != nil {
		return roaringset.BitmapLayer{}, err
	}

	return roaringset.BitmapLayer{
		Additions: gte,
		Deletions: firstLayer.Deletions,
	}, nil
}

func (r *SegmentReader) mergeGreaterThanEqual(ctx context.Context, value uint64,
	all *sroar.Bitmap,
) (*sroar.Bitmap, error) {
	ANDed := false
	result := all.Clone()
	// t_convert := time.Now()
	// result.ConvertToBitmaps()
	// fmt.Printf("  ==> convert took [%s]\n", time.Since(t_convert))
	// result2 := all.Clone()

	buf := make([]uint16, 4100)

	t_total := time.Now()
	t_next := time.Now()
	var t_and, t_or, t_condense time.Time
	var d_next, d_and, d_condense, d_or, d_next_total, d_and_total, d_condense_total, d_or_total time.Duration

	// fmt.Printf(" => [all] cardinality [%v]\n", all.GetCardinality())

	type cur struct {
		bit   uint8
		layer roaringset.BitmapLayer
	}
	ch := make(chan *cur)

	errors.GoWrapper(func() {
		for bit, layer, ok := r.cursor.Next(); ok; bit, layer, ok = r.cursor.Next() {
			if ctx.Err() != nil {
				break
			}
			ch <- &cur{bit: bit, layer: layer}
		}
		close(ch)
	}, nil)

	for bl := range ch {
		bit := bl.bit
		layer := bl.layer

		d_next = time.Since(t_next)
		d_next_total += d_next

		// if ctx.Err() != nil {
		// 	return nil, ctx.Err()
		// }

		// if bit == 57 {
		// 	break
		// }
		// if bit == 3 {
		// 	continue
		// }

		// fmt.Printf(" => [%v] cardinality [%v]\n", bit, layer.Additions.GetCardinality())

		// _ = bit
		_ = layer
		// _ = t_and
		// _ = t_or
		// _ = t_condense
		// _ = d_and
		// _ = d_or
		// _ = d_condense

		if value&(1<<(bit-1)) != 0 {
			// if !ANDed {
			// 	result = result.Clone()
			// 	// result = sroar.ConvertToBitmapContainers(result)
			// }
			ANDed = true
			t_and = time.Now()
			// fmt.Printf("  ==> AND with bit [%v]\n", bit)
			result.And2(layer.Additions, buf)
			// result2.And2(layer.Additions)
			d_and = time.Since(t_and)
			d_and_total += d_and
			t_condense = time.Now()
			// result = roaringset.Condense(result)
			d_condense = time.Since(t_condense)
			d_condense_total += d_condense
		} else if ANDed {
			t_or = time.Now()
			// fmt.Printf("  ==> OR with bit [%v]\n", bit)
			result.Or2(layer.Additions, buf)
			// result2.Or2(layer.Additions)
			d_or = time.Since(t_or)
			d_or_total += d_or
		}

		t_next = time.Now()
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// if !ANDed {
	// 	result = result.Clone()
	// }

	d_total := time.Since(t_total)
	_ = d_total
	fmt.Printf("  ==> total time [%s]\n", d_total)
	fmt.Printf("  ==> cursor time [%s]\n", d_next_total)
	fmt.Printf("  ==> or time [%s]\n", d_or_total)
	fmt.Printf("  ==> and time [%s]\n", d_and_total)
	fmt.Printf("  ==> condense time [%s]\n\n", d_condense_total)

	// fmt.Printf("  ==> result card [%v]\n", result.GetCardinality())
	// fmt.Printf("  ==> result2 card [%v]\n\n", result2.GetCardinality())

	return result, nil
}

// func (r *SegmentReader) mergeGreaterThanEqual(ctx context.Context, value uint64,
// 	all *sroar.Bitmap,
// ) (*sroar.Bitmap, error) {
// 	ANDed := false
// 	result := all.Clone()
// 	// t_convert := time.Now()
// 	// result.ConvertToBitmaps()
// 	// fmt.Printf("  ==> convert took [%s]\n", time.Since(t_convert))
// 	// result2 := all.Clone()

// 	_ = errors.GoWrapper
// 	buf := make([]uint16, 4100)

// 	t_total := time.Now()
// 	t_next := time.Now()
// 	var t_and, t_or, t_condense time.Time
// 	var d_next, d_and, d_condense, d_or, d_next_total, d_and_total, d_condense_total, d_or_total time.Duration

// 	// fmt.Printf(" => [all] cardinality [%v]\n", all.GetCardinality())

// 	for bit, layer, ok := r.cursor.Next(); ok; bit, layer, ok = r.cursor.Next() {
// 		d_next = time.Since(t_next)
// 		d_next_total += d_next

// 		if ctx.Err() != nil {
// 			return nil, ctx.Err()
// 		}

// 		// if bit == 57 {
// 		// 	break
// 		// }
// 		// if bit == 3 {
// 		// 	continue
// 		// }

// 		// fmt.Printf(" => [%v] cardinality [%v]\n", bit, layer.Additions.GetCardinality())

// 		// _ = bit
// 		_ = layer
// 		// _ = t_and
// 		// _ = t_or
// 		// _ = t_condense
// 		// _ = d_and
// 		// _ = d_or
// 		// _ = d_condense

// 		if value&(1<<(bit-1)) != 0 {
// 			// if !ANDed {
// 			// 	result = result.Clone()
// 			// 	// result = sroar.ConvertToBitmapContainers(result)
// 			// }
// 			ANDed = true
// 			t_and = time.Now()
// 			// fmt.Printf("  ==> AND with bit [%v]\n", bit)
// 			result.And2(layer.Additions, buf)
// 			// result2.And2(layer.Additions)
// 			d_and = time.Since(t_and)
// 			d_and_total += d_and
// 			t_condense = time.Now()
// 			// result = roaringset.Condense(result)
// 			d_condense = time.Since(t_condense)
// 			d_condense_total += d_condense
// 		} else if ANDed {
// 			t_or = time.Now()
// 			// fmt.Printf("  ==> OR with bit [%v]\n", bit)
// 			result.Or2(layer.Additions, buf)
// 			// result2.Or2(layer.Additions)
// 			d_or = time.Since(t_or)
// 			d_or_total += d_or
// 		}

// 		t_next = time.Now()
// 	}

// 	// if !ANDed {
// 	// 	result = result.Clone()
// 	// }

// 	d_total := time.Since(t_total)
// 	_ = d_total
// 	fmt.Printf("  ==> total time [%s]\n", d_total)
// 	fmt.Printf("  ==> cursor time [%s]\n", d_next_total)
// 	fmt.Printf("  ==> or time [%s]\n", d_or_total)
// 	fmt.Printf("  ==> and time [%s]\n", d_and_total)
// 	fmt.Printf("  ==> condense time [%s]\n\n", d_condense_total)

// 	// fmt.Printf("  ==> result card [%v]\n", result.GetCardinality())
// 	// fmt.Printf("  ==> result2 card [%v]\n\n", result2.GetCardinality())

// 	return result, nil
// }

func (r *SegmentReader) mergeBetween(ctx context.Context, valueMinInc, valueMaxExc uint64,
	all *sroar.Bitmap,
) (*sroar.Bitmap, error) {
	ANDedMin := false
	ANDedMax := false
	resultMin := all
	resultMax := all

	for bit, layer, ok := r.cursor.Next(); ok; bit, layer, ok = r.cursor.Next() {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		var b uint64 = 1 << (bit - 1)

		if valueMinInc&b != 0 {
			if !ANDedMin {
				resultMin = resultMin.Clone()
			}
			ANDedMin = true
			resultMin.And(layer.Additions)
			resultMin = roaringset.Condense(resultMin)
		} else if ANDedMin {
			resultMin.Or(layer.Additions)
		}

		if valueMaxExc&b != 0 {
			if !ANDedMax {
				resultMax = resultMax.Clone()
			}
			ANDedMax = true
			resultMax.And(layer.Additions)
			resultMax = roaringset.Condense(resultMax)
		} else if ANDedMax {
			resultMax.Or(layer.Additions)
		}
	}

	if !ANDedMin {
		resultMin = resultMin.Clone()
	}
	resultMin.AndNot(resultMax)

	return resultMin, nil
}
