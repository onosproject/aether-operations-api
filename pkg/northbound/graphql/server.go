// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/application"
	"net/http"
	"sync"
)

var log = logging.GetLogger("GqlServer")

type RocApiGqlServer struct {
	doneCh      chan bool
	wg          *sync.WaitGroup
	address     string
	grpcAddress string
}

func (s RocApiGqlServer) StartGqlServer() {

	appSrv := handler.NewDefaultServer(application.NewExecutableSchema(application.Config{
		Resolvers: applicationRoot{},
	}))

	http.Handle("/", playground.Handler("ROC API", "/query"))
	http.Handle("/query", appSrv)

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

func NewGqlServer(doneCh chan bool, wg *sync.WaitGroup, address string, grpcAddress string) (*RocApiGqlServer, error) {
	srv := RocApiGqlServer{
		doneCh:      doneCh,
		wg:          wg,
		address:     address,
		grpcAddress: grpcAddress,
	}

	return &srv, nil
}
