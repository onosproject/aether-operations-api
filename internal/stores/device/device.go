/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package device

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

var log = logging.GetLogger("Device")

type DeviceHandler struct {
	aether21 *aether_models.GnmiClient
	timeout  time.Duration
}

func (a *DeviceHandler) ListDevices(enterpriseId string, siteId string) (*v1.GetDevicesResponse, error) {
	log.Debug("listing-devices")
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
			return FromGnmi(s.Device)
		}
	}

	return nil, status.Error(codes.NotFound, fmt.Sprintf("site-%s-not-found", siteId))
}

func FromGnmi(gnmiDevices map[string]*aether_models.OnfSite_Site_Device) (*v1.GetDevicesResponse, error) {
	devices := v1.GetDevicesResponse{
		Devices: []*v1.Device{},
	}
	for _, a := range gnmiDevices {

		devices.Devices = append(devices.Devices, &v1.Device{
			Id:          utils.PointerToString(a.DeviceId),
			Name:        utils.PointerToString(a.DisplayName),
			Description: utils.PointerToString(a.Description),
			Attached:    "",
			Imei:        utils.PointerToString(a.Imei),
			Ip:          "",
			// TODO simcards are represented as string
			//SimCard:     a.SimCard,
			//DeviceGroups: nil,
		})
	}
	return &devices, nil
}

func NewDeviceHandler(gnmi *aether_models.GnmiClient, timeout time.Duration) *DeviceHandler {
	return &DeviceHandler{
		aether21: gnmi,
		timeout:  timeout,
	}
}
