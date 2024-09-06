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

package inverted

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/weaviate/weaviate/entities/additional"

	enterrors "github.com/weaviate/weaviate/entities/errors"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/sroar"
	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
	"github.com/weaviate/weaviate/adapters/repos/db/inverted/stopwords"
	"github.com/weaviate/weaviate/adapters/repos/db/inverted/terms"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv"
	"github.com/weaviate/weaviate/adapters/repos/db/priorityqueue"
	"github.com/weaviate/weaviate/adapters/repos/db/propertyspecific"
	"github.com/weaviate/weaviate/entities/inverted"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
	"github.com/weaviate/weaviate/entities/searchparams"
	"github.com/weaviate/weaviate/entities/storobj"
)

type BM25Searcher struct {
	config         schema.BM25Config
	store          *lsmkv.Store
	getClass       func(string) *models.Class
	classSearcher  ClassSearcher // to allow recursive searches on ref-props
	propIndices    propertyspecific.Indices
	propLenTracker propLengthRetriever
	logger         logrus.FieldLogger
	shardVersion   uint16
}

type propLengthRetriever interface {
	PropertyMean(prop string) (float32, error)
}

func NewBM25Searcher(config schema.BM25Config, store *lsmkv.Store,
	getClass func(string) *models.Class, propIndices propertyspecific.Indices,
	classSearcher ClassSearcher, propLenTracker propLengthRetriever,
	logger logrus.FieldLogger, shardVersion uint16,
) *BM25Searcher {
	return &BM25Searcher{
		config:         config,
		store:          store,
		getClass:       getClass,
		propIndices:    propIndices,
		classSearcher:  classSearcher,
		propLenTracker: propLenTracker,
		logger:         logger.WithField("action", "bm25_search"),
		shardVersion:   shardVersion,
	}
}

func (b *BM25Searcher) extractTermInformation(class *models.Class, params searchparams.KeywordRanking) (map[string][]string, map[string][]int, map[string][]string, map[string]float32, float64, error) {
	var stopWordDetector *stopwords.Detector
	var err error
	if class.InvertedIndexConfig != nil && class.InvertedIndexConfig.Stopwords != nil {
		stopWordDetector, err = stopwords.NewDetectorFromConfig(*(class.InvertedIndexConfig.Stopwords))
	}
	if err != nil {
		return nil, nil, nil, nil, 0, err
	}

	queryTermsByTokenization := map[string][]string{}
	duplicateBoostsByTokenization := map[string][]int{}
	propNamesByTokenization := map[string][]string{}
	propertyBoosts := make(map[string]float32, len(params.Properties))

	for _, tokenization := range helpers.Tokenizations {
		queryTerms, dupBoosts := helpers.TokenizeAndCountDuplicates(tokenization, params.Query)
		queryTermsByTokenization[tokenization] = queryTerms
		duplicateBoostsByTokenization[tokenization] = dupBoosts

		if tokenization == models.PropertyTokenizationWord {
			queryTerms, dupBoosts = b.removeStopwordsFromQueryTerms(queryTermsByTokenization[tokenization],
				duplicateBoostsByTokenization[tokenization], stopWordDetector)
			queryTermsByTokenization[tokenization] = queryTerms
			duplicateBoostsByTokenization[tokenization] = dupBoosts
		}
		propNamesByTokenization[tokenization] = make([]string, 0)
	}

	averagePropLength := 0.
	for _, propertyWithBoost := range params.Properties {
		property := propertyWithBoost
		propBoost := 1
		if strings.Contains(propertyWithBoost, "^") {
			property = strings.Split(propertyWithBoost, "^")[0]
			boostStr := strings.Split(propertyWithBoost, "^")[1]
			propBoost, _ = strconv.Atoi(boostStr)
		}
		propertyBoosts[property] = float32(propBoost)

		propMean, err := b.GetPropertyLengthTracker().PropertyMean(property)
		if err != nil {
			return nil, nil, nil, nil, 0, err
		}
		averagePropLength += float64(propMean)

		prop, err := schema.GetPropertyByName(class, property)
		if err != nil {
			return nil, nil, nil, nil, 0, err
		}

		switch dt, _ := schema.AsPrimitive(prop.DataType); dt {
		case schema.DataTypeText, schema.DataTypeTextArray:
			if _, exists := propNamesByTokenization[prop.Tokenization]; !exists {
				return nil, nil, nil, nil, 0, fmt.Errorf("cannot handle tokenization '%v' of property '%s'",
					prop.Tokenization, prop.Name)
			}
			propNamesByTokenization[prop.Tokenization] = append(propNamesByTokenization[prop.Tokenization], property)
		default:
			return nil, nil, nil, nil, 0, fmt.Errorf("cannot handle datatype '%v' of property '%s'", dt, prop.Name)
		}
	}
	averagePropLength = averagePropLength / float64(len(params.Properties))

	return queryTermsByTokenization, duplicateBoostsByTokenization, propNamesByTokenization, propertyBoosts, averagePropLength, nil
}

