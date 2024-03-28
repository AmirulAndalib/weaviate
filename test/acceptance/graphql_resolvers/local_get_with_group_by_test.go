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

package test

import (
	"fmt"
	"testing"

	graphqlhelper "github.com/weaviate/weaviate/test/helper/graphql"
	"github.com/weaviate/weaviate/test/helper/journey"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/test/helper"
)

func groupByObjects(t *testing.T) {
	t.Run("group by: people by city", func(t *testing.T) {
		getGroup := func(value interface{}) map[string]interface{} {
			group := value.(map[string]interface{})["_additional"].(map[string]interface{})["group"].(map[string]interface{})
			return group
		}
		getGroupHits := func(group map[string]interface{}) (string, []string) {
			result := []string{}
			hits := group["hits"].([]interface{})
			for _, hit := range hits {
				additional := hit.(map[string]interface{})["_additional"].(map[string]interface{})
				result = append(result, additional["id"].(string))
			}
			groupedBy := group["groupedBy"].(map[string]interface{})
			groupedByValue := groupedBy["value"].(string)
			return groupedByValue, result
		}
		query := `
		{
			Get{
				Person(
					nearObject:{
						id: "8615585a-2960-482d-b19d-8bee98ade52c"
					}
					groupBy:{
						path:["livesIn"]
						groups:4
						objectsPerGroup: 10
					}
				){
					_additional{
						id
						group{
							groupedBy{value}
							count
							maxDistance
							minDistance
							hits {
								_additional {
									id
									distance
								}
							}
						}
					}
				}
			}
		}
		`
		result := graphqlhelper.AssertGraphQL(t, helper.RootAuth, query)
		groups := result.Get("Get", "Person").AsSlice()

		require.Len(t, groups, 4)

		expectedResults := map[string][]string{}

		groupedBy1 := `weaviate://localhost/City/8f5f8e44-d348-459c-88b1-c1a44bb8f8be`
		expectedGroup1 := []string{
			"8615585a-2960-482d-b19d-8bee98ade52c",
			"3ef44474-b5e5-455d-91dc-d917b5b76165",
			"15d222c9-8c36-464b-bedb-113faa1c1e4c",
		}
		expectedResults[groupedBy1] = expectedGroup1

		groupedBy2 := `weaviate://localhost/City/9b9cbea5-e87e-4cd0-89af-e2f424fd52d6`
		expectedGroup2 := []string{
			"3ef44474-b5e5-455d-91dc-d917b5b76165",
			"15d222c9-8c36-464b-bedb-113faa1c1e4c",
		}
		expectedResults[groupedBy2] = expectedGroup2

		groupedBy3 := `weaviate://localhost/City/6ffb03f8-a853-4ec5-a5d8-302e45aaaf13`
		expectedGroup3 := []string{
			"15d222c9-8c36-464b-bedb-113faa1c1e4c",
		}
		expectedResults[groupedBy3] = expectedGroup3

		groupedBy4 := ""
		expectedGroup4 := []string{
			"5d0fa6ee-21c4-4b46-a735-f0208717837d",
		}
		expectedResults[groupedBy4] = expectedGroup4

		groupsOrder := []string{groupedBy1, groupedBy2, groupedBy4, groupedBy3}
		for i, current := range groups {
			group := getGroup(current)
			groupedBy, ids := getGroupHits(group)
			assert.Equal(t, groupsOrder[i], groupedBy)
			assert.ElementsMatch(t, expectedResults[groupedBy], ids)
		}
	})

	t.Run("group by: passages by documents", func(t *testing.T) {
		journey.GroupBySingleAndMultiShardTests(t, "")
	})
}

