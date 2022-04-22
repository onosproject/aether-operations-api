package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/labstack/gommon/log"
	gw "github.com/onosproject/roc-api/api/v1/gen/api/v1"
	"google.golang.org/grpc"
	"net/http"
	"sync"
)

type RocApiRestServer struct {
	doneCh      chan bool
	wg          *sync.WaitGroup
	address     string
	grpcAddress string
}

func (s RocApiRestServer) StartRestServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := gw.RegisterRocApiHandlerFromEndpoint(ctx, mux, s.grpcAddress, opts); err != nil {
		log.Errorf("Could not register API server: %v", err)
		return
	}

	server := &http.Server{Addr: s.address, Handler: mux}

	go func() {
		log.Infof("REST API server listening on %s", s.address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("Could not start API server: %v", err)
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

func NewRestServer(doneCh chan bool, wg *sync.WaitGroup, address string, grpcAddress string) (*RocApiRestServer, error) {
	srv := RocApiRestServer{
		doneCh:      doneCh,
		wg:          wg,
		address:     address,
		grpcAddress: grpcAddress,
	}

	return &srv, nil
}
