/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package onos_topo

import (
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"time"
)

var log = logging.GetLogger("OnosTopoDatasource")

func NewOnosTopoClient(address string) (topo.TopoClient, error) {
	log.Infow("initializing-onos-topo-client", "address", address)
	opts, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
		return nil, err
	}

	optsWithRetry := []grpc.DialOption{
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithInterval(100 * time.Millisecond))),
	}
	optsWithRetry = append(opts, optsWithRetry...)
	conn, err := grpc.Dial(address, optsWithRetry...)

	if err != nil {
		return nil, err
	}

	topoClient := topo.NewTopoClient(conn)
	return topoClient, nil
}
