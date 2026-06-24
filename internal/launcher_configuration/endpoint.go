package launcher_configuration

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListLauncherConfigurations endpoint.Endpoint[*ListLauncherConfigurationsRequest, *ListLauncherConfigurationsResponse]
	getLauncherConfiguration    endpoint.Endpoint[*GetLauncherConfigurationRequest, *LauncherConfiguration]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListLauncherConfigurations: makeListListLauncherConfigurationsEndpoint(svc),
		getLauncherConfiguration:    makeGetLauncherConfigurationEndpoint(svc),
	}
}

func makeListListLauncherConfigurationsEndpoint(svc Service) endpoint.Endpoint[*ListLauncherConfigurationsRequest, *ListLauncherConfigurationsResponse] {
	return func(ctx context.Context, req *ListLauncherConfigurationsRequest) (*ListLauncherConfigurationsResponse, error) {
		return svc.ListLauncherConfigurations(ctx, req)
	}
}

func makeGetLauncherConfigurationEndpoint(svc Service) endpoint.Endpoint[*GetLauncherConfigurationRequest, *LauncherConfiguration] {
	return func(ctx context.Context, req *GetLauncherConfigurationRequest) (*LauncherConfiguration, error) {
		return svc.GetLauncherConfiguration(ctx, req)
	}
}
