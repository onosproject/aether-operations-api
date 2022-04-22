package southbound

import (
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/roc-api/pkg/southbound/application"
	"google.golang.org/grpc"
	"os"
	"time"
)

var log = logging.GetLogger("GnmiClient")

type GnmiManager struct {
	gnmiClient *southbound.GNMIProvisioner // FIXME why not to use the standard gnmi.GnmiClient?
	address    string

	aether21 *aether_2_1_0.ServerImpl

	ApplicationHandler *application.ApplicationHandler
}

func (m GnmiManager) NewGnmiClient() (*southbound.GNMIProvisioner, error) {
	// TODO allow secure connections
	opts, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	optsWithRetry := []grpc.DialOption{
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithInterval(100 * time.Millisecond))),
	}
	optsWithRetry = append(opts, optsWithRetry...)
	gnmiConn, err := grpc.Dial(m.address, optsWithRetry...)
	if err != nil {
		log.Errorw("cannot-connect-to-gnmi-server", "err", err, "address", m.address)
		return nil, err
	}

	gnmiClient := new(southbound.GNMIProvisioner)
	err = gnmiClient.Init(gnmiConn)
	if err != nil {
		log.Error("Unable to setup GNMI provisioner", err)
		return nil, err
	}

	return gnmiClient, nil
}

func NewGnmiManager(address string) (*GnmiManager, error) {

	manager := GnmiManager{
		address: address,
	}

	client, err := manager.NewGnmiClient()
	if err != nil {
		return nil, err
	}

	manager.gnmiClient = client

	manager.aether21 = &aether_2_1_0.ServerImpl{
		GnmiClient:  client,
		GnmiTimeout: time.Duration(10 * time.Second),
	}

	manager.ApplicationHandler = application.NewApplicationHandler(client, manager.aether21)
	return &manager, nil

}
