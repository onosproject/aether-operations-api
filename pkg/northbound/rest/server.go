package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/gommon/log"
	"github.com/onosproject/roc-api/api/swagger"
	gw "github.com/onosproject/roc-api/api/v1"
	"google.golang.org/grpc"
	"io/fs"
	"mime"
	"net/http"
	"strings"
	"sync"
)

type RocApiRestServer struct {
	doneCh      chan bool
	wg          *sync.WaitGroup
	address     string
	grpcAddress string
}

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(swagger.OpenAPI, "dist")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

func (s RocApiRestServer) StartRestServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	serveMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := gw.RegisterRocApiHandlerFromEndpoint(ctx, serveMux, s.grpcAddress, opts); err != nil {
		log.Errorf("Could not register API server: %v", err)
		return
	}

	oa := getOpenAPIHandler()

	server := &http.Server{Addr: s.address, Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			serveMux.ServeHTTP(w, r)
			return
		}
		oa.ServeHTTP(w, r)
	})}

	go func() {
		log.Infof("REST API server listening on %s", s.address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("Could not start API server: %v", err)
			return
		}
	}()

	x := <-s.doneCh
	if x {
		log.Warnf("Stopping API REST server")
		_ = server.Shutdown(ctx)
	}

	s.wg.Done()
}

func NewRestServer(doneCh chan bool, wg *sync.WaitGroup, address string, grpcAddress string) (*RocApiRestServer, error) {
	srv := RocApiRestServer{
		doneCh:      doneCh,
		wg:          wg,
		address:     address,
		grpcAddress: grpcAddress,
	}

	return &srv, nil
}
