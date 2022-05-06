// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package endpoints

import (
	"context"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/gen/go/endpoints/v1"
	"github.com/onosproject/scaling-umbrella/internal/utils"
)

var log = logging.GetLogger("Endpoint")

func EndpointsToGnmi(ctx context.Context, endpoints []*v1.Endpoint) (*types.ApplicationEndpointList, error) {
	gnmiEps := types.ApplicationEndpointList{}

	for _, ep := range endpoints {
		pe := int(ep.PortEnd)
		ps := int(ep.PortStart)
		gnmiEp := types.ApplicationEndpoint{
			DisplayName: &ep.Name,
			EndpointId:  types.ListKey(ep.EndpointId),
			Mbr: &types.ApplicationEndpointMbr{
				Downlink: &ep.Mbr.Downlink,
				Uplink:   &ep.Mbr.Uplink,
			},
			PortEnd:   &pe,
			PortStart: &ps,
			Protocol:  &ep.Protocol,
		}
		if ep.Description != "" {
			gnmiEp.Description = &ep.Description
		}
		if ep.TrafficClass != "" {
			gnmiEp.TrafficClass = &ep.TrafficClass
		}
		gnmiEps = append(gnmiEps, gnmiEp)
	}

	log.Infow("gnmi-encoded-endpoints", "endpoints", gnmiEps)

	return &gnmiEps, nil
}

func FromGnmi(gnmiEp *types.ApplicationEndpointList) ([]*v1.Endpoint, error) {
	eps := []*v1.Endpoint{}

	for _, ep := range *gnmiEp {
		eps = append(eps, &v1.Endpoint{
			EndpointId:  string(ep.EndpointId),
			Description: utils.PointerToString(ep.Description),
			Name:        utils.PointerToString(ep.DisplayName),
			Mbr: &v1.MBR{
				Uplink:   utils.PointerToInt64(ep.Mbr.Uplink),
				Downlink: utils.PointerToInt64(ep.Mbr.Downlink),
			},
			PortStart: utils.IntPointerToInt32(ep.PortStart),
			PortEnd:   utils.IntPointerToInt32(ep.PortEnd),
			Protocol:  utils.PointerToString(ep.Protocol),
		})
	}
	return eps, nil
}
