/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package graphqlapp

import (
	"context"
	"github.com/onosproject/scaling-umbrella/gen/go/v1"
	gqlgen "github.com/onosproject/scaling-umbrella/gen/graph"
)

type App struct {
	EnterpriseResolver  *v1.EnterpriseServiceResolvers
	ApplicationResolver *v1.ApplicationServiceResolvers
	SiteResolver        *v1.SiteServiceResolvers
	DeviceResolver      *v1.DeviceServiceResolvers
	DeviceGroupResolver *v1.DeviceGroupServiceResolvers
	SimCardResolver     *v1.SimCardServiceResolvers
	SliceResolver       *v1.SliceServiceResolvers
	SmallCellResolver   *v1.SmallCellServiceResolvers
}

func NewApp(
	enterprise *v1.EnterpriseServiceResolvers,
	application *v1.ApplicationServiceResolvers,
	site *v1.SiteServiceResolvers,
	device *v1.DeviceServiceResolvers,
	deviceGroup *v1.DeviceGroupServiceResolvers,
	simCard *v1.SimCardServiceResolvers,
	slice *v1.SliceServiceResolvers,
	smallCell *v1.SmallCellServiceResolvers) *App {
	return &App{
		EnterpriseResolver:  enterprise,
		ApplicationResolver: application,
		SiteResolver:        site,
		DeviceResolver:      device,
		DeviceGroupResolver: deviceGroup,
		SimCardResolver:     simCard,
		SliceResolver:       slice,
		SmallCellResolver:   smallCell,
	}
}

// Mutation returns MutationResolver implementation.
func (r *App) Mutation() gqlgen.MutationResolver {
	return &MutationResolver{r}
}

// Query returns QueryResolver implementation.
func (r *App) Query() gqlgen.QueryResolver {
	return &QueryResolver{r}
}

type MutationResolver struct {
	*App
}

type QueryResolver struct {
	*App
}

func (r *MutationResolver) CreateApplication(ctx context.Context, in *v1.Application) (*v1.Application, error) {
	panic("implement me")
}

// TODO: look into onos-topo as we are not returning enterprises only connectivity-service
func (r *QueryResolver) Enterprises(ctx context.Context) (*v1.GetEnterprisesResponse, error) {
	resp, err := r.EnterpriseResolver.EnterpriseServiceGetEnterprises(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *QueryResolver) Applications(ctx context.Context, enterpriseID string) (*v1.GetApplicationsResponse, error) {
	req := v1.GetApplicationsRequest{
		EnterpriseId: enterpriseID,
	}
	resp, err := r.ApplicationResolver.ApplicationServiceGetApplications(ctx, &req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *QueryResolver) Sites(ctx context.Context, enterpriseID string) (*v1.GetSitesResponse, error) {
	req := v1.GetSitesRequest{
		EnterpriseId: enterpriseID,
	}
	resp, err := r.SiteResolver.SiteServiceGetSites(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *QueryResolver) Devices(ctx context.Context, enterpriseID string, siteID string) (*v1.GetDevicesResponse, error) {
	req := v1.GetDevicesRequest{
		EnterpriseId: enterpriseID,
		SiteId:       siteID,
	}
	resp, err := r.DeviceResolver.DeviceServiceGetDevices(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// TODO: fix
func (r *QueryResolver) DevicesGroups(ctx context.Context, enterpriseID string, siteID string) (*v1.GetDeviceGroupsResponse, error) {
	req := v1.GetDeviceGroupsRequest{
		EnterpriseId: enterpriseID,
		SiteId:       siteID,
	}
	resp, err := r.DeviceGroupResolver.DeviceGroupServiceGetDeviceGroups(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *QueryResolver) SimCards(ctx context.Context, enterpriseID string, siteID string) (*v1.GetSimCardsResponse, error) {
	req := v1.GetSimCardsRequest{
		EnterpriseId: enterpriseID,
		SiteId:       siteID,
	}
	resp, err := r.SimCardResolver.SimCardServiceGetSimCards(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *QueryResolver) Slices(ctx context.Context, enterpriseID string, siteID string) (*v1.GetSlicesResponse, error) {
	req := v1.GetSlicesRequest{
		EnterpriseId: enterpriseID,
		SiteId:       siteID,
	}
	resp, err := r.SliceResolver.SliceServiceGetSlices(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// TODO: fix
func (r *QueryResolver) SmallCells(ctx context.Context, enterpriseID string, siteID string) (*v1.GetSmallCellsResponse, error) {
	req := v1.GetSmallCellsRequest{
		EnterpriseId: enterpriseID,
		SiteId:       siteID,
	}
	resp, err := r.SmallCellResolver.SmallCellServiceGetSmallCells(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
