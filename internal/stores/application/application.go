/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package application

import (
	"context"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	roc_utils "github.com/onosproject/aether-roc-api/pkg/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/scaling-umbrella/gen/go/applications/v1"
	onos_config "github.com/onosproject/scaling-umbrella/internal/datasources/onos-config"
	"github.com/onosproject/scaling-umbrella/internal/stores/endpoints"
	"github.com/onosproject/scaling-umbrella/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

var log = logging.GetLogger("Application")

type ApplicationHandler struct {
	aether21 *aether_2_1_0.ServerImpl
}

func (a *ApplicationHandler) ListApplications(ctx context.Context, enterpriseId string) (*v1.Applications, error) {
	log.Debug("listing-applications")
	gnmiCtx, cancel := context.WithTimeout(ctx, a.aether21.GnmiTimeout)
	defer cancel()

	response, err := a.aether21.GnmiGetApplicationList(gnmiCtx, "/aether/v2.1.x/{enterprise-id}/application", types.EnterpriseId(enterpriseId))

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "applications-not-found")
	}

	return ApplicationsFromGnmi(enterpriseId, response)
}

func (a *ApplicationHandler) GetApplication(ctx context.Context, req *v1.ApplicationFilter) (*v1.Application, error) {
	log.Debug("get-application")
	gnmiCtx, cancel := context.WithTimeout(ctx, a.aether21.GnmiTimeout)
	defer cancel()

	response, err := a.aether21.GnmiGetApplication(gnmiCtx, "/aether/v2.1.x/{enterprise-id}/application/{application-id}",
		types.EnterpriseId(req.EnterpriseId), req.ApplicationId)

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "application-not-found")
	}

	return ApplicationFromGnmi(req.EnterpriseId, response)

}

func (a *ApplicationHandler) CreateApplication(ctx context.Context, app *v1.Application) (*v1.Application, error) {
	log.Debug("creating-application")
	gnmiCtx, cancel := context.WithTimeout(ctx, a.aether21.GnmiTimeout)
	defer cancel()

	gnmiPath := "/aether/v2.1.x/{enterprise-id}/application/{application-id}"
	gnmiApp, err := ApplicationToGnmi(ctx, app)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	gnmiUpdate, err := aether_2_1_0.EncodeToGnmiApplication(gnmiApp, false, false,
		types.EnterpriseId(app.EnterpriseId), "", app.ApplicationId)

	if err != nil {
		return nil, err
	}

	log.Infow("gnmiUpdate", "update", gnmiUpdate)

	gnmiSet, err := roc_utils.NewGnmiSetUpdateRequestUpdates(gnmiPath, app.EnterpriseId, gnmiUpdate, app.ApplicationId)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := a.aether21.GnmiClient.Set(gnmiCtx, gnmiSet)
	if err != nil {
		log.Errorw("gnmiSetError", "err", err)
		return nil, err
	}
	log.Infof("gnmiSetResponse", gnmiSetResponse.String())
	return app, nil

}

func (a *ApplicationHandler) DeleteApplication(ctx context.Context, req *v1.ApplicationFilter) (*v1.Empty, error) {
	log.Debug("delete-application")
	gnmiCtx, cancel := context.WithTimeout(ctx, a.aether21.GnmiTimeout)
	defer cancel()

	response, err := a.aether21.GnmiDeleteApplication(gnmiCtx, "/aether/v2.1.x/{enterprise-id}/application/{application-id}",
		types.EnterpriseId(req.EnterpriseId), req.ApplicationId)

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "application-not-found")
	}

	return &v1.Empty{}, nil
}

func ApplicationToGnmi(ctx context.Context, app *v1.Application) (*types.Application, error) {
	ep, err := endpoints.EndpointsToGnmi(ctx, app.Endpoints)
	if err != nil {
		return nil, err
	}
	gnmiApp := &types.Application{
		Address:       app.Address,
		ApplicationId: types.ListKey(app.ApplicationId),
		Description:   &app.Description,
		DisplayName:   &app.Name,
		Endpoint:      ep,
	}
	return gnmiApp, nil
}

func ApplicationFromGnmi(enterpriseId string, gnmiApp *types.Application) (*v1.Application, error) {
	eps, err := endpoints.FromGnmi(gnmiApp.Endpoint)
	if err != nil {
		return nil, err
	}
	return &v1.Application{
		ApplicationId: string(gnmiApp.ApplicationId),
		Name:          utils.PointerToString(gnmiApp.DisplayName),
		Description:   utils.PointerToString(gnmiApp.Description),
		Address:       gnmiApp.Address,
		Endpoints:     eps,
		EnterpriseId:  enterpriseId,
	}, nil
}

func ApplicationsFromGnmi(enterpriseId string, gnmiApps *types.ApplicationList) (*v1.Applications, error) {
	apps := v1.Applications{
		Applications: []*v1.Application{},
	}

	for _, a := range *gnmiApps {

		app, err := ApplicationFromGnmi(enterpriseId, &a)

		if err != nil {
			return nil, err
		}

		apps.Applications = append(apps.Applications, app)
	}
	return &apps, nil
}

func NewApplicationHandler(gnmi *onos_config.GnmiManager) *ApplicationHandler {
	return &ApplicationHandler{aether21: gnmi.Aether21}
}
