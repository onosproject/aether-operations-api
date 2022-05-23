/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package smallcell

import (
	"context"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/v1"
	onos_config "github.com/onosproject/scaling-umbrella/internal/datasources/onos-config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

var log = logging.GetLogger("SmallCell")

type SmallCellHandler struct {
	aether21 *aether_2_1_0.ServerImpl
}

func (a *SmallCellHandler) ListSmallCells(enterpriseId string, siteId string) (*v1.GetSmallCellsResponse, error) {
	log.Debug("listing-smallcells")
	ctx, cancel := context.WithTimeout(context.Background(), a.aether21.GnmiTimeout)
	defer cancel()

	//const enterpriseId = "acme" // NOTE this needs to be the same as the defaultTarget in sdcore-adapter
	response, err := a.aether21.GnmiGetSiteSmallCellList(ctx, "/aether/v2.1.x/{enterprise-id}/site/{site-id}/small-cell", types.EnterpriseId(enterpriseId), siteId)

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "smallcells-not-found")
	}

	return FromGnmi(response)
}

func FromGnmi(gnmiSmallCells *types.SiteSmallCellList) (*v1.GetSmallCellsResponse, error) {
	smc := v1.GetSmallCellsResponse{
		SmallCells: []*v1.SmallCell{},
	}

	for _, a := range *gnmiSmallCells {

		smc.SmallCells = append(smc.SmallCells, &v1.SmallCell{
			Id:          string(a.SmallCellId),
			Name:        *a.DisplayName,
			Description: *a.Description,
		})

	}
	return &smc, nil
}

func NewSmallCellHandler(gnmi *onos_config.GnmiManager) *SmallCellHandler {
	return &SmallCellHandler{aether21: gnmi.Aether21}
}
