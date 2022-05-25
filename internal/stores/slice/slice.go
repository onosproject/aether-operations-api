/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package slice

import (
	"context"
	"fmt"
	aether_models "github.com/onosproject/aether-models/models/aether-2.1.x/v2/api"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/v1"
	"github.com/onosproject/scaling-umbrella/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var log = logging.GetLogger("Slice")

type SliceHandler struct {
	aether21 *aether_models.GnmiClient
	timeout  time.Duration
}

func (a *SliceHandler) ListSlices(enterpriseId string, siteId string) (*v1.GetSlicesResponse, error) {
	log.Debug("listing-slices")
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	// FIXME we need to augment the gnmi-client-gen tool to generated methods to get items from nested lists
	// AETHER-3554 is needed for this
	res, err := a.aether21.GetSite(ctx, enterpriseId)
	if err != nil {
		return nil, err
	}

	for id, s := range res {
		if siteId == id {
			return FromGnmi(s.Slice)
		}
	}

	return nil, status.Error(codes.NotFound, fmt.Sprintf("site-%s-not-found", siteId))
}

func FromGnmi(gnmiSlices map[string]*aether_models.OnfSite_Site_Slice) (*v1.GetSlicesResponse, error) {
	slices := v1.GetSlicesResponse{
		Slices: []*v1.Slice{},
	}

	for _, a := range gnmiSlices {

		slices.Slices = append(slices.Slices, &v1.Slice{
			Id:          utils.PointerToString(a.SliceId),
			Name:        utils.PointerToString(a.DisplayName),
			Description: utils.PointerToString(a.Description),
		})

	}
	return &slices, nil
}

func NewSliceHandler(gnmi *aether_models.GnmiClient, timeout time.Duration) *SliceHandler {
	return &SliceHandler{
		aether21: gnmi,
		timeout:  timeout,
	}
}
