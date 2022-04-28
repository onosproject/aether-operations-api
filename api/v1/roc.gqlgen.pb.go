package v1

import (
	context "context"
)

type EnterpriseApiResolvers struct{ Service EnterpriseApiServer }

func (s *EnterpriseApiResolvers) EnterpriseApiGetEnterprises(ctx context.Context) (*Enterprises, error) {
	return s.Service.GetEnterprises(ctx, &Empty{})
}

type ApplicationApiResolvers struct{ Service ApplicationApiServer }

func (s *ApplicationApiResolvers) ApplicationApiGetApplications(ctx context.Context, in *EnterpriseId) (*Applications, error) {
	return s.Service.GetApplications(ctx, in)
}
func (s *ApplicationApiResolvers) ApplicationApiCreateApplication(ctx context.Context, in *Application) (*Application, error) {
	return s.Service.CreateApplication(ctx, in)
}

type EmptyInput = Empty
type EnterpriseIdInput = EnterpriseId
