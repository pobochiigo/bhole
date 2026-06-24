package launcher

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListLaunchers endpoint.Endpoint[*ListLaunchersRequest, *ListLaunchersResponse]
	getLauncher    endpoint.Endpoint[*GetLauncherRequest, *Launcher]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListLaunchers: makeListListLaunchersEndpoint(svc),
		getLauncher:    makeGetLauncherEndpoint(svc),
	}
}

func makeListListLaunchersEndpoint(svc Service) endpoint.Endpoint[*ListLaunchersRequest, *ListLaunchersResponse] {
	return func(ctx context.Context, req *ListLaunchersRequest) (*ListLaunchersResponse, error) {
		return svc.ListLaunchers(ctx, req)
	}
}

func makeGetLauncherEndpoint(svc Service) endpoint.Endpoint[*GetLauncherRequest, *Launcher] {
	return func(ctx context.Context, req *GetLauncherRequest) (*Launcher, error) {
		return svc.GetLauncher(ctx, req)
	}
}
