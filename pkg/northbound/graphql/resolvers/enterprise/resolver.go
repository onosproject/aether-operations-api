package enterprise

import (
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/enterprise"
)

type EnterpriseResolver struct {
	grpcServer v1.EnterpriseServiceServer
}

//func (e EnterpriseResolver) Mutation() enterprise.MutationResolver {
//	return &v1.ApplicationApiResolvers{
//		Service: v1.ApplicationApiServer(nil),
//	}
//}

func (e *EnterpriseResolver) Query() enterprise.QueryResolver {
	return &v1.EnterpriseServiceResolvers{
		Service: e.grpcServer,
	}
}

func NewEnterpriseResolver(grpcServer v1.EnterpriseServiceServer) *EnterpriseResolver {
	return &EnterpriseResolver{
		grpcServer: grpcServer,
	}
}
