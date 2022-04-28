// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package application

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/pkg/southbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var log = logging.GetLogger("ApplicationGrpcServer")

type ApplicationServiceGrpcServer struct {
	southboundManager *southbound.GnmiManager
	doneCh            chan bool
	v1.UnimplementedApplicationServiceServer
}

func (r *ApplicationServiceGrpcServer) StartGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterApplicationServiceServer(srv, r)
}

func (r *ApplicationServiceGrpcServer) GetApplications(ctx context.Context, entId *v1.EnterpriseId) (*v1.Applications, error) {
	return r.southboundManager.ApplicationHandler.ListApplications(entId.EnterpriseId)
}

func (r *ApplicationServiceGrpcServer) CreateApplication(context.Context, *v1.Application) (*v1.Application, error) {
	return nil, status.Error(codes.Unimplemented, "create-applications-not-implemented")
}

func NewGrpcServer(doneCh chan bool, sbManager *southbound.GnmiManager) *ApplicationServiceGrpcServer {
	srv := ApplicationServiceGrpcServer{
		southboundManager: sbManager,
		doneCh:            doneCh,
	}

	return &srv
}