func (b *BM25Searcher) BM25F(ctx context.Context, filterDocIds helpers.AllowList,
	className schema.ClassName, limit int, keywordRanking searchparams.KeywordRanking, additional additional.Properties,
) ([]*storobj.Object, []float32, error) {
	// WEAVIATE-471 - If a property is not searchable, return an error
	for _, property := range keywordRanking.Properties {
		if !PropertyHasSearchableIndex(b.getClass(className.String()), property) {
			return nil, nil, inverted.NewMissingSearchableIndexError(property)
		}
	}
	class := b.getClass(className.String())
	if class == nil {
		return nil, nil, fmt.Errorf("could not find class %s in schema", className)
	}

	objs, scores, err := b.wand(ctx, filterDocIds, class, keywordRanking, limit, additional)
	if err != nil {
		return nil, nil, errors.Wrap(err, "wand")
	}
	return objs, scores, nil
}

func (b *BM25Searcher) GetPropertyLengthTracker() *JsonShardMetaData {
	return b.propLenTracker.(*JsonShardMetaData)
}

func (b *BM25Searcher) wand(
	ctx context.Context, filterDocIds helpers.AllowList, class *models.Class, params searchparams.KeywordRanking, limit int, additionalProps additional.Properties,
) ([]*storobj.Object, []float32, error) {
	useWandDisk := os.Getenv("USE_WAND_DISK") == "true"
	validateWandDisk := os.Getenv("USE_WAND_DISK") == "validate"
	if useWandDisk {
		return b.wandDisk(ctx, filterDocIds, class, params, limit)
	} else if validateWandDisk {
		return b.validateWand(ctx, filterDocIds, class, params, limit, additionalProps)
	} else {
		return b.wandMem(ctx, filterDocIds, class, params, limit, additionalProps)
	}
}

