package application

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/roc-api/api/v1"
	"github.com/onosproject/roc-api/pkg/southbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var log = logging.GetLogger("ApplicationGrpcServer")

type ApplicationApiGrpcServer struct {
	southboundManager *southbound.GnmiManager
	doneCh            chan bool
	v1.UnimplementedApplicationApiServer
}

func (r *ApplicationApiGrpcServer) StartGrpcServer(srv grpc.ServiceRegistrar) {
	v1.RegisterApplicationApiServer(srv, r)
}

func (r *ApplicationApiGrpcServer) GetApplications(ctx context.Context, entId *v1.EnterpriseId) (*v1.Applications, error) {
	return r.southboundManager.ApplicationHandler.ListApplications(entId.EnterpriseId)
}

func (r *ApplicationApiGrpcServer) CreateApplication(context.Context, *v1.Application) (*v1.Application, error) {
	return nil, status.Error(codes.Unimplemented, "create-applications-not-implemented")
}

//func (r *RocApiGrpcServer) GetEnterprises(ctx context.Context, empty *v1.Empty) (*v1.Enterprise, error) {
//	return nil, status.Error(codes.Unimplemented, "get-enterprises-not-implemented")
//}

func NewGrpcServer(doneCh chan bool, sbManager *southbound.GnmiManager) *ApplicationApiGrpcServer {
	srv := ApplicationApiGrpcServer{
		southboundManager: sbManager,
		doneCh:            doneCh,
	}

	return &srv
}
