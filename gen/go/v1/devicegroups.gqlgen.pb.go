package v1

import (
	context "context"
)

type DeviceGroupServiceResolvers struct{ Service DeviceGroupServiceServer }

func (s *DeviceGroupServiceResolvers) DeviceGroupServiceGetDeviceGroups(ctx context.Context, in *GetDeviceGroupsRequest) (*GetDeviceGroupsResponse, error) {
	return s.Service.GetDeviceGroups(ctx, in)
}

type DeviceGroupInput = DeviceGroup
type GetDeviceGroupsResponseInput = GetDeviceGroupsResponse
type GetDeviceGroupsRequestInput = GetDeviceGroupsRequest