func (b *BM25Searcher) validateWand(
	ctx context.Context, filterDocIds helpers.AllowList, class *models.Class, params searchparams.KeywordRanking, limit int, additionalProps additional.Properties,
) ([]*storobj.Object, []float32, error) {
	objsD, scoresD, errD := b.wandDisk(ctx, filterDocIds, class, params, limit)
	objsM, scoresM, errM := b.wandMem(ctx, filterDocIds, class, params, limit, additionalProps)
	if errD != nil {
		return nil, nil, errD
	}
	if errM != nil {
		return nil, nil, errM
	}

	var err error
	// compare results and scores
	if len(objsD) != len(objsM) {
		fmt.Printf("different number of results: disk %d, mem %d", len(objsD), len(objsM))
		err = fmt.Errorf("different number of results: disk %d, mem %d", len(objsD), len(objsM))
	}
	if len(scoresD) != len(scoresM) {
		fmt.Printf("different number of scores: disk %d, mem %d", len(objsD), len(objsM))
		err = fmt.Errorf("different number of scores: disk %d, mem %d", len(scoresD), len(scoresM))
	}

	if err != nil {
		for i := range objsM {
			if len(objsD) <= i && len(objsM) <= i {
				fmt.Printf("disk %v,%v mem %v,%v\n", scoresD[i], objsD[i].ID(), scoresM[i], objsM[i].ID())
			}
		}
		// return nil, nil, err
	}

	for i := range objsM {
		if len(objsD) <= i && len(objsM) <= i {
			if objsD[i].ID() != objsM[i].ID() {
				err = fmt.Errorf("different IDs at index %d: disk %v,%v mem %v,%v", i, scoresD[i], objsD[i].ID(), scoresM[i], objsM[i].ID())
				fmt.Printf("different IDs at index %d: disk %v,%v mem %v,%v\n", i, scoresD[i], objsD[i].ID(), scoresM[i], objsM[i].ID())
				break
			}
			if scoresD[i] != scoresM[i] {
				err = fmt.Errorf("different scores at index %d: disk %v,%v mem %v,%v", i, scoresD[i], objsD[i].ID(), scoresM[i], objsM[i].ID())
				fmt.Printf("different scores at index %d: disk %v,%v mem %v,%v\n", i, scoresD[i], objsD[i].ID(), scoresM[i], objsM[i].ID())
				break
			}
		}
	}

	if err != nil {
		for i := range objsM {
			if len(objsD) <= i && len(objsM) <= i {
				fmt.Printf("disk %v,%v mem %v,%v\n", scoresD[i], objsD[i].ID(), scoresM[i], objsM[i].ID())
			}
		}
		// return nil, nil, err
	}

	return objsM, scoresM, errM
}

// global var to store wand times
/*
var (
	wandMemTimes     []float64 = make([]float64, 5)
	wandMemCounter   int       = 0
	wandMemLastClass string
	wandMemStats     []float64 = make([]float64, 1)
)
*/
func (b *BM25Searcher) wandMem(ctx context.Context, filterDocIds helpers.AllowList, class *models.Class, params searchparams.KeywordRanking, limit int, additionalProps additional.Properties) ([]*storobj.Object, []float32, error) {
	/*
		if wandMemLastClass == "" {
			wandMemLastClass = string(class.Class)
		}
		if wandMemLastClass != string(class.Class) {
			fmt.Printf(" MEM,%v", wandMemLastClass)
			sum := 0.
			for i := 0; i < len(wandMemTimes); i++ {
				fmt.Printf(",%8.2f", wandMemTimes[i]/float64(wandMemCounter))
				sum += wandMemTimes[i] / float64(wandMemCounter)
			}
			fmt.Printf(",%8.2f", sum)

			// for i := 0; i < len(wandMemStats); i++ {
			//	fmt.Printf(",%8.2f", wandMemStats[i]/float64(wandMemCounter))
			// }
			fmt.Printf("\n")
			wandMemCounter = 0
			wandMemLastClass = string(class.Class)
			wandMemTimes = make([]float64, len(wandMemTimes))
			wandMemStats = make([]float64, len(wandMemStats))
		}

		wandMemCounter++
		wandTimesId := 0

		// start timer
		startTime := float64(time.Now().UnixNano()) / 1e6
	*/

	// get number of objects
	N := float64(b.store.Bucket(helpers.ObjectsBucketLSM).Count())

	// stopword filtering for word tokenization
	queryTermsByTokenization, duplicateBoostsByTokenization, propNamesByTokenization, propertyBoosts, averagePropLength, err := b.extractTermInformation(class, params)
	if err != nil {
		return nil, nil, err
	}

	// wandMemTimes[wandTimesId] += float64(time.Now().UnixNano())/1e6 - startTime
	// wandTimesId++

	objsM, scoresM, errM := b.wandMemScoring(ctx, queryTermsByTokenization, duplicateBoostsByTokenization, propNamesByTokenization, propertyBoosts, averagePropLength, N, filterDocIds, params, limit)

	return objsM, scoresM, errM
}

