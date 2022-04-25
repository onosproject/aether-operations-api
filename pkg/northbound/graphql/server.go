package graphql

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/roc-api/api/v1"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(s.grpcAddress, opts...)
	if err != nil {
		log.Errorf("Could not start gRPC client: %v", err)
		return
	}
	if err := v1.RegisterRocApiGraphqlHandler(mux, conn); err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", mux)

	server := &http.Server{Addr: s.address, Handler: mux}

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
