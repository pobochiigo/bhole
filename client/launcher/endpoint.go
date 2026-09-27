package launcher

import (
	"context"

	bizlauncher "github.com/pobochiigo/bhole/launcher"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListLaunchers endpoint.Endpoint[*bizlauncher.ListLaunchersRequest, *bizlauncher.ListLaunchersResponse]
	getLauncher       endpoint.Endpoint[*bizlauncher.GetLauncherRequest, *bizlauncher.Launcher]
}

func (c *endpoints) ListLaunchers(ctx context.Context, req *bizlauncher.ListLaunchersRequest) (*bizlauncher.ListLaunchersResponse, error) {
	return c.listListLaunchers(ctx, req)
}

func (c *endpoints) GetLauncher(ctx context.Context, req *bizlauncher.GetLauncherRequest) (*bizlauncher.Launcher, error) {
	return c.getLauncher(ctx, req)
}
