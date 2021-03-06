package onos_config

import (
	"crypto/tls"
	"fmt"
	aether_models "github.com/onosproject/aether-models/models/aether-2.1.x/v2/api"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

var log = logging.GetLogger("OnosConfigDatasource")

type GnmiManager struct {
	gnmiClient *southbound.GNMIProvisioner // FIXME why not to use the standard gnmi.GnmiClient?
	address    string

	Aether21 *aether_2_1_0.ServerImpl
}

func (m GnmiManager) newGnmiClient() (*southbound.GNMIProvisioner, error) {
	log.Infow("initializing-onos-config-client", "address", m.address)
	// TODO handle secure connections
	opts, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
		return nil, err
	}

	optsWithRetry := []grpc.DialOption{
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithInterval(100 * time.Millisecond))),
	}
	optsWithRetry = append(opts, optsWithRetry...)

	gnmiConn, err := grpc.Dial(m.address, optsWithRetry...)
	if err != nil {
		return nil, fmt.Errorf("cannot-connect-to-gnmi-server: %s", err)
	}

	gnmiClient := new(southbound.GNMIProvisioner)
	err = gnmiClient.Init(gnmiConn)
	if err != nil {
		return nil, fmt.Errorf("unable-to-setup-gnmi-provisioner: %s", err)
	}

	return gnmiClient, nil
}

func NewOnosConfigClient(address string) (*GnmiManager, error) {
	manager := GnmiManager{
		address: address,
	}

	client, err := manager.newGnmiClient()
	if err != nil {
		return nil, err
	}

	manager.gnmiClient = client

	manager.Aether21 = &aether_2_1_0.ServerImpl{
		GnmiClient:  client,
		GnmiTimeout: 10 * time.Second,
	}

	return &manager, nil
}

func NewAetherModels_2_1_0_Client(address string) (*aether_models.GnmiClient, error) {
	// TODO handle secure connections
	cert, err := tls.X509KeyPair([]byte(certs.DefaultClientCrt), []byte(certs.DefaultClientKey))
	if err != nil {
		return nil, err
	}
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			}),
		),
	}
	gnmiConn, err := grpc.Dial(address, options...)
	if err != nil {
		return nil, err
	}

	return aether_models.NewAetherGnmiClient(gnmiConn), nil
}
