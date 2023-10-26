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

package schema

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/raft"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	command "github.com/weaviate/weaviate/cloud/proto/cluster"
	"github.com/weaviate/weaviate/cloud/store"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
	"github.com/weaviate/weaviate/usecases/config"
	"github.com/weaviate/weaviate/usecases/scaler"
	"github.com/weaviate/weaviate/usecases/sharding"
)

func newTestHandler(t *testing.T, db store.DB) (*Handler, func() raft.Future) {
	cfg := config.Config{}
	cluster, writer, clusterstate := startRaftCluster(t)
	writer.SetDB(db)
	reader := writer.SchemaReader()
	logger, _ := test.NewNullLogger()
	handler, err := NewHandler(
		writer, reader, &fakeValidator{}, logger, &fakeAuthorizer{nil},
		cfg, dummyParseVectorConfig, &fakeVectorizerValidator{}, dummyValidateInvertedConfig,
		&fakeModuleConfig{}, clusterstate, &fakeScaleOutManager{})
	require.Nil(t, err)
	return &handler, cluster.Shutdown
}

func startRaftCluster(t *testing.T) (*store.Cluster, *store.Store, clusterState) {
	node := "node-1"
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {}),
	)
	url := srv.URL
	srv.Close()
	candidates := []store.Candidate{
		{
			ID:       node,
			Address:  url,
			NonVoter: false,
		},
	}
	urlSplit := strings.Split(url, ":")
	host := strings.TrimPrefix(urlSplit[1], "//")

	root := t.TempDir()
	clusterstate := newFakeClusterState()

	writer := store.New(store.Config{
		WorkDir:              root,
		NodeID:               node,
		Host:                 host,
		RaftPort:             config.DefaultRaftPort,
		RaftElectionTimeout:  500 * time.Millisecond,
		RaftHeartbeatTimeout: 500 * time.Millisecond,
		Parser:               NewParser(clusterstate, dummyParseVectorConfig),
	})

	raftNode, err := writer.Open(false, candidates)
	require.Nil(t, err)
	cluster := store.NewCluster(raftNode, strings.TrimPrefix(url, "http://"))
	if err = cluster.Open(); err != nil {
		fmt.Printf("cluster.open %v", err.Error())
		os.Exit(1)
	}
	// Allow time to elect leader
	time.Sleep(time.Second)
	return &cluster, &writer, clusterstate
}

type fakeDB struct {
	mock.Mock
}

func (f *fakeDB) AddClass(cmd command.AddClassRequest) error {
	return nil
}

func (f *fakeDB) UpdateClass(cmd command.UpdateClassRequest) error {
	return nil
}

func (f *fakeDB) DeleteClass(class string) error {
	return nil
}

func (f *fakeDB) AddProperty(prop string, cmd command.AddPropertyRequest) error {
	return nil
}

func (f *fakeDB) AddTenants(class string, cmd *command.AddTenantsRequest) error {
	return nil
}
func (f *fakeDB) UpdateTenants(class string, cmd *command.UpdateTenantsRequest) error {
	return nil
}

func (f *fakeDB) DeleteTenants(class string, cmd *command.DeleteTenantsRequest) error {
	return nil
}

func (f *fakeDB) UpdateShardStatus(cmd *command.UpdateShardStatusRequest) error {
	return nil
}

func (f *fakeDB) GetShardsStatus(class string) (models.ShardStatusList, error) {
	args := f.Called(class)
	return args.Get(0).(models.ShardStatusList), nil
}

type fakeAuthorizer struct {
	err error
}

func (f *fakeAuthorizer) Authorize(principal *models.Principal, verb, resource string) error {
	return f.err
}

type fakeScaleOutManager struct{}

func (f *fakeScaleOutManager) Scale(ctx context.Context,
	className string, updated sharding.Config, _, _ int64,
) (*sharding.State, error) {
	return nil, nil
}

func (f *fakeScaleOutManager) SetSchemaManager(sm scaler.SchemaManager) {
}

type fakeShardReader struct {
	mock.Mock
}

func (f *fakeShardReader) GetShardsStatus(class string) (models.ShardStatusList, error) {
	args := f.Called(class)
	return args.Get(0).(models.ShardStatusList), nil
}

type fakeValidator struct{}

func (f *fakeValidator) ValidateVectorIndexConfigUpdate(ctx context.Context,
	old, updated schema.VectorIndexConfig,
) error {
	return nil
}

func (f *fakeValidator) ValidateInvertedIndexConfigUpdate(ctx context.Context,
	old, updated *models.InvertedIndexConfig,
) error {
	return nil
}

type fakeModuleConfig struct{}

func (f *fakeModuleConfig) SetClassDefaults(class *models.Class) {
	defaultConfig := map[string]interface{}{
		"my-module1": map[string]interface{}{
			"my-setting": "default-value",
		},
	}

	asMap, ok := class.ModuleConfig.(map[string]interface{})
	if !ok {
		class.ModuleConfig = defaultConfig
		return
	}

	module, ok := asMap["my-module1"]
	if !ok {
		class.ModuleConfig = defaultConfig
		return
	}

	asMap, ok = module.(map[string]interface{})
	if !ok {
		class.ModuleConfig = defaultConfig
		return
	}

	if _, ok := asMap["my-setting"]; !ok {
		asMap["my-setting"] = "default-value"
		defaultConfig["my-module1"] = asMap
		class.ModuleConfig = defaultConfig
	}
}

func (f *fakeModuleConfig) SetSinglePropertyDefaults(class *models.Class,
	prop *models.Property,
) {
}

func (f *fakeModuleConfig) ValidateClass(ctx context.Context, class *models.Class) error {
	return nil
}

type fakeVectorizerValidator struct {
	valid []string
}

func (f *fakeVectorizerValidator) ValidateVectorizer(moduleName string) error {
	for _, valid := range f.valid {
		if moduleName == valid {
			return nil
		}
	}

	return fmt.Errorf("invalid vectorizer %q", moduleName)
}

type fakeClusterState struct {
	hosts       []string
	syncIgnored bool
	skipRepair  bool
}

func newFakeClusterState(hosts ...string) *fakeClusterState {
	return &fakeClusterState{
		hosts: func() []string {
			if len(hosts) == 0 {
				return []string{"node-1"}
			}
			return hosts
		}(),
	}
}

func (f *fakeClusterState) SchemaSyncIgnored() bool {
	return f.syncIgnored
}

func (f *fakeClusterState) SkipSchemaRepair() bool {
	return f.skipRepair
}

func (f *fakeClusterState) Hostnames() []string {
	return f.hosts
}

func (f *fakeClusterState) AllNames() []string {
	return f.hosts
}

func (f *fakeClusterState) Candidates() []string {
	return f.hosts
}

func (f *fakeClusterState) LocalName() string {
	return "node1"
}

func (f *fakeClusterState) NodeCount() int {
	return 1
}

func (f *fakeClusterState) ClusterHealthScore() int {
	return 0
}

func (f *fakeClusterState) ResolveParentNodes(string, string,
) (map[string]string, error) {
	return nil, nil
}

func (f *fakeClusterState) NodeHostname(string) (string, bool) {
	return "", false
}

type fakeVectorConfig struct {
	raw interface{}
}

func (f fakeVectorConfig) IndexType() string {
	return "fake"
}

func dummyParseVectorConfig(in interface{}) (schema.VectorIndexConfig, error) {
	return fakeVectorConfig{raw: in}, nil
}

func dummyValidateInvertedConfig(in *models.InvertedIndexConfig) error {
	return nil
}
