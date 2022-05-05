/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package slice

import (
	"context"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/slices/v1"
	onos_config "github.com/onosproject/scaling-umbrella/internal/datasources/onos-config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

var log = logging.GetLogger("Device")

type SliceHandler struct {
	aether21 *aether_2_1_0.ServerImpl
}

func (a *SliceHandler) ListSlices(enterpriseId string, siteId string) (*v1.GetSlicesResponse, error) {
	log.Debug("listing-slices")
	ctx, cancel := context.WithTimeout(context.Background(), a.aether21.GnmiTimeout)
	defer cancel()

	//const enterpriseId = "acme" // NOTE this needs to be the same as the defaultTarget in sdcore-adapter
	response, err := a.aether21.GnmiGetSiteSliceList(ctx, "/aether/v2.1.x/{enterprise-id}/site/{site-id}/slice", types.EnterpriseId(enterpriseId), siteId)

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "slices-not-found")
	}

	return FromGnmi(response)
}

func FromGnmi(gnmiSlices *types.SiteSliceList) (*v1.GetSlicesResponse, error) {
	slices := v1.GetSlicesResponse{
		Slices: []*v1.Slice{},
	}

	for _, a := range *gnmiSlices {

		slices.Slices = append(slices.Slices, &v1.Slice{
			SliceId:     string(a.SliceId),
			Name:        *a.DisplayName,
			Description: *a.Description,
		})

	}
	return &slices, nil
}

func NewSliceHandler(gnmi *onos_config.GnmiManager) *SliceHandler {
	return &SliceHandler{aether21: gnmi.Aether21}
}
