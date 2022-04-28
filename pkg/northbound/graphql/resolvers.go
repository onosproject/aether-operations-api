package graphql

import (
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/application"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/enterprise"
)

type enterpriseRoot struct{}

//func (e enterpriseRoot) Mutation() enterprise.MutationResolver {
//	return &v1.ApplicationApiResolvers{
//		Service: v1.ApplicationApiServer(nil),
//	}
//}

func (e enterpriseRoot) Query() enterprise.QueryResolver {
	return &v1.EnterpriseServiceResolvers{
		Service: v1.EnterpriseServiceServer(nil),
	}
}

type applicationRoot struct {
}

func (a applicationRoot) Mutation() application.MutationResolver {
	return &v1.ApplicationServiceResolvers{
		Service: v1.ApplicationServiceServer(nil),
	}
}

func (a applicationRoot) Query() application.QueryResolver {
	return &v1.ApplicationServiceResolvers{
		Service: v1.ApplicationServiceServer(nil),
	}
}
