package v1

import (
	context "context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type EnterpriseServiceResolvers struct{ Service EnterpriseServiceServer }

func (s *EnterpriseServiceResolvers) EnterpriseServiceGetEnterprises(ctx context.Context) (*Enterprises, error) {
	return s.Service.GetEnterprises(ctx, &emptypb.Empty{})
}

type EnterpriseInput = Enterprise
type EnterprisesInput = Enterprises
