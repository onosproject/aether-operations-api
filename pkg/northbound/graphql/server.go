// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/application"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/enterprise"
	ar "github.com/onosproject/scaling-umbrella/pkg/northbound/graphql/resolvers/application"
	er "github.com/onosproject/scaling-umbrella/pkg/northbound/graphql/resolvers/enterprise"
	"github.com/onosproject/scaling-umbrella/pkg/northbound/grpc"
	"net/http"
	"sync"
)

var log = logging.GetLogger("GqlServer")

type RocApiGqlServer struct {
	doneCh     chan bool
	wg         *sync.WaitGroup
	address    string
	grpcServer *grpc.RocApiGrpcServer
}

func (s RocApiGqlServer) StartGqlServer() {

	appResolver := ar.NewApplicationResolver(s.grpcServer.ApplicationService)
	appSrv := handler.NewDefaultServer(application.NewExecutableSchema(application.Config{
		Resolvers: appResolver,
	}))

	entResolver := er.NewEnterpriseResolver(s.grpcServer.EnterpriseService)
	entSrv := handler.NewDefaultServer(enterprise.NewExecutableSchema(enterprise.Config{
		Resolvers: entResolver,
	}))

	http.Handle("/enterprise-playground", playground.Handler("ROC API", "/enterprise-query"))
	http.Handle("/enterprise-query", entSrv)
	http.Handle("/application-playground", playground.Handler("ROC API", "/application-query"))
	http.Handle("/application-query", appSrv)

	go func() {
		log.Infof("GraphQL API server listening on %s", s.address)
		log.Fatal(http.ListenAndServe(":8081", nil))
		//if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		//	log.Errorf("Could not start GraphQL server: %v", err)
		//	return
		//}
	}()

	x := <-s.doneCh
	if x {
		log.Warnf("Stopping API REST server")
	}

	s.wg.Done()
}

func NewGqlServer(doneCh chan bool, wg *sync.WaitGroup, address string, grpcServer *grpc.RocApiGrpcServer) (*RocApiGqlServer, error) {
	srv := RocApiGqlServer{
		doneCh:     doneCh,
		wg:         wg,
		address:    address,
		grpcServer: grpcServer,
	}

	return &srv, nil
}
