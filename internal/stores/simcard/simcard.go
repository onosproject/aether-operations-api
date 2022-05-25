/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package simcard

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

var log = logging.GetLogger("SimCard")

type SimCardHandler struct {
	aether21 *aether_models.GnmiClient
	timeout  time.Duration
}

func (a *SimCardHandler) ListSimCards(enterpriseId string, siteId string) (*v1.GetSimCardsResponse, error) {
	log.Debug("listing-simcards")
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
			return FromGnmi(s.SimCard)
		}
	}

	return nil, status.Error(codes.NotFound, fmt.Sprintf("site-%s-not-found", siteId))
}

func FromGnmi(gnmiSimCards map[string]*aether_models.OnfSite_Site_SimCard) (*v1.GetSimCardsResponse, error) {
	simcards := v1.GetSimCardsResponse{
		SimCards: []*v1.SimCard{},
	}

	for _, a := range gnmiSimCards {

		simcards.SimCards = append(simcards.SimCards, &v1.SimCard{
			Id:          utils.PointerToString(a.SimId),
			Name:        utils.PointerToString(a.DisplayName),
			Description: utils.PointerToString(a.Description),
		})
	}
	return &simcards, nil
}

func NewSimCardHandler(gnmi *aether_models.GnmiClient, timeout time.Duration) *SimCardHandler {
	return &SimCardHandler{
		aether21: gnmi,
		timeout:  timeout,
	}
}
