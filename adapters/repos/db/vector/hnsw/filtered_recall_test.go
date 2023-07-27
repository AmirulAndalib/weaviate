//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

//go:build benchmarkRecall
// +build benchmarkRecall

package hnsw

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"sync"
	"testing"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/hnsw/distancer"
	"github.com/weaviate/weaviate/entities/cyclemanager"
	ent "github.com/weaviate/weaviate/entities/vectorindex/hnsw"
)

//"github.com/weaviate/weaviate/adapters/repos/db/helpers"

func init() {
	go func() {
		runtime.SetBlockProfileRate(1)
		http.ListenAndServe("localhost:6060", nil)
	}()
}

func TestFilteredRecall(t *testing.T) {
	efConstruction := 256
	ef := 256
	maxNeighbors := 64

	type Vector struct {
		ID     int   `json:"id"`
		Vector []int `json:"vector"`
	}

	type Filters struct {
		ID        int         `json:"id"`
		FilterMap map[int]int `json:"filterMap"`
	}

	type vecWithFilters struct {
		ID        int         `json:"id"`
		Vector    []int       `json:"vector"`
		FilterMap map[int]int `json:"filterMap"`
	}

	var indexVectors []Vector
	var indexFilters []Filters
	var queryVectors []Vectors
	var queryFilters []Filters
	var truths [][]uint64
	var vectorIndex *hnsw

	t.Run("Loading vectors for testing...", func(t *testing.T) {
		// vectors.json
		/*
			[{"id": 0, "vector": [0,15,35,...]},{"id": 0, "vector": [119,15,4,...]},...]
		*/
		indexVectorsJSON, err := ioutil.ReadFile("indexVectors.json")
		require.Nil(t, err)
		err = json.Unmarshal(vectorsJSON, &indexVectors)
		require.Nil(t, err)
		// vectorFilters.json
		/*
			[{"id": 0, "filterMap": {0: 1, 1: 3, ...}}, {"id": 1, "filterMap": {0: 2, 1: 4}}, ...]
		
			For now, running one test at a time, future - loop through filter paths
		*/
		indexFiltersJSON, err := ioutil.ReadFile("indexFilters.json")
		require.Nil(t, err)
		err = json.Unmarshal(indexFiltersJSON, &filters)
		require.Nil(t, err)

		indexVectorsWithFilters := mergeData(indexVectorsJSON, indexFiltersJSON) // returns []vecWithFilters
		
		/* =================================================
			SAME JOINING OF VECTORS AND FILTERS FOR QUERIES
		   =================================================
		*/

		queryVectorsJSON, err := ioutil.ReadFile("queryVectors.json")
		require.Nil(t, err)
		err = json.Unmarshal(queryJSON, &queries)
		require.Nil(t, err)
		
		queryFiltersJSON, err := ioutil.ReadFile("queryFilters.json")
		require.Nil(t, err)
		err = json.Unmarshal(queryFiltersJSON, &filters)

		queryVectorsWithFilters := mergeData(queryVectorsJSON, queryFiltersJSON)

		truthsJSON, err := ioutil.ReadFile("filtered_recall_truths.json")
		require.Nil(t, err)
		err = json.Unmarshal(truthsJSON, &truths)
		require.Nil(t, err)
	})

	t.Run("importing into hnsw", func(t *testing.T) {
		fmt.Printf("importing into hnsw\n")

		index, err := New(Config{
			RootPath:              "doesnt-matter-as-committlogger-is-mocked-out",
			ID:                    "recallbenchmark",
			MakeCommitLoggerThunk: MakeNoopCommitLogger,
			DistanceProvider:      distancer.NewCosineDistanceProvider(),
			VectorForIDThunk: func(ctx context.Context, id uint64) ([]float32, error) {
				return vectors[int(id)].Vector, nil
			},
		}, ent.UserConfig{
			MaxConnections: maxNeighbors,
			EFConstruction: efConstruction,
			EF:             ef,
		}, cyclemanager.NewNoop())

		filterToIDs := make(map[int][]uint64)

		require.Nil(t, err)
		vectorIndex = index

		workerCount := runtime.GOMAXPROCS(0)
		jobsForWorker := make([][]vecWithFilters, workerCount)

		before := time.Now()
		for i, vecWithFilters := range indexVectorsWithFilters {
			workerID := i % workerCount
			jobsForWorker[workerID] = append(jobsForWorker[workerID], vecWithFilters)
		}

		wg := &sync.WaitGroup{}
		mutex := &sync.Mutex{}
		for workerID, jobs := range jobsForWorker {
			wg.Add(1)
			go func(workerID int, myJobs []LabeledVector) {
				defer wg.Done()
				for i, vec := range myJobs {
					originalIndex := (i * workerCount) + workerID
					nodeId := uint64(originalIndex)
					err := vectorIndex.filteredAdd(nodeId, vec.Vector, vec.FilterMap) // change signature to add vec.Label
					require.Nil(t, err)
					mutex.Lock()
					if _, ok := filterToIDs[vec.Label]; !ok {
						filterToIDs[vec.Label] = []uint64{nodeId}
					} else {
						filterToIDs[vec.Label] = append(filterToIDs[vec.Label], nodeId)
					}
					mutex.Unlock()
					require.Nil(t, err)
				}
			}(workerID, jobs)
		}

		wg.Wait()
		fmt.Printf("importing took %s\n", time.Since(before))

		fmt.Printf("With k=20")

		k := 100

		var relevant_retrieved int
		var recall float32

		for i := 0; i < len(queries); i++ {
			// change to queryFilters
			queryFilter := queries[i].filterMap
			//construct an allowList from the []uint64 of ids that match the filter
			queryAllowList := helpers.NewAllowList(filterToIDs[queryFilter]...)
			// ok hard to test just searchLayerByVector
			/*
				h.searchLayerByVectorWithDistancer(searchVec, eps, 1, level, 0, false, allowList, byteDistancer)
			*/
			results, _, err := vectorIndex.SearchByVector(queries[i].Vector, k, queryFilterMap, queryAllowList)
			//results, _, err := vectorIndex.SearchByVector(queries[i].Vector, k, nil)
			// it shouldn't matter if it has the allowList or not
			// ^ because it's only connected to nodes that share the same filter
			// confirmed, same recall #

			require.Nil(t, err)

			relevant_retrieved = matchesInLists(truths[i], results)
			fmt.Print(float32(relevant_retrieved) / 100)
			fmt.Print("\n")
			recall += float32(relevant_retrieved) / 100

			// Would I want to see a histogram of recalls per query?

			fmt.Print("LABEL \n")
			fmt.Print(queries[i].Label)
			fmt.Print("\n RESULTS \n")
			fmt.Print(results)
			fmt.Print("\n TRUTHS \n")
			fmt.Print(truths[i])
			fmt.Print("\n")
		}

		recall = float32(recall) / float32(len(queries))
		fmt.Printf("recall is %f\n", recall)
		assert.True(t, recall >= 0.09)
	})
}

func loadVectors(filename string) []Vector {
	var vectors []Vector

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &vectors)
	if err != nil {
		log.Fatal(err)
	}

	return vectors
}

func loadFilters(filename string) []Filter {
	var filters []Filters

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &filters)
	if err != nil {
		log.Fatal(err)
	}

	return filters
}

func mergeData(vectors []Vector, filters []Filter) []Result {
	// Create a map for quick lookup of filters
	IDtoFilterMap := make(map[int]map[int]int)
	for _, filter := range filters {
		IDtoFilterMap[filter.ID] = filter.FilterMap
	}
	// Merge vectors and filters
	var results []vecWithFilters
	for _, vector := range vectors {
		result := vecWithFilters{
			ID:        vector.ID,
			Vector:    vector.Vector,
			FilterMap: IDtoFilterMap[vector.ID]
		}
		results = append(results, result)
	}

	return results
}

func matchesInLists(control []uint64, results []uint64) int {
	desired := map[uint64]struct{}{}
	for _, relevant := range control {
		desired[relevant] = struct{}{}
	}

	var matches int
	for _, candidate := range results {
		_, ok := desired[candidate]
		if ok {
			matches++
		}
	}

	return matches
}
