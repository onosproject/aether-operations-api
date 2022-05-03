/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package simcard

import (
	"context"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	v1 "github.com/onosproject/scaling-umbrella/gen/go/simcards/v1"
	onos_config "github.com/onosproject/scaling-umbrella/internal/datasources/onos-config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

var log = logging.GetLogger("SimCard")

type SimCardHandler struct {
	aether21 *aether_2_1_0.ServerImpl
}

func (a *SimCardHandler) ListSimCards(enterpriseId string, siteId string) (*v1.GetSimCardsResponse, error) {
	log.Debug("listing-simcards")
	ctx, cancel := context.WithTimeout(context.Background(), a.aether21.GnmiTimeout)
	defer cancel()

	//const enterpriseId = "acme" // NOTE this needs to be the same as the defaultTarget in sdcore-adapter
	response, err := a.aether21.GnmiGetSiteSimCardList(ctx, "/aether/v2.1.x/{enterprise-id}/site/{site-id}/sim-card", types.EnterpriseId(enterpriseId), siteId)

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "simcards-not-found")
	}

	return FromGnmi(response)
}

func FromGnmi(gnmiSimCards *types.SiteSimCardList) (*v1.GetSimCardsResponse, error) {
	simcards := v1.GetSimCardsResponse{
		SimCards: []*v1.SimCard{},
	}

	for _, a := range *gnmiSimCards {

		simcards.SimCards = append(simcards.SimCards, &v1.SimCard{
			Id:          string(a.SimId),
			Name:        *a.DisplayName,
			Description: *a.Description,
		})
	}
	return &simcards, nil
}

func NewSimCardHandler(gnmi *onos_config.GnmiManager) *SimCardHandler {
	return &SimCardHandler{aether21: gnmi.Aether21}
}
