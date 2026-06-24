package launch

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizlaunch "com.gitlab/pobochiigo/bhole/internal/launch"
)

type endpoints struct {
	listListLaunches endpoint.Endpoint[*bizlaunch.ListLaunchesRequest, *bizlaunch.ListLaunchesResponse]
	getLaunch    endpoint.Endpoint[*bizlaunch.GetLaunchRequest, *bizlaunch.Launch]
}

func (c *endpoints) ListLaunches(ctx context.Context, req *bizlaunch.ListLaunchesRequest) (*bizlaunch.ListLaunchesResponse, error) {
	return c.listListLaunches(ctx, req)
}

func (c *endpoints) GetLaunch(ctx context.Context, req *bizlaunch.GetLaunchRequest) (*bizlaunch.Launch, error) {
	return c.getLaunch(ctx, req)
}
