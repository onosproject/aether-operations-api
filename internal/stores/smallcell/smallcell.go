/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package smallcell

import (
	"context"
	"fmt"
	aether_models "github.com/onosproject/aether-models/models/aether-2.1.x/api"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/v1"
	"github.com/onosproject/scaling-umbrella/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var log = logging.GetLogger("SmallCell")

type SmallCellHandler struct {
	aether21 *aether_models.GnmiClient
	timeout  time.Duration
}

func (a *SmallCellHandler) ListSmallCells(enterpriseId string, siteId string) (*v1.GetSmallCellsResponse, error) {
	log.Debug("listing-smallcells")
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	res, err := a.aether21.GetSite(ctx, enterpriseId)
	if err != nil {
		return nil, err
	}

	for id, s := range res {
		if siteId == id {
			return FromGnmi(s.SmallCell)
		}
	}

	return nil, status.Error(codes.NotFound, fmt.Sprintf("site-%s-not-found", siteId))
}

func FromGnmi(gnmiSmallCells map[string]*aether_models.OnfSite_Site_SmallCell) (*v1.GetSmallCellsResponse, error) {
	smc := v1.GetSmallCellsResponse{
		SmallCells: []*v1.SmallCell{},
	}

	for _, a := range gnmiSmallCells {

		smc.SmallCells = append(smc.SmallCells, &v1.SmallCell{
			Id:          utils.PointerToString(a.SmallCellId),
			Name:        utils.PointerToString(a.DisplayName),
			Description: utils.PointerToString(a.Description),
		})

	}
	return &smc, nil
}

func NewSmallCellHandler(gnmi *aether_models.GnmiClient, timeout time.Duration) *SmallCellHandler {
	return &SmallCellHandler{
		aether21: gnmi,
		timeout:  timeout,
	}
}
