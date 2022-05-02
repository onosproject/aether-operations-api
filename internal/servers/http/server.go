// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/api/swagger"
	generated "github.com/onosproject/scaling-umbrella/gen/graph"
	"github.com/onosproject/scaling-umbrella/internal/graph/resolvers"
	rocGrpcServer "github.com/onosproject/scaling-umbrella/internal/servers/grpc"
	"github.com/onosproject/scaling-umbrella/internal/stores/application"
	"github.com/onosproject/scaling-umbrella/internal/stores/device"
	"github.com/onosproject/scaling-umbrella/internal/stores/enterprise"
	"github.com/onosproject/scaling-umbrella/internal/stores/site"
	"github.com/onosproject/scaling-umbrella/internal/stores/slice"
	"google.golang.org/grpc"
	"io/fs"
	"net/http"
	"strings"
	"sync"
)

var log = logging.GetLogger("RestServer")

type RocHttpServer struct {
	doneCh     chan bool
	wg         *sync.WaitGroup
	address    string
	grpcConn   *grpc.ClientConn
	grpcServer *rocGrpcServer.RocApiGrpcServer
	gin        *gin.Engine
	mux        *runtime.ServeMux
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	realpath := strings.ReplaceAll(path, prefix, "")
	if realpath == "" {
		realpath = "/"
	}
	_, err := e.Open(realpath)
	if err != nil {
		return false
	}
	return true
}

func getStaticOapiFiles() (static.ServeFileSystem, error) {
	files, err := fs.Sub(swagger.OpenAPI, "dist")
	if err != nil {
		return nil, err
	}
	return embedFileSystem{FileSystem: http.FS(files)}, err
}

func (s RocHttpServer) RegisterRestGatewayHandlers() error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := enterprise.RegisterGatewayHandler(ctx, s.mux, s.grpcConn); err != nil {
		return err
	}
	if err := application.RegisterGatewayHandler(ctx, s.mux, s.grpcConn); err != nil {
		return err
	}
	if err := site.RegisterGatewayHandler(ctx, s.mux, s.grpcConn); err != nil {
		return err
	}
	if err := device.RegisterGatewayHandler(ctx, s.mux, s.grpcConn); err != nil {
		return err
	}
	if err := slice.RegisterGatewayHandler(ctx, s.mux, s.grpcConn); err != nil {
		return err
	}

	// map the gRPC gateway to GIN
	// note that /api/v1 must match the endpoint declaration in the protos
	// if that's not possible we'll have to look into rewriting the paths
	// within the handler
	s.gin.Group("/api/v1/*any").Any("", gin.WrapH(s.mux))
	return nil
}

// Defining the Graphql handler
func (s RocHttpServer) graphqlHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(generated.NewExecutableSchema(resolvers.NewResolver(s.grpcServer)))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func (s RocHttpServer) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (s RocHttpServer) RegisterGraphqlHandlers() {
	s.gin.POST("/graphql", s.graphqlHandler())
	s.gin.GET("/graphiql", s.playgroundHandler())
}

func (s RocHttpServer) StartHttpServer() error {

	if err := s.RegisterRestGatewayHandlers(); err != nil {
		return err
	}

	s.RegisterGraphqlHandlers()

	go func() {
		log.Infof("REST API server listening on %s", s.address)

		if err := s.gin.Run(s.address); err != nil {
			log.Errorf("Could not start API server: %v", err)
			return
		}
	}()

	return nil

}

func NewHttpServer(doneCh chan bool, wg *sync.WaitGroup, address string, grpcAddress string, grpcServer *rocGrpcServer.RocApiGrpcServer) (*RocHttpServer, error) {

	// create a Mux server (required by grpc-gateway)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// create a gRPC connection to our inter server to proxy requests from the gateway
	conn, err := grpc.Dial(grpcAddress, opts...)
	if err != nil {
		return nil, err
	}

	// create a Gin Server to handle both the Gateway and GraphQL requests
	// NOTE consider https://chenyitian.gitbooks.io/gin-web-framework/content/docs/38.html
	// for graceful shutdowns
	server := gin.New()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"content-type"},
		AllowCredentials: true,
	}))
	server.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"graphiql":      fmt.Sprintf("http://%s/graphiql", address),
			"graphql":       fmt.Sprintf("http://%s/graphql", address),
			"v1":            fmt.Sprintf("http://%s/api/v1", address),
			"openapi-specs": fmt.Sprintf("http://%s/docs", address),
		})
	})
	server.Use(gin.Logger()) // NOTE we might want to replace with a custom logger that uses our format

	//serve the OpenAPI specs
	oapiFiles, err := getStaticOapiFiles()
	if err != nil {
		return nil, fmt.Errorf("cannot-get-oapi-files: %s", err)
	}
	server.Use(static.Serve("/docs", oapiFiles))
	health := server.Group("/health")
	{
		health.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	srv := RocHttpServer{
		doneCh:     doneCh,
		wg:         wg,
		address:    address,
		grpcConn:   conn,
		grpcServer: grpcServer,
		gin:        server,
		mux:        mux,
	}

	return &srv, nil
}
