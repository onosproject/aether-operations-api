/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package site

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/onosproject/aether-operations-api/gen/go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SiteServiceGrpcServer struct {
	handler *SiteHandler
	v1.UnimplementedSiteServiceServer
}

func (r *SiteServiceGrpcServer) RegisterGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterSiteServiceServer(srv, r)
}

func (r *SiteServiceGrpcServer) GetSites(ctx context.Context, entId *v1.GetSitesRequest) (*v1.GetSitesResponse, error) {
	return r.handler.ListSites(entId.EnterpriseId)
}

func (r *SiteServiceGrpcServer) CreateSite(context.Context, *v1.Site) (*v1.Site, error) {
	return nil, status.Error(codes.Unimplemented, "create-site-not-implemented")
}

func NewGrpcServer(handler *SiteHandler) *SiteServiceGrpcServer {
	srv := SiteServiceGrpcServer{
		handler: handler,
	}

	return &srv
}

func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, grpcConn *grpc.ClientConn) error {
	if err := v1.RegisterSiteServiceHandler(ctx, mux, grpcConn); err != nil {
		log.Errorf("Could not register SiteService handler: %v", err)
		return err
	}
	return nil
}
