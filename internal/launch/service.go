package launch

import "context"

type Service interface {
	ListLaunches(ctx context.Context, req *ListLaunchesRequest) (*ListLaunchesResponse, error)
	GetLaunch(ctx context.Context, req *GetLaunchRequest) (*Launch, error)
}
