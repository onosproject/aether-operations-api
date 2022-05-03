package simcard

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/simcards/v1"
	"google.golang.org/grpc"
)

type SimCardServiceGrpcServer struct {
	handler *SimCardHandler
	v1.UnimplementedSimCardServiceServer
}

func (r *SimCardServiceGrpcServer) RegisterGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterSimCardServiceServer(srv, r)
}

func (r *SimCardServiceGrpcServer) GetSimCards(ctx context.Context, req *v1.GetSimCardsRequest) (*v1.GetSimCardsResponse, error) {
	return r.handler.ListSimCards(req.EnterpriseId, req.SiteId)
}

func NewGrpcServer(handler *SimCardHandler) *SimCardServiceGrpcServer {
	srv := SimCardServiceGrpcServer{
		handler: handler,
	}

	return &srv
}

func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, grpcConn *grpc.ClientConn) error {
	if err := v1.RegisterSimCardServiceHandler(ctx, mux, grpcConn); err != nil {
		log.Errorf("Could not register SimCardService handler: %v", err)
		return err
	}
	return nil
}