func (b *BM25Searcher) wandMemScoring(ctx context.Context, queryTermsByTokenization map[string][]string, duplicateBoostsByTokenization map[string][]int, propNamesByTokenization map[string][]string, propertyBoosts map[string]float32, averagePropLength float64, N float64, filterDocIds helpers.AllowList, params searchparams.KeywordRanking, limit int) ([]*storobj.Object, []float32, error) {
	// wandTimesId := 1
	// startTime := float64(time.Now().UnixNano()) / 1e6

	tokenCount := 0
	for tokenization, propNames := range propNamesByTokenization {
		if len(propNames) > 0 {
			tokenCount += len(queryTermsByTokenization[tokenization])
		}
	}

	results := make([]terms.Term, tokenCount)
	indices := make([]map[uint64]int, tokenCount)

	eg := enterrors.NewErrorGroupWrapper(b.logger)
	eg.SetLimit(_NUMCPU)

	tokenIndex := 0
	queryTerms := make([]string, 0)
	for tokenization, propNames := range propNamesByTokenization {
		propNames := propNames
		if len(propNames) > 0 {
			for queryTermId, queryTerm := range queryTermsByTokenization[tokenization] {
				queryTerms = append(queryTerms, queryTerm)
				tokenization := tokenization
				queryTerm := queryTerm
				queryTermId := queryTermId
				tokenIndexInner := tokenIndex

				dupBoost := duplicateBoostsByTokenization[tokenization][queryTermId]

				eg.Go(func() (err error) {
					termResult, docIndices, termErr := b.createTerm(ctx, N, filterDocIds, queryTerm, queryTermId, propNames, propertyBoosts, dupBoost, params.AdditionalExplanations)
					if termErr != nil {
						err = termErr
						return
					}
					results[tokenIndexInner] = termResult
					indices[tokenIndexInner] = docIndices
					return
				}, "query_term", queryTerm, "prop_names", propNames, "has_filter", filterDocIds != nil)

				tokenIndex++
			}
		}
	}

	if err := eg.Wait(); err != nil {
		return nil, nil, err
	}
	// all results. Sum up the length of the results from all terms to get an upper bound of how many results there are
	if limit == 0 {
		limit = int(N)
	}

	// wandMemTimes[wandTimesId] += float64(time.Now().UnixNano())/1e6 - startTime
	// wandTimesId++
	// startTime = float64(time.Now().UnixNano()) / 1e6

	// the results are needed in the original order to be able to locate frequency/property length for the top-results
	resultsOriginalOrder := terms.Terms{}
	resultsTerms := terms.Terms{}
	resultsTerms.T = results
	resultsTerms.Count = len(queryTerms)

	resultsOriginalOrder.T = make([]terms.Term, len(results))

	copy(resultsOriginalOrder.T, resultsTerms.T)

	topKHeap := b.getTopKHeap(limit, resultsTerms, averagePropLength, params.AdditionalExplanations)

	// wandMemTimes[wandTimesId] += float64(time.Now().UnixNano())/1e6 - startTime
	// wandTimesId++
	// startTime = float64(time.Now().UnixNano()) / 1e6

	objects, scores, err := b.getTopKObjects(topKHeap, resultsOriginalOrder, params.AdditionalExplanations, queryTerms)

	// wandMemTimes[wandTimesId] += float64(time.Now().UnixNano())/1e6 - startTime
	// wandTimesId++

	//for i, obj := range objects {
	//	fmt.Printf("%v: %v %v\n", obj.DocID, scores[i], obj.AdditionalProperties())
	//}

	return objects, scores, err
}

