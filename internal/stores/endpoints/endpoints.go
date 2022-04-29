// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package endpoints

import (
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	v1 "github.com/onosproject/scaling-umbrella/api/v1"
	"github.com/onosproject/scaling-umbrella/internal/utils"
)

func FromGnmi(gnmiEp *types.ApplicationEndpointList) ([]*v1.Endpoint, error) {
	eps := []*v1.Endpoint{}
	for _, ep := range *gnmiEp {
		eps = append(eps, &v1.Endpoint{
			ID:          string(ep.EndpointId),
			Description: utils.PointerToString(ep.Description),
			DisplayName: utils.PointerToString(ep.DisplayName),
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
