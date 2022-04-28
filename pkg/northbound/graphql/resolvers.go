package graphql

import (
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/application"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/enterprise"
)

type enterpriseRoot struct {
	grpcServer v1.EnterpriseServiceServer
}

//func (e enterpriseRoot) Mutation() enterprise.MutationResolver {
//	return &v1.ApplicationApiResolvers{
//		Service: v1.ApplicationApiServer(nil),
//	}
//}

func (e *enterpriseRoot) Query() enterprise.QueryResolver {
	return &v1.EnterpriseServiceResolvers{
		Service: e.grpcServer,
	}
}

func NewEnterpriseResolver(grpcServer v1.EnterpriseServiceServer) *enterpriseRoot {
	return &enterpriseRoot{
		grpcServer: grpcServer,
	}
}

type applicationRoot struct {
	grpcServer v1.ApplicationServiceServer
}

func (a *applicationRoot) Mutation() application.MutationResolver {
	return &v1.ApplicationServiceResolvers{
		Service: a.grpcServer,
	}
}

func (a *applicationRoot) Query() application.QueryResolver {
	return &v1.ApplicationServiceResolvers{
		Service: a.grpcServer,
	}
}

func NewApplicationResolver(grpcServer v1.ApplicationServiceServer) *applicationRoot {
	return &applicationRoot{
		grpcServer: grpcServer,
	}
}
