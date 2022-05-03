/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package resolvers

import (
	"context"
	"fmt"
	appsv1 "github.com/onosproject/scaling-umbrella/gen/go/applications/v1"
	devicesv1 "github.com/onosproject/scaling-umbrella/gen/go/devices/v1"
	sitesv1 "github.com/onosproject/scaling-umbrella/gen/go/sites/v1"
	generated "github.com/onosproject/scaling-umbrella/gen/graph/graphql"
	"github.com/onosproject/scaling-umbrella/gen/graph/graphql/model"
	"github.com/onosproject/scaling-umbrella/internal/servers/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TODO: decouple independent resolvers

type Resolver struct {
	grpcServer *grpc.RocApiGrpcServer
	//enterprises TODO
	//applications TODO
}

func NewResolver(s *grpc.RocApiGrpcServer) generated.Config {
	r := Resolver{
		grpcServer: s,
		//enterprises: TODO
		//applications: TODO
	}
	return generated.Config{
		Resolvers: &r,
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &MutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &QueryResolver{r} }

type MutationResolver struct{ *Resolver }
type QueryResolver struct{ *Resolver }

func (r *MutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *QueryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *QueryResolver) Enterprises(ctx context.Context) ([]*model.Enterprise, error) {
	var enterprises []*model.Enterprise

	grpcEnterprises, _ := r.grpcServer.Services.EnterpriseService.GetEnterprises(ctx, &emptypb.Empty{})

	for _, e := range grpcEnterprises.Enterprises {

		enterprise := model.Enterprise{
			ID:          e.EnterpriseId,
			Name:        e.Name,
			Description: e.Description,
			//Sites:        sites,
			//Applications: apps,
		}
		enterprises = append(enterprises, &enterprise)
	}

	return enterprises, nil
}

func (r *QueryResolver) Applications(ctx context.Context, enterpriseID string) ([]*model.Application, error) {
	var apps []*model.Application

	protoEntId := appsv1.EnterpriseId{
		EnterpriseId: enterpriseID,
	}
	grpcApps, _ := r.grpcServer.Services.ApplicationService.GetApplications(ctx, &protoEntId)
	for _, a := range grpcApps.Applications {
		app := model.Application{
			ID:   a.ApplicationId,
			Name: a.Name,
		}
		apps = append(apps, &app)
	}

	return apps, nil
}

func (r *QueryResolver) Endpoints(ctx context.Context) ([]*model.Endpoint, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *QueryResolver) Sites(ctx context.Context, enterpriseID string) ([]*model.Site, error) {
	var sites []*model.Site

	// -----
	var devices []*model.Device
	// -----

	protoReq := sitesv1.GetSitesRequest{
		EnterpriseId: enterpriseID,
	}
	grpcSites, _ := r.grpcServer.Services.SiteService.GetSites(ctx, &protoReq)
	for _, a := range grpcSites.Sites {

		// -----
		devProto := devicesv1.GetDevicesRequest{
			EnterpriseId: enterpriseID,
			SiteId:       a.SiteId,
		}
		grpcDevices, _ := r.grpcServer.Services.DeviceService.GetDevices(ctx, &devProto)
		for _, d := range grpcDevices.Devices {
			d := model.Device{
				ID:          d.DeviceId,
				Name:        d.Name,
				Description: d.Description,
				SimCard:     nil,
				IP:          nil,
			}
			devices = append(devices, &d)
		}
		// ----

		site := model.Site{
			ID:      a.SiteId,
			Name:    a.Name,
			Devices: devices,
		}
		sites = append(sites, &site)
	}

	return sites, nil
}

func (r *QueryResolver) SmallCells(ctx context.Context) ([]*model.SmallCell, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *QueryResolver) Slices(ctx context.Context) ([]*model.Slice, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *QueryResolver) DeviceGroups(ctx context.Context) ([]*model.DeviceGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *QueryResolver) Devices(ctx context.Context, enterpriseID string, siteID string) ([]*model.Device, error) {
	var devices []*model.Device

	protoReq := devicesv1.GetDevicesRequest{
		EnterpriseId: enterpriseID,
		SiteId:       siteID,
	}
	grpcDevices, _ := r.grpcServer.Services.DeviceService.GetDevices(ctx, &protoReq)
	for _, a := range grpcDevices.Devices {
		d := model.Device{
			ID:          a.DeviceId,
			Name:        a.Name,
			Description: a.Description,
			SimCard:     nil,
			IP:          nil,
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *QueryResolver) SimCards(ctx context.Context) ([]*model.SimCard, error) {
	panic(fmt.Errorf("not implemented"))
}
