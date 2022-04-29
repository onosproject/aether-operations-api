/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package grpc

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/internal/stores"
	"github.com/onosproject/scaling-umbrella/internal/stores/application"
	"github.com/onosproject/scaling-umbrella/internal/stores/enterprise"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
)

var log = logging.GetLogger("RocApiGrpcServer")

type IRocApiGrpcServer interface {
	RegisterGrpcServer(srv grpc.ServiceRegistrar)
}

type RocApiGrpcServices struct {
	ApplicationService v1.ApplicationServiceServer
	EnterpriseService  v1.EnterpriseServiceServer
}

type RocApiGrpcServer struct {
	doneCh   chan bool
	wg       *sync.WaitGroup
	address  string
	Services *RocApiGrpcServices
	Servers  []IRocApiGrpcServer
}

func (s *RocApiGrpcServer) StartGrpcServer() error {
	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", s.address)
	log.Infow("starting-grpc-server", "address", s.address)
	if err != nil {
		return err
	}

	for _, s := range s.Servers {
		s.RegisterGrpcServer(grpcServer)
	}

	reflection.Register(grpcServer)

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Errorw("grpc-server-error", "err", err)
		}
	}()

	x := <-s.doneCh
	if x {
		// if the API channel is closed, stop the gRPC server
		grpcServer.Stop()
		log.Info("Stopping gRPC server")
	}

	s.wg.Done()

	return nil
}

// NewGrpcServer creates a new gRPC server with handlers for all the services
// defined in the protos
func NewGrpcServer(doneCh chan bool, wg *sync.WaitGroup, address string, s *stores.Stores) *RocApiGrpcServer {

	srv := RocApiGrpcServer{
		doneCh:   doneCh,
		wg:       wg,
		address:  address,
		Services: &RocApiGrpcServices{},
	}

	appServer := application.NewGrpcServer(s.Application)
	srv.Services.ApplicationService = appServer
	srv.Servers = append(srv.Servers, appServer)

	enterpriseServer := enterprise.NewGrpcServer(s.Enterprise)
	srv.Services.EnterpriseService = enterpriseServer
	srv.Servers = append(srv.Servers, enterpriseServer)

	return &srv
}
