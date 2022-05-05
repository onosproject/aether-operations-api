/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package device

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/onosproject/scaling-umbrella/gen/go/devices/v1"
	"google.golang.org/grpc"
)

type DeviceServiceGrpcServer struct {
	handler *DeviceHandler
	v1.UnimplementedDeviceServiceServer
}

func (r *DeviceServiceGrpcServer) RegisterGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterDeviceServiceServer(srv, r)
}

func (r *DeviceServiceGrpcServer) GetDevices(ctx context.Context, req *v1.GetDevicesRequest) (*v1.GetDevicesResponse, error) {
	return r.handler.ListDevices(req.EnterpriseId, req.SiteId)
}

func NewGrpcServer(handler *DeviceHandler) *DeviceServiceGrpcServer {
	srv := DeviceServiceGrpcServer{
		handler: handler,
	}

	return &srv
}

func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, grpcConn *grpc.ClientConn) error {
	if err := v1.RegisterDeviceServiceHandler(ctx, mux, grpcConn); err != nil {
		log.Errorf("Could not register DeviceService handler: %v", err)
		return err
	}
	return nil
}
