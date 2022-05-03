/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package enterprise

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/onosproject/scaling-umbrella/api/v1/gqlgen/enterprise"
	//v1 "github.com/onosproject/scaling-umbrella/api/v1"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/enterprises/v1"
)

type EnterpriseResolver struct {
	grpcServer v1.EnterpriseServiceServer
}

//func (e EnterpriseResolver) Mutation() enterprise.MutationResolver {
//	return &v1.ApplicationApiResolvers{
//		Service: v1.ApplicationApiServer(nil),
//	}
//}

//func (e *EnterpriseResolver) Query() enterprise.QueryResolver {
//	return &v1.EnterpriseServiceResolvers{
//		Service: e.grpcServer,
//	}
//}

func NewEnterpriseResolver(grpcServer v1.EnterpriseServiceServer) *EnterpriseResolver {
	return &EnterpriseResolver{
		grpcServer: grpcServer,
	}
}

func RegisterGraphQlHandler(server v1.EnterpriseServiceServer, router *gin.Engine) {
	r := NewEnterpriseResolver(server)
	s := handler.NewDefaultServer(enterprise.NewExecutableSchema(enterprise.Config{
		Resolvers: r,
	}))

	router.POST("/enterprise-query", func(c *gin.Context) {
		s.ServeHTTP(c.Writer, c.Request)
	})

	p := playground.Handler("ROC Application API", "/enterprise-query")
	router.GET("/enterprise-playground", func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	})
}
