package graphql

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"sync"
)

var log = logging.GetLogger("GqlServer")

type RocApiGqlServer struct {
	doneCh      chan bool
	wg          *sync.WaitGroup
	address     string
	grpcAddress string
}

//func (s RocApiGqlServer) StartGqlServer() {
//
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	mux := runtime.NewServeMux()
//
//	opts := []grpc.DialOption{grpc.WithInsecure()}
//	conn, err := grpc.Dial(s.grpcAddress, opts...)
//	if err != nil {
//		log.Errorf("Could not start gRPC client: %v", err)
//		return
//	}
//	if err := v1.RegisterApplicationApiGraphqlHandler(mux, conn); err != nil {
//		log.Fatal(err)
//	}
//	// FIXME seems like the GraphQL Gateway plugin from ysugimoto only generates one service
//	//if err := v1.RegisterEnterpriseApiGraphqlHandler(mux, conn); err != nil {
//	//	log.Fatal(err)
//	//}
//	http.Handle("/graphql", mux)
//
//	server := &http.Server{Addr: s.address, Handler: mux}
//
//	go func() {
//		log.Infof("GraphQL API server listening on %s", s.address)
//		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			log.Errorf("Could not start GraphQL server: %v", err)
//			return
//		}
//	}()
//
//	x := <-s.doneCh
//	if x {
//		log.Warnf("Stopping API REST server")
//		_ = server.Shutdown(ctx)
//	}
//
//	s.wg.Done()
//}
//
//func NewGqlServer(doneCh chan bool, wg *sync.WaitGroup, address string, grpcAddress string) (*RocApiGqlServer, error) {
//	srv := RocApiGqlServer{
//		doneCh:      doneCh,
//		wg:          wg,
//		address:     address,
//		grpcAddress: grpcAddress,
//	}
//
//	return &srv, nil
//}
