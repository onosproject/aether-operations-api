/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package stores

import (
	"github.com/onosproject/scaling-umbrella/internal/datasources"
	"github.com/onosproject/scaling-umbrella/internal/stores/application"
	"github.com/onosproject/scaling-umbrella/internal/stores/enterprise"
)

type Stores struct {
	Application *application.ApplicationHandler
	Enterprise  *enterprise.EnterpriseHandler
}

// RegisterStores will create one Store per supported resource
// and makes them available to the servers
func RegisterStores(ds *datasources.Datasources) (*Stores, error) {
	return &Stores{
		Application: application.NewApplicationHandler(ds.OnosConfig),
		Enterprise:  enterprise.NewEnterpriseHandler(ds.OnosTopo),
	}, nil
}
