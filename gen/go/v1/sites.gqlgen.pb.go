package v1

import (
	context "context"
)

type SiteServiceResolvers struct{ Service SiteServiceServer }

func (s *SiteServiceResolvers) SiteServiceGetSites(ctx context.Context, in *GetSitesRequest) (*GetSitesResponse, error) {
	return s.Service.GetSites(ctx, in)
}

type SiteInput = Site
type GetSitesResponseInput = GetSitesResponse
type GetSitesRequestInput = GetSitesRequest
