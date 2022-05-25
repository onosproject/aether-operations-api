/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package datasources

import (
	aether_models "github.com/onosproject/aether-models/models/aether-2.1.x/v2/api"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/scaling-umbrella/internal/config"
	onosConfig "github.com/onosproject/scaling-umbrella/internal/datasources/onos-config"
	onosTopo "github.com/onosproject/scaling-umbrella/internal/datasources/onos-topo"
)

type Datasources struct {
	OnosConfig         *onosConfig.GnmiManager
	OnosTopo           topo.TopoClient
	AetherModels_2_1_0 *aether_models.GnmiClient
}

// RegisterDatasources will initialize all the required
// southbound datasources and store them so that they are available to the Stores
func RegisterDatasources(cfg *config.Config) (*Datasources, error) {

	onosConfigDs, err := onosConfig.NewOnosConfigClient(cfg.DataSources.OnosConfigAddress)
	if err != nil {
		return nil, err
	}

	onosTopoDs, err := onosTopo.NewOnosTopoClient(cfg.DataSources.OnosTopoAddress)
	if err != nil {
		return nil, err
	}

	aether210Ds, err := onosConfig.NewAetherModels_2_1_0_Client(cfg.DataSources.OnosConfigAddress)
	if err != nil {
		return nil, err
	}

	// TODO: prometheus connection

	return &Datasources{
		OnosConfig:         onosConfigDs,
		OnosTopo:           onosTopoDs,
		AetherModels_2_1_0: aether210Ds,
		//Prometheus:   TODO,
	}, nil

}
