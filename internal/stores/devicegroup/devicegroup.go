/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package devicegroup

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

var log = logging.GetLogger("DeviceGroup")

type DeviceGroupHandler struct {
	aether21 *aether_models.GnmiClient
	timeout  time.Duration
}

func (a *DeviceGroupHandler) ListDeviceGroups(enterpriseId string, siteId string) (*v1.GetDeviceGroupsResponse, error) {
	log.Debug("listing-DeviceGroups")
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
			return FromGnmi(s.DeviceGroup)
		}
	}

	return nil, status.Error(codes.NotFound, fmt.Sprintf("site-%s-not-found", siteId))
}

func FromGnmi(gnmiDeviceGroups map[string]*aether_models.OnfSite_Site_DeviceGroup) (*v1.GetDeviceGroupsResponse, error) {
	deviceGroups := v1.GetDeviceGroupsResponse{
		DeviceGroups: []*v1.DeviceGroup{},
	}

	for _, a := range gnmiDeviceGroups {
		deviceGroups.DeviceGroups = append(deviceGroups.DeviceGroups, &v1.DeviceGroup{
			Id:          utils.PointerToString(a.DeviceGroupId),
			Name:        utils.PointerToString(a.DisplayName),
			Description: utils.PointerToString(a.Description),
		})
	}

	return &deviceGroups, nil
}

func NewDeviceGroupHandler(gnmi *aether_models.GnmiClient, timeout time.Duration) *DeviceGroupHandler {
	return &DeviceGroupHandler{
		aether21: gnmi,
		timeout:  timeout,
	}
}
