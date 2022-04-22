package application

import (
	"context"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	v1 "github.com/onosproject/roc-api/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

type ApplicationHandler struct {
	gnmiClient southbound.GnmiClient
	aether21   *aether_2_1_0.ServerImpl
}

func (a ApplicationHandler) ListApplications() (*v1.Applications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), a.aether21.GnmiTimeout)
	defer cancel()

	const enterpriseId = "defaultent"
	response, err := a.aether21.GnmiGetApplicationList(ctx, "/aether/v2.1.x/{enterprise-id}/application", enterpriseId)

	if err != nil {
		return nil, err
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return nil, status.Error(codes.NotFound, "applications-not-found")
	}
	apps := v1.Applications{
		Applications: []*v1.Application{},
	}

	for _, a := range *response {

		eps := []*v1.Endpoint{}

		for _, ep := range *a.Endpoint {
			eps = append(eps, &v1.Endpoint{
				ID:          string(ep.EndpointId),
				Description: *ep.Description,
				DisplayName: *ep.DisplayName,
				Mbr: &v1.MBR{
					Uplink:   *ep.Mbr.Uplink,
					Downlink: *ep.Mbr.Downlink,
				},
				PortStart: int32(*ep.PortStart),
				PortEnd:   int32(*ep.PortEnd),
				Protocol:  *ep.Protocol,
			})
		}

		apps.Applications = append(apps.Applications, &v1.Application{
			ID:          string(a.ApplicationId),
			Description: *a.Description,
			Endpoint:    eps})
	}
	return &apps, nil
}

func NewApplicationHandler(client southbound.GnmiClient, aether21 *aether_2_1_0.ServerImpl) *ApplicationHandler {
	return &ApplicationHandler{gnmiClient: client, aether21: aether21}
}