func groupByBm25(t *testing.T) {
	t.Run("group by: companies by city bm25", func(t *testing.T) {

		getGroup := func(value interface{}) map[string]interface{} {
			group := value.(map[string]interface{})["_additional"].(map[string]interface{})["group"].(map[string]interface{})
			return group
		}
		getGroupHits := func(group map[string]interface{}) (string, []string) {
			result := []string{}
			hits := group["hits"].([]interface{})
			for _, hit := range hits {
				additional := hit.(map[string]interface{})["_additional"].(map[string]interface{})
				result = append(result, additional["id"].(string))
			}
			groupedBy := group["groupedBy"].(map[string]interface{})
			groupedByValue := groupedBy["value"].(string)
			return groupedByValue, result
		}

		query := `
		{
			Get{
				CompanyGroup(
					bm25:{
						query:"Inc Apple Microsoft"
					}
					groupBy:{
						path:["city"]
						groups:4
						objectsPerGroup: 10
					}
				){
					_additional{
						group{
							id
							groupedBy{value path}
							count
							maxDistance
							minDistance
							hits {
								name city
								_additional {
									id
									distance
								}
							}
						}
					}
				}
			}
		}
		`
		result := graphqlhelper.AssertGraphQL(t, helper.RootAuth, query)
		groups := result.Get("Get", "CompanyGroup").AsSlice()
		fmt.Printf("groups: %+v\n", groups)

		require.Len(t, groups, 3)

		group1 := getGroup(groups[0])
		groupby, hits := getGroupHits(group1)

		t.Logf("groupby: %s, hits: %+v\n", groupby, hits)
		require.Equal(t, "dusseldorf", groupby)
		require.Len(t, hits, 2)
		require.Equal(t, hits[0], "1fa3b21e-ca4f-4db7-a432-7fc6a23c534d")
		require.Equal(t, hits[1], "1b2cfdba-d4ba-4cf8-abda-e719ef35ac33")

		group2 := getGroup(groups[1])
		groupby, hits = getGroupHits(group2)
		t.Logf("groupby: %s, hits: %+v\n", groupby, hits)
		require.Equal(t, "berlin", groupby)
		require.Len(t, hits, 2)
		require.Equal(t, hits[0], "177fec91-1292-4928-8f53-f0ff49c76900")
		require.Equal(t, hits[1], "1343f51d-7e05-4084-bd66-d504db3b6bec")

		group3 := getGroup(groups[2])
		groupby, hits = getGroupHits(group3)
		t.Logf("groupby: %s, hits: %+v\n", groupby, hits)
		require.Equal(t, "amsterdam", groupby)
		require.Len(t, hits, 3)
		require.Equal(t, hits[0], "171d2b4c-3da1-4684-9c5e-aabd2a4f2998")
		require.Equal(t, hits[1], "1f75ed97-39dd-4294-bff7-ecabd7923062")
		require.Equal(t, hits[2], "1c2e21fc-46fe-4999-b41c-a800595129af")

	})

}

