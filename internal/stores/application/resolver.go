/*
* SPDX-FileCopyrightText: $today.year-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
 */

package application

//type ApplicationResolver struct {
//	grpcServer v1.ApplicationServiceServer
//}
//
//func (a *ApplicationResolver) Mutation() generated.MutationResolver {
//	return ApplicationResolver{
//		grpcServer: a.grpcServer,
//	}
//}
//
//func (a *ApplicationResolver) Mutation() application.MutationResolver {
//	return &v1.ApplicationServiceResolvers{
//		Service: a.grpcServer,
//	}
//}
//
//func NewApplicationResolver(grpcServer v1.ApplicationServiceServer) *ApplicationResolver {
//	return &ApplicationResolver{
//		grpcServer: grpcServer,
//	}
//}
//
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
