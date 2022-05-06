package v1

import (
	context "context"
)

type ApplicationServiceResolvers struct{ Service ApplicationServiceServer }

func (s *ApplicationServiceResolvers) ApplicationServiceGetApplications(ctx context.Context, in *GetApplicationsRequest) (*Applications, error) {
	return s.Service.GetApplications(ctx, in)
}
func (s *ApplicationServiceResolvers) ApplicationServiceGetApplication(ctx context.Context, in *ApplicationFilter) (*Application, error) {
	return s.Service.GetApplication(ctx, in)
}
func (s *ApplicationServiceResolvers) ApplicationServiceCreateOrUpdateApplication(ctx context.Context, in *Application) (*Application, error) {
	return s.Service.CreateOrUpdateApplication(ctx, in)
}
func (s *ApplicationServiceResolvers) ApplicationServiceDeleteApplication(ctx context.Context, in *ApplicationFilter) (*bool, error) {
	_, err := s.Service.DeleteApplication(ctx, in)
	return nil, err
}

type ApplicationInput = Application
type ApplicationsInput = Applications
type GetApplicationsRequestInput = GetApplicationsRequest
type ApplicationFilterInput = ApplicationFilter
