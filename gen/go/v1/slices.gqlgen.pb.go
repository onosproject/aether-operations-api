package v1

import (
	context "context"
)

type SliceServiceResolvers struct{ Service SliceServiceServer }

func (s *SliceServiceResolvers) SliceServiceGetSlices(ctx context.Context, in *GetSlicesRequest) (*GetSlicesResponse, error) {
	return s.Service.GetSlices(ctx, in)
}

type SliceInput = Slice
type GetSlicesResponseInput = GetSlicesResponse
type GetSlicesRequestInput = GetSlicesRequest
