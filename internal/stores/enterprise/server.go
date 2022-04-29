/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package enterprise

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"google.golang.org/grpc"
)

type EnterpriseServiceGrpcServer struct {
	handler *EnterpriseHandler
	v1.UnimplementedEnterpriseServiceServer
}

func (r *EnterpriseServiceGrpcServer) RegisterGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterEnterpriseServiceServer(srv, r)
}

func (r *EnterpriseServiceGrpcServer) GetEnterprises(ctx context.Context, empty *v1.Empty) (*v1.Enterprises, error) {
	return r.handler.ListEnterprises()
}

func NewGrpcServer(handler *EnterpriseHandler) *EnterpriseServiceGrpcServer {
	srv := EnterpriseServiceGrpcServer{
		handler: handler,
	}

	return &srv
}

func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, grpcConn *grpc.ClientConn) error {
	if err := v1.RegisterEnterpriseServiceHandler(ctx, mux, grpcConn); err != nil {
		log.Errorf("Could not register EnterpriseService handler: %v", err)
		return err
	}
	return nil
}
