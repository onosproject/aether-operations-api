/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package stores

import (
	"github.com/onosproject/scaling-umbrella/internal/config"
	"github.com/onosproject/scaling-umbrella/internal/datasources"
	"github.com/onosproject/scaling-umbrella/internal/stores/application"
	"github.com/onosproject/scaling-umbrella/internal/stores/device"
	"github.com/onosproject/scaling-umbrella/internal/stores/devicegroup"
	"github.com/onosproject/scaling-umbrella/internal/stores/enterprise"
	"github.com/onosproject/scaling-umbrella/internal/stores/simcard"
	"github.com/onosproject/scaling-umbrella/internal/stores/site"
	"github.com/onosproject/scaling-umbrella/internal/stores/slice"
	"github.com/onosproject/scaling-umbrella/internal/stores/smallcell"
)

type Stores struct {
	Enterprise  *enterprise.EnterpriseHandler
	Application *application.ApplicationHandler
	Site        *site.SiteHandler
	Device      *device.DeviceHandler
	SimCard     *simcard.SimCardHandler
	DeviceGroup *devicegroup.DeviceGroupHandler
	Slice       *slice.SliceHandler
	SmallCell   *smallcell.SmallCellHandler
}

// RegisterStores will create one Store per supported resource
// and makes them available to the servers
func RegisterStores(ds *datasources.Datasources, cfg *config.Config) (*Stores, error) {
	return &Stores{
		Enterprise:  enterprise.NewEnterpriseHandler(ds.OnosTopo),
		Application: application.NewApplicationHandler(ds.AetherModels_2_1_0, cfg.DataSources.OnosConfigTimeout),
		Site:        site.NewSiteHandler(ds.AetherModels_2_1_0, cfg.DataSources.OnosConfigTimeout),
		Device:      device.NewDeviceHandler(ds.AetherModels_2_1_0, cfg.DataSources.OnosConfigTimeout),
		SimCard:     simcard.NewSimCardHandler(ds.AetherModels_2_1_0, cfg.DataSources.OnosConfigTimeout),
		DeviceGroup: devicegroup.NewDeviceGroupHandler(ds.AetherModels_2_1_0, cfg.DataSources.OnosConfigTimeout),
		Slice:       slice.NewSliceHandler(ds.AetherModels_2_1_0, cfg.DataSources.OnosConfigTimeout),
		SmallCell:   smallcell.NewSmallCellHandler(ds.AetherModels_2_1_0, cfg.DataSources.OnosConfigTimeout),
	}, nil
}