func (b *BM25Searcher) removeStopwordsFromQueryTerms(queryTerms []string,
	duplicateBoost []int, detector *stopwords.Detector,
) ([]string, []int) {
	if detector == nil || len(queryTerms) == 0 {
		return queryTerms, duplicateBoost
	}

	i := 0
WordLoop:
	for {
		if i == len(queryTerms) {
			return queryTerms, duplicateBoost
		}
		queryTerm := queryTerms[i]
		if detector.IsStopword(queryTerm) {
			queryTerms[i] = queryTerms[len(queryTerms)-1]
			queryTerms = queryTerms[:len(queryTerms)-1]
			duplicateBoost[i] = duplicateBoost[len(duplicateBoost)-1]
			duplicateBoost = duplicateBoost[:len(duplicateBoost)-1]

			continue WordLoop
		}

		i++
	}
}

func (b *BM25Searcher) getTopKHeapExaustive(limit int, results terms.Terms, averagePropLength float64, additionalExplanations bool,
) *priorityqueue.Queue[[]*terms.DocPointerWithScore] {
	eg := enterrors.NewErrorGroupWrapper(b.logger)
	eg.SetLimit(len(results.T))

	scores := make([]map[uint64]float64, len(results.T))
	docInfos := make([]map[uint64]*terms.DocPointerWithScore, len(results.T))

	for i, term := range results.T {
		term := term
		i := i

		scores[i] = make(map[uint64]float64)
		docInfos[i] = make(map[uint64]*terms.DocPointerWithScore)
		eg.Go(func() error {
			analysed := 0
			for {
				if term.IsExhausted() {

					fmt.Printf("analysed: %v\n", analysed)
					return nil
				}

				id, score, docInfo := term.ScoreAndAdvance(averagePropLength, b.config)

				docInfos[i][id] = &docInfo
				scores[i][id] = score
				analysed++
			}
		}, "term", term.QueryTerm())

	}

	if err := eg.Wait(); err != nil {
		return nil
	}

	finalScores := make(map[uint64]float64, len(scores))

	for score := range scores {
		for id, s := range scores[score] {
			finalScores[id] += s
		}
	}

	topKHeap := priorityqueue.NewMinScoreAndId[[]*terms.DocPointerWithScore](limit)

	for id, score := range finalScores {
		if topKHeap.Len() > 0 {
			fmt.Printf("%v: %v %v\n", id, score, topKHeap.Top().Dist)
		}
		if topKHeap.Len() < limit || topKHeap.Top().Dist < float32(score) {
			docInfosLocal := make([]*terms.DocPointerWithScore, len(results.T))
			for i, docInfo := range docInfos {
				docInfosLocal[i] = docInfo[id]
			}

			topKHeap.InsertWithValue(id, float32(score), docInfosLocal)
			for topKHeap.Len() > limit {
				topKHeap.Pop()
			}
		}
	}
	return topKHeap
}

