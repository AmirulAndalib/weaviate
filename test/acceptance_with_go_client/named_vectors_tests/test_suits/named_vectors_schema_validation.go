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

package test_suits

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	wvt "github.com/weaviate/weaviate-go-client/v5/weaviate"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
)

func testSchemaValidation(host string) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		client, err := wvt.NewClient(wvt.Config{Scheme: "http", Host: host})
		require.Nil(t, err)

		cleanup := func() {
			err := client.Schema().AllDeleter().Do(context.Background())
			require.Nil(t, err)
		}

		t.Run("none existent name of the vectorizer module", func(t *testing.T) {
			cleanup()
			className := "NamedVector"
			nonExistenVectorizer := "non_existent_vectorizer_module"
			class := &models.Class{
				Class: className,
				Properties: []*models.Property{
					{
						Name: "text", DataType: []string{schema.DataTypeText.String()},
					},
				},
				VectorConfig: map[string]models.VectorConfig{
					nonExistenVectorizer: {
						Vectorizer: map[string]interface{}{
							nonExistenVectorizer: map[string]interface{}{
								"vectorizeClassName": false,
							},
						},
						VectorIndexType: "hnsw",
					},
				},
			}

			err := client.Schema().ClassCreator().WithClass(class).Do(ctx)
			require.Error(t, err)
			assert.ErrorContains(t, err, "vectorizer: no module with name \\\"non_existent_vectorizer_module\\\" present")
		})

		t.Run("properties check", func(t *testing.T) {
			cleanup()
			className := "NamedVector"
			class := &models.Class{
				Class: className,
				Properties: []*models.Property{
					{
						Name: "text", DataType: []string{schema.DataTypeText.String()},
					},
				},
				VectorConfig: map[string]models.VectorConfig{
					c11y: {
						Vectorizer: map[string]interface{}{
							text2vecContextionary: map[string]interface{}{
								"vectorizeClassName": false,
								"properties":         []interface{}{"text", 1111},
							},
						},
						VectorIndexType: "hnsw",
					},
				},
			}

			err := client.Schema().ClassCreator().WithClass(class).Do(ctx)
			require.Error(t, err)
			assert.ErrorContains(t, err, "properties field value: 1111 must be a string")
		})

		t.Run("very long target vector name", func(t *testing.T) {
			tests := []struct {
				targetVectorName   string
				validationErrorMsg string
			}{
				{
					targetVectorName:   fmt.Sprintf("%s1", transformers_bq_very_long_230_chars),
					validationErrorMsg: "Target vector name should not be longer than 230 characters",
				},
				{
					targetVectorName:   "invalid-characters",
					validationErrorMsg: "Weaviate target vector names are restricted to valid GraphQL names",
				},
			}
			for _, tt := range tests {
				cleanup()
				class := &models.Class{
					Class: className,
					Properties: []*models.Property{
						{
							Name: "text", DataType: []string{schema.DataTypeText.String()},
						},
					},
					VectorConfig: map[string]models.VectorConfig{
						tt.targetVectorName: {
							Vectorizer: map[string]interface{}{
								text2vecContextionary: map[string]interface{}{
									"vectorizeClassName": false,
								},
							},
							VectorIndexType:   "flat",
							VectorIndexConfig: bqFlatIndexConfig(),
						},
					},
				}

				err := client.Schema().ClassCreator().WithClass(class).Do(ctx)
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.validationErrorMsg)
			}
		})

		t.Run("default vector", func(t *testing.T) {
			cleanup()
			oneTargetVector := "oneTargetVector"
			class := &models.Class{
				Class: className,
				Properties: []*models.Property{
					{
						Name: "text", DataType: []string{schema.DataTypeText.String()},
					},
				},
				VectorConfig: map[string]models.VectorConfig{
					oneTargetVector: {
						Vectorizer: map[string]interface{}{
							text2vecContextionary: map[string]interface{}{
								"vectorizeClassName": false,
							},
						},
						VectorIndexType:   "flat",
						VectorIndexConfig: bqFlatIndexConfig(),
					},
				},
			}

			err := client.Schema().ClassCreator().WithClass(class).Do(ctx)
			require.NoError(t, err)

			objWrapper, err := client.Data().Creator().
				WithClassName(className).
				WithID(id1).
				WithProperties(map[string]interface{}{
					"text": "default target vector",
				}).
				Do(ctx)
			require.NoError(t, err)
			require.NotNil(t, objWrapper)
			assert.Len(t, objWrapper.Object.Vectors, 1)

			nearTextWithoutTargetVector := client.GraphQL().NearTextArgBuilder().
				WithConcepts([]string{"book"})

			resultVectors := getVectorsWithNearText(t, client, className, id1, nearTextWithoutTargetVector, oneTargetVector)
			assert.Len(t, resultVectors, 1)
		})

		t.Run("generative module wrong configuration - legacy configuration", func(t *testing.T) {
			class := &models.Class{
				Class: "GenerativeOpenAIModuleLegacyValidation",
				Properties: []*models.Property{
					{
						Name:     "text",
						DataType: []string{schema.DataTypeText.String()},
					},
				},
				ModuleConfig: map[string]interface{}{
					"generative-openai": map[string]interface{}{
						"model": "wrong-model",
					},
				},
				Vectorizer:      text2vecContextionary,
				VectorIndexType: "hnsw",
			}
			err := client.Schema().ClassCreator().WithClass(class).Do(ctx)
			require.Error(t, err)
			assert.ErrorContains(t, err, "wrong OpenAI model name")
		})

		t.Run("generative module wrong configuration - multiple vectors", func(t *testing.T) {
			class := &models.Class{
				Class: "GenerativeOpenAIModuleValidation",
				Properties: []*models.Property{
					{
						Name:     "text",
						DataType: []string{schema.DataTypeText.String()},
					},
				},
				VectorConfig: map[string]models.VectorConfig{
					c11y: {
						Vectorizer: map[string]interface{}{
							text2vecContextionary: map[string]interface{}{
								"vectorizeClassName": false,
							},
						},
						VectorIndexType: "hnsw",
					},
					transformers_flat: {
						Vectorizer: map[string]interface{}{
							text2vecTransformers: map[string]interface{}{
								"vectorizeClassName": false,
							},
						},
						VectorIndexType: "flat",
					},
				},
				ModuleConfig: map[string]interface{}{
					"generative-openai": map[string]interface{}{
						"model": "wrong-model",
					},
				},
			}
			err := client.Schema().ClassCreator().WithClass(class).Do(ctx)
			require.Error(t, err)
			assert.ErrorContains(t, err, "wrong OpenAI model name")
		})

		t.Run("generative module proper configuration - multiple vectors", func(t *testing.T) {
			class := &models.Class{
				Class: "GenerativeOpenAIModuleValidationProperConfig",
				Properties: []*models.Property{
					{
						Name:     "text",
						DataType: []string{schema.DataTypeText.String()},
					},
				},
				VectorConfig: map[string]models.VectorConfig{
					c11y: {
						Vectorizer: map[string]interface{}{
							text2vecContextionary: map[string]interface{}{
								"vectorizeClassName": false,
							},
						},
						VectorIndexType: "hnsw",
					},
					transformers_flat: {
						Vectorizer: map[string]interface{}{
							text2vecTransformers: map[string]interface{}{
								"vectorizeClassName": false,
							},
						},
						VectorIndexType: "flat",
					},
				},
				ModuleConfig: map[string]interface{}{
					"generative-openai": map[string]interface{}{
						"model": "gpt-4",
					},
				},
			}
			err := client.Schema().ClassCreator().WithClass(class).Do(ctx)
			require.NoError(t, err)
		})
	}
}
