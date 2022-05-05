/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
 */

package application

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/applications/v1"
	application "github.com/onosproject/scaling-umbrella/gen/graph/applications/v1"
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

func (a *ApplicationResolver) Application() application.ApplicationResolver {
	return nil
}

func (a *ApplicationResolver) ApplicationInput() application.ApplicationInputResolver {
	return nil
}

func NewApplicationResolver(grpcServer v1.ApplicationServiceServer) *ApplicationResolver {
	return &ApplicationResolver{
		grpcServer: grpcServer,
	}
}

func RegisterGraphQlHandler(server v1.ApplicationServiceServer, router *gin.Engine) {
	r := NewApplicationResolver(server)
	s := handler.NewDefaultServer(application.NewExecutableSchema(application.Config{
		Resolvers: r,
	}))

	router.POST("/application-query", func(c *gin.Context) {
		s.ServeHTTP(c.Writer, c.Request)
	})

	p := playground.Handler("ROC Application API", "/application-query")
	router.GET("/application-playground", func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	})
}
