/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package application

import (
	"context"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/gen/go/applications/v1"
	onos_config "github.com/onosproject/scaling-umbrella/internal/datasources/onos-config"
	"github.com/onosproject/scaling-umbrella/internal/stores/endpoints"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

var log = logging.GetLogger("Application")

type ApplicationHandler struct {
	aether21 *aether_2_1_0.ServerImpl
}

func (a *ApplicationHandler) ListApplications(enterpriseId string) (*v1.Applications, error) {
	log.Debug("listing-application")
	ctx, cancel := context.WithTimeout(context.Background(), a.aether21.GnmiTimeout)
	defer cancel()

	//const enterpriseId = "acme" // NOTE this needs to be the same as the defaultTarget in sdcore-adapter
	response, err := a.aether21.GnmiGetApplicationList(ctx, "/aether/v2.1.x/{enterprise-id}/application", types.EnterpriseId(enterpriseId))

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "applications-not-found")
	}

	return FromGnmi(response)
}

func FromGnmi(gnmiApps *types.ApplicationList) (*v1.Applications, error) {
	apps := v1.Applications{
		Applications: []*v1.Application{},
	}

	for _, a := range *gnmiApps {

		eps, err := endpoints.FromGnmi(a.Endpoint)

		if err != nil {
			return nil, err
		}

		apps.Applications = append(apps.Applications, &v1.Application{
			ApplicationId: string(a.ApplicationId),
			Description:   *a.Description,
			Endpoints:     eps})
	}
	return &apps, nil
}

func NewApplicationHandler(gnmi *onos_config.GnmiManager) *ApplicationHandler {
	return &ApplicationHandler{aether21: gnmi.Aether21}
}