func (b *BM25Searcher) getTopKObjects(topKHeap *priorityqueue.Queue[[]*terms.DocPointerWithScore],
	results terms.Terms, additionalExplanations bool, allTerms []string,
) ([]*storobj.Object, []float32, error) {
	objectsBucket := b.store.Bucket(helpers.ObjectsBucketLSM)
	if objectsBucket == nil {
		return nil, nil, errors.Errorf("objects bucket not found")
	}

	objects := make([]*storobj.Object, 0, topKHeap.Len())
	scores := make([]float32, 0, topKHeap.Len())
	ids := make([]uint64, 0, topKHeap.Len())

	buf := make([]byte, 8)
	for topKHeap.Len() > 0 {
		res := topKHeap.Pop()
		binary.LittleEndian.PutUint64(buf, res.ID)
		objectByte, err := objectsBucket.GetBySecondary(0, buf)
		if err != nil {
			return nil, nil, err
		}

		// If there is a crash and WAL recovery, the inverted index may have objects that are not in the objects bucket.
		// This is an issue that needs to be fixed, but for now we need to reduce the huge amount of log messages that
		// are generated by this issue. Logging the first time we encounter a missing object in a query still resulted
		// in a huge amount of log messages and it will happen on all queries, so we not log at all for now.
		// The user has already been alerted about ppossible data loss when the WAL recovery happened.
		// TODO: consider deleting these entries from the inverted index and alerting the user
		if len(objectByte) == 0 {
			continue
		}

		obj, err := storobj.FromBinary(objectByte)
		if err != nil {
			return nil, nil, err
		}

		if additionalExplanations {
			if obj.AdditionalProperties() == nil {
				obj.Object.Additional = make(map[string]interface{})
			}
			for j := range res.Value {
				if res.Value[j] == nil {
					continue
				}
				queryTerm := allTerms[j]
				obj.Object.Additional["BM25F_"+queryTerm+"_frequency"] = res.Value[j].Frequency
				obj.Object.Additional["BM25F_"+queryTerm+"_propLength"] = res.Value[j].PropLength
			}

		}
		objects = append(objects, obj)
		scores = append(scores, res.Dist)

	}

	// handle case that an object was removed
	if len(objects) != len(scores) {
		j := 0
		for i := range scores {
			if j >= len(objects) {
				break
			}
			if objects[j].DocID != ids[i] {
				continue
			}
			scores[j] = scores[i]
			j++
		}
		scores = scores[:j]
	}

	return objects, scores, nil
}

func (b *BM25Searcher) getTopKHeap(limit int, results terms.Terms, averagePropLength float64, additionalExplanations bool,
) *priorityqueue.Queue[[]*terms.DocPointerWithScore] {
	topKHeap := priorityqueue.NewMinScoreAndId[[]*terms.DocPointerWithScore](limit)
	worstDist := float64(-10000) // tf score can be negative
	for {
		results.FullSort()
		if results.CompletelyExhausted() || results.Pivot(worstDist) {
			return topKHeap
		}

		id, score, docInfos := results.ScoreNext(averagePropLength, b.config, additionalExplanations)

		if topKHeap.Len() < limit || topKHeap.Top().Dist < float32(score) {
			topKHeap.InsertWithValue(id, float32(score), docInfos)
			for topKHeap.Len() > limit {
				topKHeap.Pop()
			}
			// only update the worst distance when the queue is full, otherwise results can be missing if the first
			// entry that is checked already has a very high score
			if topKHeap.Len() >= limit {
				worstDist = float64(topKHeap.Top().Dist)
			}
		}
	}
}

