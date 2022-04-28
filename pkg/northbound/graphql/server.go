// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"net/http"
	"runtime/debug"
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

	srv := handler.NewDefaultServer(v1.(todo.New()))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send this panic somewhere
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8081", nil))

	go func() {
		log.Infof("GraphQL API server listening on %s", s.address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("Could not start GraphQL server: %v", err)
			return
		}
	}()

	x := <-s.doneCh
	if x {
		log.Warnf("Stopping API REST server")
		_ = server.Shutdown(ctx)
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
