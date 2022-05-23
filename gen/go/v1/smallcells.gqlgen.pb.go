package v1

import (
	context "context"
)

type SmallCellServiceResolvers struct{ Service SmallCellServiceServer }

func (s *SmallCellServiceResolvers) SmallCellServiceGetSmallCells(ctx context.Context, in *GetSmallCellsRequest) (*GetSmallCellsResponse, error) {
	return s.Service.GetSmallCells(ctx, in)
}

type SmallCellInput = SmallCell
type GetSmallCellsResponseInput = GetSmallCellsResponse
type GetSmallCellsRequestInput = GetSmallCellsRequest
