package v1

import (
	context "context"
)

type DeviceServiceResolvers struct{ Service DeviceServiceServer }

func (s *DeviceServiceResolvers) DeviceServiceGetDevices(ctx context.Context, in *GetDevicesRequest) (*GetDevicesResponse, error) {
	return s.Service.GetDevices(ctx, in)
}

type DeviceInput = Device
type GetDevicesResponseInput = GetDevicesResponse
type GetDevicesRequestInput = GetDevicesRequest
