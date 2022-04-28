/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package application

import (
	"context"
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApplicationServiceGrpcServer struct {
	handler *ApplicationHandler
	v1.UnimplementedApplicationServiceServer
}

func (r *ApplicationServiceGrpcServer) StartGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterApplicationServiceServer(srv, r)
}

func (r *ApplicationServiceGrpcServer) GetApplications(ctx context.Context, entId *v1.EnterpriseId) (*v1.Applications, error) {
	return r.handler.ListApplications(entId.EnterpriseId)
}

func (r *ApplicationServiceGrpcServer) CreateApplication(context.Context, *v1.Application) (*v1.Application, error) {
	return nil, status.Error(codes.Unimplemented, "create-applications-not-implemented")
}

func NewGrpcServer(handler *ApplicationHandler) *ApplicationServiceGrpcServer {
	srv := ApplicationServiceGrpcServer{
		handler: handler,
	}

	return &srv
}
