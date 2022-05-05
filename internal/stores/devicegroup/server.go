/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package devicegroup

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/devicegroups/v1"
	"google.golang.org/grpc"
)

type DeviceGroupServiceGrpcServer struct {
	handler *DeviceGroupHandler
	v1.UnimplementedDeviceGroupServiceServer
}

func (r *DeviceGroupServiceGrpcServer) RegisterGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterDeviceGroupServiceServer(srv, r)
}

func (r *DeviceGroupServiceGrpcServer) GetDeviceGroups(ctx context.Context, req *v1.GetDeviceGroupsRequest) (*v1.GetDeviceGroupsResponse, error) {
	return r.handler.ListDeviceGroups(req.EnterpriseId, req.SiteId)
}

func NewGrpcServer(handler *DeviceGroupHandler) *DeviceGroupServiceGrpcServer {
	srv := DeviceGroupServiceGrpcServer{
		handler: handler,
	}

	return &srv
}

func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, grpcConn *grpc.ClientConn) error {
	if err := v1.RegisterDeviceGroupServiceHandler(ctx, mux, grpcConn); err != nil {
		log.Errorf("Could not register DeviceGroupService handler: %v", err)
		return err
	}
	return nil
}
