/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package application

import (
	"context"
	aether_models "github.com/onosproject/aether-models/models/aether-2.1.x/v2/api"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/gen/go/v1"
	"github.com/onosproject/scaling-umbrella/internal/stores/endpoints"
	"time"
)

var log = logging.GetLogger("Application")

type ApplicationHandler struct {
	aether21 *aether_models.GnmiClient
	timeout  time.Duration
}

func (a *ApplicationHandler) ListApplications(enterpriseId string) (*v1.GetApplicationsResponse, error) {
	log.Debug("listing-application")
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	res, err := a.aether21.GetApplication(ctx, enterpriseId)
	if err != nil {
		return nil, err
	}

	return FromGnmi(res)
}

func FromGnmi(gnmiApps map[string]*aether_models.OnfApplication_Application) (*v1.GetApplicationsResponse, error) {
	apps := &v1.GetApplicationsResponse{
		Applications: []*v1.Application{},
	}

	for _, a := range gnmiApps {

		eps, err := endpoints.FromGnmiClient(a.Endpoint)
		if err != nil {
			return nil, err
		}

		apps.Applications = append(apps.Applications, &v1.Application{
			Id:          *a.ApplicationId,
			Description: *a.Description,
			Endpoints:   eps,
		})
	}

	return apps, nil
}

func NewApplicationHandler(gnmi *aether_models.GnmiClient, timeout time.Duration) *ApplicationHandler {
	return &ApplicationHandler{
		aether21: gnmi,
		timeout:  timeout,
	}
}
