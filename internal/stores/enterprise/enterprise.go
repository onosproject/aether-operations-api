/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package enterprise

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/gen/go/enterprises/v1"
)

var log = logging.GetLogger("Enterprise")

type EnterpriseHandler struct {
	client topo.TopoClient
}

func (e EnterpriseHandler) ListEnterprises() (*v1.Enterprises, error) {

	// TODO better understand the Enterprise info stored in TOPO
	// seems like we're only storing basic things:
	// ID: acme
	// Kind ID: aether
	// Labels: <None>
	// Source Id's:
	// Target Id's: uuid:7ccafdaf-350c-40d6-9335-fd8dfbd6a512
	// Aspects:
	// - onos.topo.Configurable={"address":"sdcore-adapter-v2-1:5150","type":"aether","version":"2.1.x"}
	// - onos.topo.Asset={"name":"ACME Enterprise"}
	// - onos.topo.TLSOptions={"insecure":true}
	// - onos.topo.MastershipState={"term":"1","nodeId":"uuid:7ccafdaf-350c-40d6-9335-fd8dfbd6a512"}
	// - onos.topo.Location={"lat":52.515,"lng":13.3885}

	log.Debug("listing-enterprises")
	res, err := e.client.List(context.Background(), &topo.ListRequest{
		Filters: &topo.Filters{
			KindFilter: &topo.Filter{
				Filter: &topo.Filter_In{In: &topo.InFilter{Values: []string{"aether"}}},
			},
			WithAspects: []string{"onos.topo.Location"},
			ObjectTypes: []topo.Object_Type{topo.Object_ENTITY},
		},
		SortOrder: topo.SortOrder_ASCENDING,
	})

	if err != nil {
		return nil, err
	}

	// TODO move this in a FromGrpc method to convert the format
	enterprises := &v1.Enterprises{Enterprises: []*v1.Enterprise{}}

	for _, o := range res.Objects {

		asset := &topo.Asset{}
		err := o.GetAspect(asset)
		if err != nil {
			return nil, err
		}

		e := &v1.Enterprise{
			EnterpriseId: string(o.ID),
			Name:         asset.Name,
		}
		enterprises.Enterprises = append(enterprises.Enterprises, e)
	}

	return enterprises, nil
}

func NewEnterpriseHandler(client topo.TopoClient) *EnterpriseHandler {
	return &EnterpriseHandler{
		client: client,
	}
}
