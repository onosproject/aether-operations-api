/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package application

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/onosproject/scaling-umbrella/gen/go/applications/v1"
	"google.golang.org/grpc"
)

type ApplicationServiceGrpcServer struct {
	handler *ApplicationHandler
	v1.UnimplementedApplicationServiceServer
}

func (r *ApplicationServiceGrpcServer) RegisterGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterApplicationServiceServer(srv, r)
}

func (r *ApplicationServiceGrpcServer) GetApplications(ctx context.Context, req *v1.GetApplicationsRequest) (*v1.Applications, error) {
	return r.handler.ListApplications(ctx, req.EnterpriseId)
}

func (r *ApplicationServiceGrpcServer) GetApplication(ctx context.Context, req *v1.ApplicationFilter) (*v1.Application, error) {
	return r.handler.GetApplication(ctx, req)
}

func (r *ApplicationServiceGrpcServer) CreateOrUpdateApplication(ctx context.Context, app *v1.Application) (*v1.Application, error) {
	return r.handler.CreateApplication(ctx, app)
}

func (r *ApplicationServiceGrpcServer) DeleteApplication(ctx context.Context, req *v1.ApplicationFilter) (*v1.Empty, error) {
	return r.handler.DeleteApplication(ctx, req)
}

func NewGrpcServer(handler *ApplicationHandler) *ApplicationServiceGrpcServer {
	srv := ApplicationServiceGrpcServer{
		handler: handler,
	}

	return &srv
}

func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, grpcConn *grpc.ClientConn) error {
	if err := v1.RegisterApplicationServiceHandler(ctx, mux, grpcConn); err != nil {
		log.Errorf("Could not register ApplicationService handler: %v", err)
		return err
	}
	return nil
}
