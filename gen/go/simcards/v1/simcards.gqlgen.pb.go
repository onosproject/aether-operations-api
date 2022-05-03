package v1

import (
	context "context"
)

type SimCardServiceResolvers struct{ Service SimCardServiceServer }

func (s *SimCardServiceResolvers) SimCardServiceGetSimCards(ctx context.Context, in *GetSimCardsRequest) (*GetSimCardsResponse, error) {
	return s.Service.GetSimCards(ctx, in)
}

type SimCardInput = SimCard
type GetSimCardsResponseInput = GetSimCardsResponse
type GetSimCardsRequestInput = GetSimCardsRequest
