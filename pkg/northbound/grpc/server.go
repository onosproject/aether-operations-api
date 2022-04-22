package grpc

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/roc-api/api/v1"
	"github.com/onosproject/roc-api/pkg/southbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
)

var log = logging.GetLogger("GrpcServer")

type RocApiGrpcServer struct {
	southboundManager *southbound.GnmiManager
	doneCh            chan bool
	wg                *sync.WaitGroup
	address           string
}

func (r *RocApiGrpcServer) StartGrpcServer() error {
	lis, err := net.Listen("tcp", r.address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	v1.RegisterRocApiServer(grpcServer, r)

	reflection.Register(grpcServer)

	go func() {
		log.Infow("starting-grpc-server", "address", r.address)
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Errorw("grpc-server-error", "err", err)
		}
	}()

	x := <-r.doneCh
	if x {
		// if the API channel is closed, stop the gRPC server
		grpcServer.Stop()
		log.Info("Stopping gRPC server")
	}

	r.wg.Done()
	return nil
}

func (r *RocApiGrpcServer) GetApplications(ctx context.Context, empty *v1.Empty) (*v1.Applications, error) {
	return r.southboundManager.ApplicationHandler.ListApplications()
}

func (r *RocApiGrpcServer) GetEnterprises(ctx context.Context, empty *v1.Empty) (*v1.Enterprise, error) {
	//TODO implement me
	panic("implement me")
}

func NewGrpcServer(doneCh chan bool, wg *sync.WaitGroup, address string, sbManager *southbound.GnmiManager) (*RocApiGrpcServer, error) {
	srv := RocApiGrpcServer{
		southboundManager: sbManager,
		doneCh:            doneCh,
		wg:                wg,
		address:           address,
	}

	return &srv, nil
}
