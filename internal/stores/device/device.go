/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package device

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

var log = logging.GetLogger("Device")

type DeviceHandler struct {
	aether21 *aether_2_1_0.ServerImpl
}

func (a *DeviceHandler) ListDevices(enterpriseId string, siteId string) (*v1.GetDevicesResponse, error) {
	log.Debug("listing-devices")
	ctx, cancel := context.WithTimeout(context.Background(), a.aether21.GnmiTimeout)
	defer cancel()

	//const enterpriseId = "acme" // NOTE this needs to be the same as the defaultTarget in sdcore-adapter
	response, err := a.aether21.GnmiGetSiteDeviceList(ctx, "/aether/v2.1.x/{enterprise-id}/site/{site-id}/device", types.EnterpriseId(enterpriseId), siteId)

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "devices-not-found")
	}

	return FromGnmi(response)
}

func FromGnmi(gnmiDevices *types.SiteDeviceList) (*v1.GetDevicesResponse, error) {
	devices := v1.GetDevicesResponse{
		Devices: []*v1.Device{},
	}

	for _, a := range *gnmiDevices {

		devices.Devices = append(devices.Devices, &v1.Device{
			Id:          string(a.DeviceId),
			Name:        *a.DisplayName,
			Description: *a.Description,
			Attached:    "",
			Imei:        *a.Imei,
			Ip:          "",
			//SimCard:     a.SimCard,
			//DeviceGroups: nil,
		})
	}
	return &devices, nil
}

func NewDeviceHandler(gnmi *onos_config.GnmiManager) *DeviceHandler {
	return &DeviceHandler{aether21: gnmi.Aether21}
}
