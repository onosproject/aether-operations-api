// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/internal/config"
	"github.com/onosproject/scaling-umbrella/internal/datasources"
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
	_, err := datasources.RegisterDatasources(cfg)
	if err != nil {
		log.Fatalw("cannot-setup-datasources", "err", err)
	}

	// create the southbound manager
	//gnmiManager, err := southbound.NewGnmiManager(onosConfigAddress)
	//if err != nil {
	//	log.Fatalw("cannot-start-gnmi-client", "err", err)
	//}
	//
	//// start the gRPC server
	//doneChannel := make(chan bool)
	//wg := sync.WaitGroup{}
	//
	//sigs := make(chan os.Signal, 1)
	//signal.Notify(sigs, syscall.SIGTERM)
	//go func() {
	//	<-sigs
	//	close(doneChannel)
	//}()
	//
	//wg.Add(1)
	//grpcSrv := grpc.NewGrpcServer(doneChannel, &wg, grpcEndpoint, gnmiManager)
	//go grpcSrv.StartGrpcServer()
	//
	//wg.Add(1)
	//restSrv, err := rest.NewRestServer(doneChannel, &wg, restEndpoint, grpcEndpoint)
	//if err != nil {
	//	log.Fatal("cannot start rest server")
	//}
	//go restSrv.StartRestServer()
	//
	//wg.Add(1)
	//gqlSrv, err := graphql.NewGqlServer(doneChannel, &wg, gqlEndpoint, grpcSrv)
	//if err != nil {
	//	log.Fatal("cannot start graphql server")
	//}
	//go gqlSrv.StartGqlServer()
	//
	//wg.Wait()
}
