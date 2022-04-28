// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package enterprise

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/pkg/southbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("EnterpriseGrpcServer")

type EnterpriseServiceGrpcServer struct {
	southboundManager *southbound.GnmiManager
	doneCh            chan bool
	v1.UnimplementedEnterpriseServiceServer
}

func (r *EnterpriseServiceGrpcServer) StartGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterEnterpriseServiceServer(srv, r)
}

func (r *EnterpriseServiceGrpcServer) GetEnterprises(ctx context.Context, empty *v1.Empty) (*v1.Enterprises, error) {
	return r.southboundManager.EnterpriseHandler.ListEnterprises()
}

func NewGrpcServer(doneCh chan bool, sbManager *southbound.GnmiManager) *EnterpriseServiceGrpcServer {
	srv := EnterpriseServiceGrpcServer{
		southboundManager: sbManager,
		doneCh:            doneCh,
	}

	return &srv
}
