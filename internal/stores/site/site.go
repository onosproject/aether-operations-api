/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package site

import (
	"context"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/v1"
	onos_config "github.com/onosproject/scaling-umbrella/internal/datasources/onos-config"
	"github.com/onosproject/scaling-umbrella/internal/stores/device"
	"github.com/onosproject/scaling-umbrella/internal/stores/slice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

var log = logging.GetLogger("Application")

type SiteHandler struct {
	aether21 *aether_2_1_0.ServerImpl
}

func (a *SiteHandler) ListSites(enterpriseId string) (*v1.GetSitesResponse, error) {
	log.Debug("listing-sites")
	ctx, cancel := context.WithTimeout(context.Background(), a.aether21.GnmiTimeout)
	defer cancel()

	//const enterpriseId = "acme" // NOTE this needs to be the same as the defaultTarget in sdcore-adapter
	response, err := a.aether21.GnmiGetSiteList(ctx, "/aether/v2.1.x/{enterprise-id}/site", types.EnterpriseId(enterpriseId))

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "applications-not-found")
	}

	return FromGnmi(response)
}

func FromGnmi(gnmiSites *types.SiteList) (*v1.GetSitesResponse, error) {
	sites := v1.GetSitesResponse{
		Sites: []*v1.Site{},
	}

	for _, a := range *gnmiSites {

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
			Id:          string(a.SiteId),
			Name:        *a.DisplayName,
			Description: *a.Description,
			Devices:     d.Devices,
			Slices:      s.Slices,
			//SmallCells: s.SmallCells,
		})
	}

	return &sites, nil
}

func NewSiteHandler(gnmi *onos_config.GnmiManager) *SiteHandler {
	return &SiteHandler{aether21: gnmi.Aether21}
}
