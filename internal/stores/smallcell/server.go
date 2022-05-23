/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package smallcell

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/onosproject/scaling-umbrella/gen/go/v1"
	"google.golang.org/grpc"
)

type SmallCellServiceGrpcServer struct {
	handler *SmallCellHandler
	v1.UnimplementedSmallCellServiceServer
}

func (r *SmallCellServiceGrpcServer) RegisterGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterSmallCellServiceServer(srv, r)
}

func (r *SmallCellServiceGrpcServer) GetSmallCells(ctx context.Context, req *v1.GetSmallCellsRequest) (*v1.GetSmallCellsResponse, error) {
	return r.handler.ListSmallCells(req.EnterpriseId, req.SiteId)
}

func NewGrpcServer(handler *SmallCellHandler) *SmallCellServiceGrpcServer {
	srv := SmallCellServiceGrpcServer{
		handler: handler,
	}

	return &srv
}

func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, grpcConn *grpc.ClientConn) error {
	if err := v1.RegisterSmallCellServiceHandler(ctx, mux, grpcConn); err != nil {
		log.Errorf("Could not register SmallCellService handler: %v", err)
		return err
	}
	return nil
}
