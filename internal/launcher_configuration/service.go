package launcher_configuration

import "context"

type Service interface {
	ListLauncherConfigurations(ctx context.Context, req *ListLauncherConfigurationsRequest) (*ListLauncherConfigurationsResponse, error)
	GetLauncherConfiguration(ctx context.Context, req *GetLauncherConfigurationRequest) (*LauncherConfiguration, error)
}
