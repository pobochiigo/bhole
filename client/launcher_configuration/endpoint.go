package launcher_configuration

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizlauncher_configuration "com.gitlab/pobochiigo/bhole/internal/launcher_configuration"
)

type endpoints struct {
	listListLauncherConfigurations endpoint.Endpoint[*bizlauncher_configuration.ListLauncherConfigurationsRequest, *bizlauncher_configuration.ListLauncherConfigurationsResponse]
	getLauncherConfiguration    endpoint.Endpoint[*bizlauncher_configuration.GetLauncherConfigurationRequest, *bizlauncher_configuration.LauncherConfiguration]
}

func (c *endpoints) ListLauncherConfigurations(ctx context.Context, req *bizlauncher_configuration.ListLauncherConfigurationsRequest) (*bizlauncher_configuration.ListLauncherConfigurationsResponse, error) {
	return c.listListLauncherConfigurations(ctx, req)
}

func (c *endpoints) GetLauncherConfiguration(ctx context.Context, req *bizlauncher_configuration.GetLauncherConfigurationRequest) (*bizlauncher_configuration.LauncherConfiguration, error) {
	return c.getLauncherConfiguration(ctx, req)
}
