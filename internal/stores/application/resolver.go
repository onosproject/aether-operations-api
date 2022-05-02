/*
* SPDX-FileCopyrightText: $today.year-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
 */

package application

import (
	"context"

	//generated "github.com/onosproject/scaling-umbrella/gen/graph"
	"github.com/onosproject/scaling-umbrella/gen/go/applications/v1"
)

type ApplicationResolver struct {
	grpcServer v1.ApplicationServiceServer
}

//type ApplicationServiceResolvers struct {
//	Service v1.ApplicationServiceServer
//}

//type ApplicationInput = v1.Application
//type ApplicationsInput = v1.Applications
//type EnterpriseIdInput = v1.EnterpriseId

func (a *ApplicationResolver) ApplicationServiceGetApplications(ctx context.Context, in *v1.EnterpriseId) (*v1.Applications, error) {
	return a.grpcServer.GetApplications(ctx, in)
}
func (a *ApplicationResolver) ApplicationServiceCreateApplication(ctx context.Context, in *v1.Application) (*v1.Application, error) {
	return a.grpcServer.CreateApplication(ctx, in)
}

//func (a *ApplicationResolver) Mutation() generated.MutationResolver {
//	return ApplicationResolver{
//		grpcServer: a.grpcServer,
//	}
//}
//
//func (a *ApplicationResolver) Query() generated.QueryResolver {
//	return &ApplicationServiceResolvers{
//		Service: a.grpcServer,
//	}
//}

func NewApplicationResolver(grpcServer v1.ApplicationServiceServer) *ApplicationResolver {
	return &ApplicationResolver{
		grpcServer: grpcServer,
	}
}

//func RegisterGraphQlHandler(server v1.ApplicationServiceServer, router *gin.Engine) {
//	r := NewApplicationResolver(server)
//	s := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
//		Resolvers: r,
//	}))
//
//	router.POST("/application-query", func(c *gin.Context) {
//		s.ServeHTTP(c.Writer, c.Request)
//	})
//
//	p := playground.Handler("ROC Application API", "/application-query")
//	router.GET("/application-playground", func(c *gin.Context) {
//		p.ServeHTTP(c.Writer, c.Request)
//	})
//}
