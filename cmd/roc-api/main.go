package main

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/roc-api/pkg/northbound/grpc"
	"github.com/onosproject/roc-api/pkg/northbound/rest"
	"github.com/onosproject/roc-api/pkg/southbound"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const onosConfigAddress = "localhost:5150"
const grpcEndpoint = "0.0.0.0:50060"
const restEndpoint = "0.0.0.0:8181"

var log = logging.GetLogger("RocApi")

func main() {
	log.Info("roc-api started")

	// create the southbound manager
	gnmiManager, err := southbound.NewGnmiManager(onosConfigAddress)
	if err != nil {
		log.Fatalw("cannot-start-gnmi-client", "err", err)
	}

	// start the gRPC server
	doneChannel := make(chan bool)
	wg := sync.WaitGroup{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		<-sigs
		close(doneChannel)
	}()

	wg.Add(1)
	grpcSrv, err := grpc.NewGrpcServer(doneChannel, &wg, grpcEndpoint, gnmiManager)
	if err != nil {
		log.Fatal("cannot start grpc server")
	}
	go grpcSrv.StartGrpcServer()

	wg.Add(1)
	restSrv, err := rest.NewRestServer(doneChannel, &wg, restEndpoint, grpcEndpoint)
	if err != nil {
		log.Fatal("cannot start rest server")
	}
	go restSrv.StartRestServer()

	wg.Wait()
}
