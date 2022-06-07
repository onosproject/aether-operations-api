/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package site

import (
	"context"
	aether_models "github.com/onosproject/aether-models/models/aether-2.1.x/v2/api"
	v1 "github.com/onosproject/aether-operations-api/gen/go/v1"
	"github.com/onosproject/aether-operations-api/internal/stores/device"
	"github.com/onosproject/aether-operations-api/internal/stores/slice"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"time"
)

var log = logging.GetLogger("Application")

type SiteHandler struct {
	aether21 *aether_models.GnmiClient
	timeout  time.Duration
}

func (a *SiteHandler) ListSites(enterpriseId string) (*v1.GetSitesResponse, error) {
	log.Debug("listing-sites")
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	res, err := a.aether21.GetSite(ctx, enterpriseId)
	if err != nil {
		return nil, err
	}

	return FromGnmi(res)
}

func FromGnmi(gnmiSites map[string]*aether_models.OnfSite_Site) (*v1.GetSitesResponse, error) {
	sites := v1.GetSitesResponse{
		Sites: []*v1.Site{},
	}

	for _, a := range gnmiSites {

		d, err := device.FromGnmi(a.Device)
		if err != nil {
			return nil, err
		}

		s, err := slice.FromGnmi(a.Slice)
		if err != nil {
			return nil, err
		}

		//s, err := smallcell.FromGnmi(a.Slice)
		//if err != nil {
		//	return nil, err
		//}

		sites.Sites = append(sites.Sites, &v1.Site{
			Id:          *a.SiteId,
			Name:        *a.DisplayName,
			Description: *a.Description,
			Devices:     d.Devices,
			Slices:      s.Slices,
			//SmallCells: s.SmallCells,
		})
	}

	return &sites, nil
}

func NewSiteHandler(gnmi *aether_models.GnmiClient, timeout time.Duration) *SiteHandler {
	return &SiteHandler{
		aether21: gnmi,
		timeout:  timeout,
	}
}
