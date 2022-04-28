package graphql

import (
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	graphQl "github.com/onosproject/scaling-umbrella/api/v1/gqlgen"
)

type enterpriseRoot struct{}

func (e enterpriseRoot) Mutation() graphQl.MutationResolver {
	return &v1.ApplicationApiResolvers{
		Service: v1.ApplicationApiServer(nil),
	}
}

func (e enterpriseRoot) Query() graphQl.QueryResolver {
	return &v1.ApplicationApiResolvers{
		Service: v1.ApplicationApiServer(nil),
	}
}
