// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/pkg/northbound/grpc/application"
	"github.com/onosproject/scaling-umbrella/pkg/northbound/grpc/enterprise"
	"github.com/onosproject/scaling-umbrella/pkg/southbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
)

var log = logging.GetLogger("GrpcServer")

type IRocApiGrpcServer interface {
	StartGrpcServer(srv grpc.ServiceRegistrar)
}

type RocApiGrpcServer struct {
	southboundManager  *southbound.GnmiManager
	doneCh             chan bool
	wg                 *sync.WaitGroup
	address            string
	Servers            []IRocApiGrpcServer
	ApplicationService v1.ApplicationServiceServer
	EnterpriseService  v1.EnterpriseServiceServer
}

func (s *RocApiGrpcServer) StartGrpcServer() error {
	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", s.address)
	log.Infow("starting-grpc-server", "address", s.address)
	if err != nil {
		return err
	}

	for _, s := range s.Servers {
		s.StartGrpcServer(grpcServer)
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

func NewGrpcServer(doneCh chan bool, wg *sync.WaitGroup, address string, sbManager *southbound.GnmiManager) *RocApiGrpcServer {

	srv := RocApiGrpcServer{
		southboundManager: sbManager,
		doneCh:            doneCh,
		wg:                wg,
		address:           address,
		Servers:           []IRocApiGrpcServer{},
	}

	appServer := application.NewGrpcServer(doneCh, sbManager)
	srv.ApplicationService = appServer
	srv.Servers = append(srv.Servers, appServer)

	enterpriseServer := enterprise.NewGrpcServer(doneCh, sbManager)
	srv.EnterpriseService = enterpriseServer
	srv.Servers = append(srv.Servers, enterpriseServer)

	return &srv
}