func (b *BM25Searcher) createTerm(ctx context.Context, N float64, filterDocIds helpers.AllowList, query string, queryTermId int,
	propertyNames []string, propertyBoosts map[string]float32, duplicateTextBoost int, additional bool,
) (terms.Term, map[uint64]int, error) {
	termResult := &termMem{queryTerm: query, index: queryTermId}

	eg := enterrors.NewErrorGroupWrapper(b.logger)
	eg.SetLimit(_NUMCPU)

	allMsAndProps := make(AllMapPairsAndPropName, len(propertyNames))
	filteredDocIDsThread := make([]*sroar.Bitmap, len(propertyNames))

	for i, propName := range propertyNames {
		i := i
		propName := propName

		eg.Go(
			func() error {
				bucket := b.store.Bucket(helpers.BucketSearchableFromPropNameLSM(propName))
				if bucket == nil {
					return fmt.Errorf("could not find bucket for property %v", propName)
				}
				preM, err := bucket.MapList(ctx, []byte(query))
				if err != nil {
					return err
				}

				var m []lsmkv.MapPair
				if filterDocIds != nil {
					if filteredDocIDsThread[i] == nil {
						filteredDocIDsThread[i] = sroar.NewBitmap()
					}
					m = make([]lsmkv.MapPair, 0, len(preM))
					for _, val := range preM {
						docID := binary.BigEndian.Uint64(val.Key)
						if filterDocIds.Contains(docID) {
							m = append(m, val)
						} else {
							filteredDocIDsThread[i].Set(docID)
						}
					}
				} else {
					m = preM
				}
				allMsAndProps[i] = MapPairsAndPropName{MapPairs: m, propBoost: propertyBoosts[propName]}
				return nil
			},
		)
	}

	if err := eg.Wait(); err != nil {
		return termResult, nil, err
	}

	filteredDocIDs := sroar.NewBitmap() // to build the global n if there is a filter
	if filterDocIds != nil {
		for _, docIDs := range filteredDocIDsThread {
			if docIDs != nil {
				filteredDocIDs.Or(docIDs)
			}
		}
	}

	propetiesWithResults := 0
	largestM := 0
	for _, m := range allMsAndProps {
		if len(m.MapPairs) > largestM {
			largestM = len(m.MapPairs)
		}
		if len(m.MapPairs) > 0 {
			propetiesWithResults++
		}
	}

	if propetiesWithResults == 0 {
		termResult.exhausted = true
		termResult.idPointer = math.MaxUint64
		return termResult, nil, nil
	}

	// only include indices for the last element if there are multiple properties
	includeIndicesForLastElement := propetiesWithResults > 1

	var docMapPairs []terms.DocPointerWithScore = nil
	var docMapPairsIndices map[uint64]int = nil

	indices := make([]int, len(allMsAndProps))

	for {
		i := -1
		minKey := make([]byte, 8)
		for ti, mAndProps := range allMsAndProps {
			if indices[ti] >= len(mAndProps.MapPairs) {
				continue
			}
			ki := mAndProps.MapPairs[indices[ti]].Key
			if i == -1 || bytes.Compare(ki, minKey) < 0 {
				i = ti
				copy(minKey, ki)
			}
		}

		if i == -1 {
			break
		}

		m := allMsAndProps[i].MapPairs
		val := m[indices[i]]
		key := binary.BigEndian.Uint64(val.Key)

		indices[i]++

		propBoost := allMsAndProps[i].propBoost

		// only create maps/slices if we know how many entries there are
		if docMapPairs == nil {
			docMapPairs = make([]terms.DocPointerWithScore, 0, len(m))
			docMapPairsIndices = make(map[uint64]int, len(m))
			if len(val.Value) < 8 {
				b.logger.Warnf("Skipping pair in BM25: MapPair.Value should be 8 bytes long, but is %d.", len(val.Value))
				continue
			}
			key := binary.BigEndian.Uint64(val.Key)
			freqBits := binary.LittleEndian.Uint32(val.Value[0:4])
			propLenBits := binary.LittleEndian.Uint32(val.Value[4:8])

			docMapPairs = append(docMapPairs,
				terms.DocPointerWithScore{
					Id:         key,
					Frequency:  math.Float32frombits(freqBits) * propBoost,
					PropLength: math.Float32frombits(propLenBits),
				})
			if includeIndicesForLastElement {
				docMapPairsIndices[key] = len(docMapPairs) - 1 // current last entry
			}

		} else {
			if len(val.Value) < 8 {
				b.logger.Warnf("Skipping pair in BM25: MapPair.Value should be 8 bytes long, but is %d.", len(val.Value))
				continue
			}
			ind, ok := docMapPairsIndices[key]
			freqBits := binary.LittleEndian.Uint32(val.Value[0:4])
			propLenBits := binary.LittleEndian.Uint32(val.Value[4:8])
			if ok {
				if ind >= len(docMapPairs) {
					// the index is not valid anymore, but the key is still in the map
					b.logger.Warnf("Skipping pair in BM25: Index %d is out of range for key %d, length %d.", ind, key, len(docMapPairs))
					continue
				}
				if ind < len(docMapPairs) && docMapPairs[ind].Id != key {
					b.logger.Warnf("Skipping pair in BM25: id at %d in doc map pairs, %d, differs from current key, %d", ind, docMapPairs[ind].Id, key)
					continue
				}

				docMapPairs[ind].PropLength += math.Float32frombits(propLenBits)
				docMapPairs[ind].Frequency += math.Float32frombits(freqBits) * propBoost
			} else {
				docMapPairs = append(docMapPairs,
					terms.DocPointerWithScore{
						Id:         binary.BigEndian.Uint64(val.Key),
						Frequency:  math.Float32frombits(freqBits) * propBoost,
						PropLength: math.Float32frombits(propLenBits),
					})
				if includeIndicesForLastElement {
					docMapPairsIndices[binary.BigEndian.Uint64(val.Key)] = len(docMapPairs) - 1
				}
			}

		}
	}
	if docMapPairs == nil {
		termResult.exhausted = true
		termResult.idPointer = math.MaxUint64
		return termResult, docMapPairsIndices, nil
	}

	termResult.data = docMapPairs

	n := float64(len(docMapPairs))
	if filterDocIds != nil {
		n += float64(filteredDocIDs.GetCardinality())
	}
	termResult.idf = math.Log(float64(1)+(N-n+0.5)/(n+0.5)) * float64(duplicateTextBoost)

	// catch special case where there are no results and would panic termResult.data[0].id
	// related to #4125
	if len(termResult.data) == 0 {
		termResult.posPointer = 0
		termResult.idPointer = math.MaxUint64
		termResult.exhausted = true
		return termResult, docMapPairsIndices, nil
	}

	termResult.posPointer = 0
	termResult.idPointer = termResult.data[0].Id
	return termResult, docMapPairsIndices, nil
}

