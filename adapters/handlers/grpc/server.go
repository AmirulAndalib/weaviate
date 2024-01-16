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

package grpc

import (
	"fmt"
	"math"
	"net"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/adapters/handlers/rest/state"
	pbv0 "github.com/weaviate/weaviate/grpc/generated/protocol/v0"
	pbv1 "github.com/weaviate/weaviate/grpc/generated/protocol/v1"
	"github.com/weaviate/weaviate/usecases/auth/authentication/composer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
	"google.golang.org/grpc/health/grpc_health_v1"

	v0 "github.com/weaviate/weaviate/adapters/handlers/grpc/v0"
	v1 "github.com/weaviate/weaviate/adapters/handlers/grpc/v1"
)

const (
	maxRecvMsgSize = 1024 * 1024 * 100 // 100mb, needs to be synchronized with clients
	maxSendMsgSize = math.MaxInt32     // 2gb, needs to be synchronized with clients
)

func CreateGRPCServer(state *state.State) *GRPCServer {
	o := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(parseFromEnv("GRPC_MAX_RECV_MSG_SIZE", maxRecvMsgSize, state.Logger)),
		grpc.MaxSendMsgSize(parseFromEnv("GRPC_MAX_SEND_MSG_SIZE", maxSendMsgSize, state.Logger)),
	}

	// Add TLS creds for the GRPC connection, if defined.
	if len(state.ServerConfig.Config.GRPC.CertFile) > 0 || len(state.ServerConfig.Config.GRPC.KeyFile) > 0 {
		c, err := credentials.NewServerTLSFromFile(state.ServerConfig.Config.GRPC.CertFile,
			state.ServerConfig.Config.GRPC.KeyFile)
		if err != nil {
			state.Logger.WithField("action", "grpc_startup").
				Fatalf("grpc server TLS credential error: %s", err)
		}
		o = append(o, grpc.Creds(c))
	}

	s := grpc.NewServer(o...)
	weaviateV0 := v0.NewService()
	weaviateV1 := v1.NewService(
		state.Traverser,
		composer.New(
			state.ServerConfig.Config.Authentication,
			state.APIKey, state.OIDC),
		state.ServerConfig.Config.Authentication.AnonymousAccess.Enabled,
		state.SchemaManager,
		state.BatchManager,
	)
	pbv0.RegisterWeaviateServer(s, weaviateV0)
	pbv1.RegisterWeaviateServer(s, weaviateV1)
	grpc_health_v1.RegisterHealthServer(s, weaviateV1)

	return &GRPCServer{s}
}

func StartAndListen(s *GRPCServer, state *state.State) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d",
		state.ServerConfig.Config.GRPC.Port))
	if err != nil {
		return err
	}
	state.Logger.WithField("action", "grpc_startup").
		Infof("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func parseFromEnv(name string, defaultValue int, logger *logrus.Logger) int {
	parsedValue, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		logger.WithField("action", "grpc_startup").
			Warnf("could not parse env var: %v, with value: %v. Using default instead: %v", name, parsedValue, defaultValue)
		return defaultValue
	}
	return parsedValue
}

type GRPCServer struct {
	*grpc.Server
}
