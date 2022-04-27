// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package enterprise

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/roc-api/api/v1"
	"github.com/onosproject/roc-api/pkg/southbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var log = logging.GetLogger("EnterpriseGrpcServer")

type EnterpriseApiGrpcServer struct {
	southboundManager *southbound.GnmiManager
	doneCh            chan bool
	v1.UnimplementedEnterpriseApiServer
}

func (r *EnterpriseApiGrpcServer) StartGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterEnterpriseApiServer(srv, r)
}

func (r *EnterpriseApiGrpcServer) GetEnterprises(ctx context.Context, empty *v1.Empty) (*v1.Enterprise, error) {
	return nil, status.Error(codes.Unimplemented, "get-enterprises-not-implemented")
}

func NewGrpcServer(doneCh chan bool, sbManager *southbound.GnmiManager) *EnterpriseApiGrpcServer {
	srv := EnterpriseApiGrpcServer{
		southboundManager: sbManager,
		doneCh:            doneCh,
	}

	return &srv
}