func groupByHybridBm25(t *testing.T) {
	t.Run("group by: companies by city hybrid bm25", func(t *testing.T) {

		getGroup := func(value interface{}) map[string]interface{} {
			group := value.(map[string]interface{})["_additional"].(map[string]interface{})["group"].(map[string]interface{})
			return group
		}
		getGroupHits := func(group map[string]interface{}) (string, []string) {
			result := []string{}
			hits := group["hits"].([]interface{})
			for _, hit := range hits {
				additional := hit.(map[string]interface{})["_additional"].(map[string]interface{})
				result = append(result, additional["id"].(string))
			}
			groupedBy := group["groupedBy"].(map[string]interface{})
			groupedByValue := groupedBy["value"].(string)
			return groupedByValue, result
		}

		query := `
		{
			Get{
				CompanyGroup(
					hybrid:{
						query:"Inc Apple Microsoft"
					}
					groupBy:{
						path:["city"]
						groups:4
						objectsPerGroup: 10
					}
				){
					_additional{
						group{
							id
							groupedBy{value path}
							count
							maxDistance
							minDistance
							hits {
								name city
								_additional {
									id
									distance
								}
							}
						}
					}
				}
			}
		}
		`
		result := graphqlhelper.AssertGraphQL(t, helper.RootAuth, query)
		groups := result.Get("Get", "CompanyGroup").AsSlice()
		fmt.Printf("groups: %+v\n", groups)

		require.Len(t, groups, 3)

		group1 := getGroup(groups[2])
		groupby, hits := getGroupHits(group1)

		t.Logf("groupby: %s, hits: %+v\n", groupby, hits)
		require.Equal(t, "dusseldorf", groupby)
		require.Len(t, hits, 3)
		require.Equal(t, hits[0], "1fa3b21e-ca4f-4db7-a432-7fc6a23c534d")
		require.Equal(t, hits[1], "1b2cfdba-d4ba-4cf8-abda-e719ef35ac33")

		group2 := getGroup(groups[0])
		groupby, hits = getGroupHits(group2)
		t.Logf("groupby: %s, hits: %+v\n", groupby, hits)
		require.Equal(t, "berlin", groupby)
		require.Len(t, hits, 3)
		require.Equal(t, hits[0], "177fec91-1292-4928-8f53-f0ff49c76900")
		require.Equal(t, hits[1], "1343f51d-7e05-4084-bd66-d504db3b6bec")

		group3 := getGroup(groups[1])
		groupby, hits = getGroupHits(group3)
		t.Logf("groupby: %s, hits: %+v\n", groupby, hits)
		require.Equal(t, "amsterdam", groupby)
		require.Len(t, hits, 3)
		require.Equal(t, hits[0], "1f75ed97-39dd-4294-bff7-ecabd7923062")
		require.Equal(t, hits[1], "1c2e21fc-46fe-4999-b41c-a800595129af")
		require.Equal(t, hits[2], "171d2b4c-3da1-4684-9c5e-aabd2a4f2998")

	})

}

