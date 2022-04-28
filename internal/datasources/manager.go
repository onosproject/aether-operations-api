/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package datasources

import (
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/scaling-umbrella/internal/config"
	onos_config "github.com/onosproject/scaling-umbrella/internal/datasources/onos-config"
	onos_topo "github.com/onosproject/scaling-umbrella/internal/datasources/onos-topo"
)

type Datasources struct {
	OnosConfig *onos_config.GnmiManager
	OnosTopo   topo.TopoClient
}

func RegisterDatasources(cfg *config.Config) (*Datasources, error) {

	onosConfigDs, err := onos_config.NewOnosConfigClient(cfg.DataSources.OnosConfigAddress)
	if err != nil {
		return nil, err
	}

	onosTopoDs, err := onos_topo.NewOnosTopoClient(cfg.DataSources.OnosTopoAddress)
	if err != nil {
		return nil, err
	}

	return &Datasources{
		OnosConfig: onosConfigDs,
		OnosTopo:   onosTopoDs,
	}, nil

}
