/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package grpc

import (
	"github.com/onosproject/aether-operations-api/gen/go/v1"
	"github.com/onosproject/aether-operations-api/internal/stores"
	"github.com/onosproject/aether-operations-api/internal/stores/application"
	"github.com/onosproject/aether-operations-api/internal/stores/device"
	"github.com/onosproject/aether-operations-api/internal/stores/devicegroup"
	"github.com/onosproject/aether-operations-api/internal/stores/enterprise"
	"github.com/onosproject/aether-operations-api/internal/stores/simcard"
	"github.com/onosproject/aether-operations-api/internal/stores/site"
	"github.com/onosproject/aether-operations-api/internal/stores/slice"
	"github.com/onosproject/aether-operations-api/internal/stores/smallcell"
	"github.com/onosproject/onos-lib-go/pkg/logging"
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
	EnterpriseService  v1.EnterpriseServiceServer
	ApplicationService v1.ApplicationServiceServer
	SiteService        v1.SiteServiceServer
	DeviceService      v1.DeviceServiceServer
	SimCardService     v1.SimCardServiceServer
	DeviceGroupService v1.DeviceGroupServiceServer
	SliceService       v1.SliceServiceServer
	SmallCellService   v1.SmallCellServiceServer
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

	enterpriseServer := enterprise.NewGrpcServer(s.Enterprise)
	srv.Services.EnterpriseService = enterpriseServer
	srv.Servers = append(srv.Servers, enterpriseServer)

	appServer := application.NewGrpcServer(s.Application)
	srv.Services.ApplicationService = appServer
	srv.Servers = append(srv.Servers, appServer)

	siteServer := site.NewGrpcServer(s.Site)
	srv.Services.SiteService = siteServer
	srv.Servers = append(srv.Servers, siteServer)

	deviceServer := device.NewGrpcServer(s.Device)
	srv.Services.DeviceService = deviceServer
	srv.Servers = append(srv.Servers, deviceServer)

	simCardServer := simcard.NewGrpcServer(s.SimCard)
	srv.Services.SimCardService = simCardServer
	srv.Servers = append(srv.Servers, simCardServer)

	deviceGroupServer := devicegroup.NewGrpcServer(s.DeviceGroup)
	srv.Services.DeviceGroupService = deviceGroupServer
	srv.Servers = append(srv.Servers, deviceGroupServer)

	sliceServer := slice.NewGrpcServer(s.Slice)
	srv.Services.SliceService = sliceServer
	srv.Servers = append(srv.Servers, sliceServer)

	smallCellServer := smallcell.NewGrpcServer(s.SmallCell)
	srv.Services.SmallCellService = smallCellServer
	srv.Servers = append(srv.Servers, smallCellServer)

	return &srv
}
