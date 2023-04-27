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

package hybrid

import (
	"fmt"
	"sort"

	"github.com/go-openapi/strfmt"
	"github.com/weaviate/weaviate/entities/search"
)

func FusionScoreCombSUM(results [][]search.Result) []search.Result {
	allDocs := map[strfmt.UUID]*search.Result{}
	// Loop over each array of results and add the score of each document to the totals
	totals := map[strfmt.UUID]float32{}
	for _, resultSet := range results {
		for i, doc := range resultSet {
			allDocs[doc.ID] = &resultSet[i]
			score := doc.Score
			if _, ok := totals[doc.ID]; ok {
				totals[doc.ID] = totals[doc.ID] + score
			} else {
				totals[doc.ID] = score
			}

		}
	}

	out := []search.Result{}
	for docID, score := range totals {
		doc := *allDocs[docID]
		doc.Score = score
		out = append(out, doc)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Score > out[j].Score
	})

	return out
}

func FusionNormalize(weights []float64, results [][]*Result) []*Result {
	maximum := []float32{-100000, -100000}
	minimum := []float32{100000, 100000}

	for i := range results {
		for _, res := range results[i] {
			if res.SecondarySortValue > maximum[i] {
				maximum[i] = res.SecondarySortValue
			}

			if res.SecondarySortValue < minimum[i] {
				minimum[i] = res.SecondarySortValue
			}

		}
	}

	// normalize scores
	mapResults := make(map[strfmt.UUID]*Result)

	for i := range results {
		weight := float32(weights[i])
		for _, res := range results[i] {
			tempResult := res
			score := weight * (res.SecondarySortValue - minimum[i]) / (maximum[i] - minimum[i])

			previousResult, ok := mapResults[res.ID]
			if ok {
				score += previousResult.Score
			}
			tempResult.Score = score

			mapResults[res.ID] = tempResult
		}
	}

	concat := make([]*Result, 0, len(mapResults))
	for _, res := range mapResults {
		concat = append(concat, res)
	}

	sort.Slice(concat, func(i, j int) bool {
		a_b := float64(concat[j].Score - concat[i].Score)
		if a_b*a_b < 1e-14 {
			return concat[i].SecondarySortValue > concat[j].SecondarySortValue
		}
		return float64(concat[i].Score) > float64(concat[j].Score)
	})
	return concat
}

func FusionReciprocal(weights []float64, results [][]*Result) []*Result {
	mapResults := map[strfmt.UUID]*Result{}
	for resultSetIndex, result := range results {
		for i, res := range result {
			tempResult := res
			docId := tempResult.ID
			score := weights[resultSetIndex] / float64(i+60+1) // TODO replace 60 with a class configured variable

			if tempResult.AdditionalProperties == nil {
				tempResult.AdditionalProperties = map[string]interface{}{}
			}

			// Get previous results from the map, if any
			previousResult, ok := mapResults[docId]
			if ok {
				tempResult.AdditionalProperties["explainScore"] = fmt.Sprintf(
					"%v\n(hybrid) Document %v contributed %v to the score",
					previousResult.AdditionalProperties["explainScore"], tempResult.ID, score)
				score += float64(previousResult.Score)
			} else {
				tempResult.AdditionalProperties["explainScore"] = fmt.Sprintf(
					"%v\n(hybrid) Document %v contributed %v to the score",
					tempResult.ExplainScore, tempResult.ID, score)
			}
			tempResult.AdditionalProperties["rank_score"] = score
			tempResult.AdditionalProperties["score"] = score

			tempResult.Score = float32(score)
			mapResults[docId] = tempResult
		}
	}

	// Sort the results
	var (
		concat = make([]*Result, len(mapResults))
		i      = 0
	)
	for _, res := range mapResults {
		res.ExplainScore = res.AdditionalProperties["explainScore"].(string)
		concat[i] = res
		i++
	}

	sort.Slice(concat, func(i, j int) bool {
		a_b := float64(concat[j].Score - concat[i].Score)
		if a_b*a_b < 1e-14 {
			return concat[i].SecondarySortValue > concat[j].SecondarySortValue
		}
		return float64(concat[i].Score) > float64(concat[j].Score)
	})
	return concat
}

func FusionScoreConcatenate(results [][]*search.Result) []*search.Result {
	// Concatenate the results
	concatenatedResults := []*search.Result{}
	for _, result := range results {
		concatenatedResults = append(concatenatedResults, result...)
	}

	sort.Slice(concatenatedResults, func(i, j int) bool {
		a := concatenatedResults[i].Score

		b := concatenatedResults[j].Score

		return a > b
	})
	return concatenatedResults
}
