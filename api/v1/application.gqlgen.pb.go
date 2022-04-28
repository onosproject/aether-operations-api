package v1

import (
	context "context"
)

type ApplicationServiceResolvers struct{ Service ApplicationServiceServer }

func (s *ApplicationServiceResolvers) ApplicationServiceGetApplications(ctx context.Context, in *EnterpriseId) (*Applications, error) {
	return s.Service.GetApplications(ctx, in)
}
func (s *ApplicationServiceResolvers) ApplicationServiceCreateApplication(ctx context.Context, in *Application) (*Application, error) {
	return s.Service.CreateApplication(ctx, in)
}

type ApplicationInput = Application
type ApplicationsInput = Applications
type EnterpriseIdInput = EnterpriseId
