package onos_config

import (
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"time"
)

var log = logging.GetLogger("OnosConfigDatasource")

func NewOnosConfigClient(address string) (*southbound.GNMIProvisioner, error) {
	log.Infow("initializing-onos-config-client", "address", address)
	// TODO handle secure connections
	opts, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
		return nil, err
	}

	optsWithRetry := []grpc.DialOption{
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithInterval(100 * time.Millisecond))),
	}
	optsWithRetry = append(opts, optsWithRetry...)

	gnmiConn, err := grpc.Dial(address, optsWithRetry...)
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
