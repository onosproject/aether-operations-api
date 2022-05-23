/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package slice

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/onosproject/scaling-umbrella/gen/go/v1"
	"google.golang.org/grpc"
)

type SliceServiceGrpcServer struct {
	handler *SliceHandler
	v1.UnimplementedSliceServiceServer
}

func (r *SliceServiceGrpcServer) RegisterGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterSliceServiceServer(srv, r)
}

func (r *SliceServiceGrpcServer) GetSlices(ctx context.Context, req *v1.GetSlicesRequest) (*v1.GetSlicesResponse, error) {
	return r.handler.ListSlices(req.EnterpriseId, req.SiteId)
}

func NewGrpcServer(handler *SliceHandler) *SliceServiceGrpcServer {
	srv := SliceServiceGrpcServer{
		handler: handler,
	}

	return &srv
}

func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, grpcConn *grpc.ClientConn) error {
	if err := v1.RegisterSliceServiceHandler(ctx, mux, grpcConn); err != nil {
		log.Errorf("Could not register SliceService handler: %v", err)
		return err
	}
	return nil
}
