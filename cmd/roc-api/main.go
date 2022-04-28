// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/internal/config"
	"github.com/onosproject/scaling-umbrella/internal/datasources"
	"github.com/onosproject/scaling-umbrella/internal/servers/grpc"
	"github.com/onosproject/scaling-umbrella/internal/servers/rest"
	"github.com/onosproject/scaling-umbrella/internal/stores"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const onosConfigAddress = "localhost:5150"
const grpcEndpoint = "0.0.0.0:50060"
const restEndpoint = "0.0.0.0:8080"
const gqlEndpoint = "0.0.0.0:8081" // FIXME use the same mux server for REST and GraphQL

var log = logging.GetLogger("RocApi")

func main() {

	cfg := config.GetConfig()

	log.Infow("roc-api started", "cfg", cfg)

	// setup the datasources
	ds, err := datasources.RegisterDatasources(cfg)
	if err != nil {
		log.Fatalw("cannot-setup-datasources", "err", err)
	}

	// setup the stores
	s, err := stores.RegisterStores(ds)
	if err != nil {
		log.Fatalw("cannot-setup-stores", "err", err)
	}

	// create a channel to manage the servers lifecycle
	doneChannel := make(chan bool)
	// create a Waitgroup to wait for the server to exit before shutting down
	wg := sync.WaitGroup{}

	// listen on SIGTERM to signal to the servers to shutdown (via doneChannel)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		<-sigs
		close(doneChannel)
	}()

	wg.Add(1)
	grpcSrv := grpc.NewGrpcServer(doneChannel, &wg, cfg.ServersConfig.GrpcAddress, s)
	go grpcSrv.StartGrpcServer()

	wg.Add(1)
	restSrv, err := rest.NewRestServer(doneChannel, &wg, cfg.ServersConfig.RestAddress, cfg.ServersConfig.GrpcAddress)
	if err != nil {
		log.Fatal("cannot start rest server")
	}
	go restSrv.StartRestServer()
	//
	//wg.Add(1)
	//gqlSrv, err := graphql.NewGqlServer(doneChannel, &wg, gqlEndpoint, grpcSrv)
	//if err != nil {
	//	log.Fatal("cannot start graphql server")
	//}
	//go gqlSrv.StartGqlServer()
	//
	wg.Wait()
}