func groupByHybridNearVector(t *testing.T) {
	defaultLimit := 100
	t.Run("group hybrid nearVector: with implicit unlimited search - no limit provided (with distance)", func(t *testing.T) {
		query := `
		{
			Get {
				RansomNote(
          hybrid: {
            searches: {
					nearVector: {
						distance: 1.8
						vector: [-0.07853702, -0.33730024, 0.62998116, 0.08703484, -0.0011832615, 0.23041481, -0.091878965, 0.1184808, 0.060692377, 0.1748896, 0.53659165, 0.12019079, 0.54373807, -0.43369776, 0.1843199, -0.19319294, 0.122559674, -0.09465141, -0.14069664, 0.031092037, -0.1763922, 0.0074394196, -0.2586067, 0.10990611, -0.18623954, -0.038631044, -0.22795723, 0.09296776, -0.31110525, -0.37963995, -0.19045947, 0.48089907, 0.46725857, 0.28541213, 0.08294283, -0.18865398, 0.09647029, 0.2321466, -0.03435125, -0.09602424, -0.3831683, -0.027315892, 0.4215511, -0.35865632, 0.41955224, 0.090477064, 0.29026023, -0.48485047, -0.24656451, -0.06478625, 0.07755969, -0.049564634, 0.026147474, -0.028342195, -0.035627227, 0.49309397, 0.3705331, 0.04615483, 0.14789912, -0.01220134, 0.300666, -0.246646, 0.0038986988, 0.16730541, 0.46581128, -0.04931062, 0.040290095, 0.32867354, -0.18300997, 0.30411696, -0.1969807, 0.4727539, -0.31915516, -0.32722718, 0.12694982, 0.22583495, -0.014532595, -0.14432396, 0.2469766, 0.14872919, -0.06750808, 0.06351308, -0.287756, -0.32118404, 0.25326216, 0.45288888, -0.36307186, 0.05369787, -0.3283361, 0.07754738, 0.38473788, -0.5120014, -0.3344492, -0.1102767, -0.16755687, -0.3575448, -0.2555015, -0.42976367, -0.2695758, 0.04068499, 0.591914, -0.008395256, 0.2613976, -0.51722556, -0.22581989, 0.036452737, 0.42190477, -0.256124, 0.25849363, -0.073515825, -0.08690646, 0.013338611, 0.14928678, 0.16371651, 0.111465625, -0.117571846, -0.44434816, 0.07603647, 0.4188736, -0.16967061, 0.040313836, 0.41924894, -0.36998197, 0.23204626, -0.23309743, -0.18061559, 1.0674918, -0.51468146, -0.37230963, 0.02214618, -0.5616187, -0.07473461, -0.3314831, -0.24691144, -0.34061712, -0.1511554, 0.33711013, 0.1055847, -0.047220375, -0.06317049, -0.22572862, -0.21646689, 0.090705894, 0.018562902, 0.020744732, -0.5342965, -0.23402104, -0.17680043, 0.1363935, -0.17916845, 0.37825805, -0.07233101, -0.28244817, 0.4055966, 0.19701958, 0.6236174, 0.078134544, 0.46439224, -0.60451704, 0.16722181, -0.20011653, 0.36931068, -0.39967215, 0.21178648, 0.47920865, -0.033521328, 0.57077545, -0.8003054, -0.4028354, 0.27799648, -0.23070334, 0.57747835, 0.49984616, -0.12409506, -0.26694623, -0.20168623, -0.19415514, -0.4626071, 0.10374411, 0.24878122, 0.47352287, -0.6494472, -0.26087105, 0.418008, -0.2789803, -0.60986733, -0.54914564, 0.4734504, 0.04347568, -0.13549352, 0.1530609, 0.085039385, -0.014595425, -0.1106091, 0.014441653, 0.14899726, -0.107090004, 0.03979257, 0.20897605, -0.040235993, 0.1928343, -0.048328623, 0.5435577, -0.1704212, -0.016530415, 0.11402996, 0.24666561, -0.62601864, 0.6729872, -0.21594357, -0.3161654, 0.2899072, -0.05281632, 0.026857251, 0.13927892, 0.26362655, 0.37995058, -0.056429606, 0.27310744, -0.34237143, -0.6419976, -0.02513231, -0.18217334, 0.021232722, -0.35155025, 0.055071075, -0.22192729, 0.4597671, 0.09872845, -0.41803727, -0.08897542, -0.63276047, 0.38059604, 0.45347637, 0.52723855, 0.25096536, -0.3165448, 0.43803728, 0.02419832, 0.317004, -0.059602205, 0.15561013, 0.11867607, 0.7157601, 0.08024589, -0.013107148, 0.3127224, -0.08844044, 0.5374578, 0.39421698, -0.054171022, 0.0913302, -0.081881694, 0.24596375, -0.2841653, -0.5482517, -0.5673938, 0.05889957, -0.1146344, 0.39452744, -0.03414711, 0.32027423, 0.2599335, -0.31470263, 0.45967287, -0.5710101, -0.21222454, 0.38154987, -0.21218868, -0.4366558, 0.13715877, 0.23925674, 0.34832072, -0.03769251, 0.25530148, 0.10662722, -0.5269836, 0.32952255, 0.46165445, 0.3794754, -0.061259665, 0.02883365, -0.3199015, 0.40625557, -0.3794913, 0.42420092, 0.4631467, 0.54236996, 0.031472385, 0.2635622, -0.25566247, -0.040713936, 0.48734123, 0.2742017, -0.15524681, 0.025654443, 0.056942068, -0.48883253, 0.60433495, 0.03514151]
					}
          }
          }
				) {
					_additional {
						vector
					}
					contents
				}
			}
		}
	`
		result := graphqlhelper.AssertGraphQL(t, helper.RootAuth, query)
		notes := result.Get("Get", "RansomNote").AsSlice()
		require.NotEmpty(t, notes)
		require.Equal(t, len(notes), defaultLimit)
	})
}
