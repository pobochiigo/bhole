package launcher

import "context"

type Service interface {
	ListLaunchers(ctx context.Context, req *ListLaunchersRequest) (*ListLaunchersResponse, error)
	GetLauncher(ctx context.Context, req *GetLauncherRequest) (*Launcher, error)
}
