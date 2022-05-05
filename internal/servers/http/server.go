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
	generated "github.com/onosproject/scaling-umbrella/gen/graph/graphql"
	"github.com/onosproject/scaling-umbrella/gen/openapiv2"
	"github.com/onosproject/scaling-umbrella/internal/config"
	"github.com/onosproject/scaling-umbrella/internal/graph/resolvers"
	rocGrpcServer "github.com/onosproject/scaling-umbrella/internal/servers/grpc"
	"github.com/onosproject/scaling-umbrella/internal/stores/application"
	"github.com/onosproject/scaling-umbrella/internal/stores/device"
	"github.com/onosproject/scaling-umbrella/internal/stores/enterprise"
	"github.com/onosproject/scaling-umbrella/internal/stores/simcard"
	"github.com/onosproject/scaling-umbrella/internal/stores/site"
	"github.com/onosproject/scaling-umbrella/internal/stores/slice"
	"github.com/onosproject/scaling-umbrella/internal/utils"
	"google.golang.org/grpc"
	"io/fs"
	"net/http"
	"strings"
	"sync"
	"time"
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
	files, err := fs.Sub(openapiv2.OpenAPI, "dist")
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
	if err := simcard.RegisterGatewayHandler(ctx, s.mux, s.grpcConn); err != nil {
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

	// TODO it would be good to collect the registered endpoinds to
	// dinamically generate a navigation page
	application.RegisterGraphQlHandler(s.grpcServer.Services.ApplicationService, s.gin)
	enterprise.RegisterGraphQlHandler(s.grpcServer.Services.EnterpriseService, s.gin)
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

func NewHttpServer(doneCh chan bool, wg *sync.WaitGroup, config config.Config, grpcServer *rocGrpcServer.RocApiGrpcServer) (*RocHttpServer, error) {

	// create a Mux server (required by grpc-gateway)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// create a gRPC connection to our inter server to proxy requests from the gateway
	conn, err := grpc.Dial(config.ServersConfig.GrpcAddress, opts...)
	if err != nil {
		return nil, err
	}

	// create a Gin Server to handle both the Gateway and GraphQL requests
	// NOTE consider https://chenyitian.gitbooks.io/gin-web-framework/content/docs/38.html
	// for graceful shutdowns
	server := gin.New()

	corsOrigins := append(config.ServersConfig.Cors, "*")

	server.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins,
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"content-type"},
		AllowCredentials: true,
	}))
	server.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"graphiql":      fmt.Sprintf("http://%s/graphiql", config.ServersConfig.HttpAddress),
			"graphql":       fmt.Sprintf("http://%s/graphql", config.ServersConfig.HttpAddress),
			"v1":            fmt.Sprintf("http://%s/api/v1", config.ServersConfig.HttpAddress),
			"openapi-specs": fmt.Sprintf("http://%s/docs", config.ServersConfig.HttpAddress),
		})
	})

	server.Use(func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := utils.GetDurationInMillseconds(start)

		fields := []interface{}{
			"duration", duration,
			"client_ip", utils.GetClientIP(c),
			"method", c.Request.Method,
			"path", c.Request.RequestURI,
			"status", c.Writer.Status(),
			"referrer", c.Request.Referer(),
		}

		if c.Writer.Status() >= 500 {
			log.Errorw(c.Errors.String(), fields...)
		} else if c.Writer.Status() >= 400 {
			log.Warnw("", fields...)
		} else {
			log.Debugw("", fields...)
		}
	})

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
		address:    config.ServersConfig.HttpAddress,
		grpcConn:   conn,
		grpcServer: grpcServer,
		gin:        server,
		mux:        mux,
	}

	return &srv, nil
}
