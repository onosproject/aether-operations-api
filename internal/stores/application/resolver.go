package application

import (
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/application"
)

type ApplicationResolver struct {
	grpcServer v1.ApplicationServiceServer
}

func (a *ApplicationResolver) Mutation() application.MutationResolver {
	return &v1.ApplicationServiceResolvers{
		Service: a.grpcServer,
	}
}

func (a *ApplicationResolver) Query() application.QueryResolver {
	return &v1.ApplicationServiceResolvers{
		Service: a.grpcServer,
	}
}

func NewApplicationResolver(grpcServer v1.ApplicationServiceServer) *ApplicationResolver {
	return &ApplicationResolver{
		grpcServer: grpcServer,
	}
}
