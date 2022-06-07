// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package endpoints

import (
	aether_models "github.com/onosproject/aether-models/models/aether-2.1.x/v2/api"
	"github.com/onosproject/aether-operations-api/gen/go/v1"
	"github.com/onosproject/aether-operations-api/internal/utils"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
)

// deprecated
func FromGnmi(gnmiEp *types.ApplicationEndpointList) ([]*v1.Endpoint, error) {
	eps := []*v1.Endpoint{}
	for _, ep := range *gnmiEp {
		eps = append(eps, &v1.Endpoint{
			Id:          string(ep.EndpointId),
			Description: utils.PointerToString(ep.Description),
			Name:        utils.PointerToString(ep.DisplayName),
			Mbr: &v1.MBR{
				Uplink:   utils.PointerToInt64(ep.Mbr.Uplink),
				Downlink: utils.PointerToInt64(ep.Mbr.Downlink),
			},
			PortStart: int32(*ep.PortStart),
			PortEnd:   int32(*ep.PortEnd),
			Protocol:  *ep.Protocol,
		})
	}
	return eps, nil
}

func FromGnmiClient(gnmiEp map[string]*aether_models.OnfApplication_Application_Endpoint) ([]*v1.Endpoint, error) {
	eps := []*v1.Endpoint{}
	for _, ep := range gnmiEp {
		eps = append(eps, &v1.Endpoint{
			Id:          utils.PointerToString(ep.EndpointId),
			Description: utils.PointerToString(ep.Description),
			Name:        utils.PointerToString(ep.DisplayName),
			Mbr: &v1.MBR{
				Uplink:   utils.PointerToInt64(ep.Mbr.Uplink),
				Downlink: utils.PointerToInt64(ep.Mbr.Downlink),
			},
			PortStart: int32(*ep.PortStart),
			PortEnd:   int32(*ep.PortEnd),
			Protocol:  *ep.Protocol,
		})
	}
	return eps, nil
}
