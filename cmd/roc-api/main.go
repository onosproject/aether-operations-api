// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/internal/config"
	"github.com/onosproject/scaling-umbrella/internal/datasources"
	"github.com/onosproject/scaling-umbrella/internal/servers/grpc"
	"github.com/onosproject/scaling-umbrella/internal/servers/http"
	"github.com/onosproject/scaling-umbrella/internal/stores"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

var log = logging.GetLogger("RocApi")

func stringToLogLevel(l string) logging.Level {
	switch strings.ToLower(l) {
	case "debug":
		return logging.DebugLevel
	case "info":
		return logging.InfoLevel
	case "warn":
		return logging.WarnLevel
	case "error":
		return logging.ErrorLevel
	default:
		log.Warnf("Provided log level (%s) is unknown, defaulting to INFO. Valid values are: DEBUG, INFO, WARN, ERROR")
		return logging.InfoLevel
	}
}

func main() {

	cfg := config.GetConfig()

	l := stringToLogLevel(cfg.LogLevel)
	log.SetLevel(l)

	if l != logging.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}

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
	go func() {
		if err := grpcSrv.StartGrpcServer(); err != nil {
			log.Fatalw("cannot-start-grpc-server", "err", err)
		}
	}()

	wg.Add(1)
	httpSrv, err := http.NewHttpServer(doneChannel, &wg, *cfg, grpcSrv)
	if err != nil {
		log.Fatalw("cannot-create-rest-server", "err", err)
	}
	go func() {
		if err := httpSrv.StartHttpServer(); err != nil {
			log.Fatalw("cannot-start-http-server", "err", err)
		}
	}()

	wg.Wait()
}