type termMem struct {
	// doubles as max impact (with tf=1, the max impact would be 1*idf), if there
	// is a boost for a queryTerm, simply apply it here once
	idf float64

	idPointer  uint64
	posPointer uint64
	data       []terms.DocPointerWithScore
	exhausted  bool
	queryTerm  string
	index      int
}

func (t *termMem) ScoreAndAdvance(averagePropLength float64, config schema.BM25Config) (uint64, float64, terms.DocPointerWithScore) {
	id := t.idPointer
	pair := t.data[t.posPointer]
	freq := float64(pair.Frequency)
	tf := freq / (freq + config.K1*(1-config.B+config.B*float64(pair.PropLength)/averagePropLength))

	// advance
	t.posPointer++
	if t.posPointer >= uint64(len(t.data)) {
		t.idPointer = math.MaxUint64
		t.exhausted = true
	} else {
		t.idPointer = t.data[t.posPointer].Id
	}

	return id, tf * t.idf, pair
}

func (t *termMem) AdvanceAtLeast(minID uint64) {
	for t.idPointer < minID {
		t.posPointer++
		if t.posPointer >= uint64(len(t.data)) {
			t.idPointer = math.MaxUint64
			t.exhausted = true
			return
		}
		t.idPointer = t.data[t.posPointer].Id
	}
}

func (t *termMem) IsExhausted() bool {
	return t.exhausted
}

func (t *termMem) IdPointer() uint64 {
	return t.idPointer
}

func (t *termMem) QueryTerm() string {
	return t.queryTerm
}

func (t *termMem) IDF() float64 {
	return t.idf
}

func (t *termMem) QueryTermIndex() int {
	return t.index
}

type MapPairsAndPropName struct {
	propBoost float32
	MapPairs  []lsmkv.MapPair
}

type AllMapPairsAndPropName []MapPairsAndPropName

// provide sort interface
func (m AllMapPairsAndPropName) Len() int {
	return len(m)
}

func (m AllMapPairsAndPropName) Less(i, j int) bool {
	return len(m[i].MapPairs) < len(m[j].MapPairs)
}

func (m AllMapPairsAndPropName) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func PropertyHasSearchableIndex(class *models.Class, tentativePropertyName string) bool {
	if class == nil {
		return false
	}

	propertyName := strings.Split(tentativePropertyName, "^")[0]
	p, err := schema.GetPropertyByName(class, propertyName)
	if err != nil {
		return false
	}
	return HasSearchableIndex(p)
}
