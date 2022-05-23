/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package stores

import (
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
func RegisterStores(ds *datasources.Datasources) (*Stores, error) {
	return &Stores{
		Enterprise:  enterprise.NewEnterpriseHandler(ds.OnosTopo),
		Application: application.NewApplicationHandler(ds.OnosConfig),
		Site:        site.NewSiteHandler(ds.OnosConfig),
		Device:      device.NewDeviceHandler(ds.OnosConfig),
		SimCard:     simcard.NewSimCardHandler(ds.OnosConfig),
		DeviceGroup: devicegroup.NewDeviceGroupHandler(ds.OnosConfig),
		Slice:       slice.NewSliceHandler(ds.OnosConfig),
		SmallCell:   smallcell.NewSmallCellHandler(ds.OnosConfig),
	}, nil
}
