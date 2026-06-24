package launch

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListLaunches endpoint.Endpoint[*ListLaunchesRequest, *ListLaunchesResponse]
	getLaunch    endpoint.Endpoint[*GetLaunchRequest, *Launch]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListLaunches: makeListListLaunchesEndpoint(svc),
		getLaunch:    makeGetLaunchEndpoint(svc),
	}
}

func makeListListLaunchesEndpoint(svc Service) endpoint.Endpoint[*ListLaunchesRequest, *ListLaunchesResponse] {
	return func(ctx context.Context, req *ListLaunchesRequest) (*ListLaunchesResponse, error) {
		return svc.ListLaunches(ctx, req)
	}
}

func makeGetLaunchEndpoint(svc Service) endpoint.Endpoint[*GetLaunchRequest, *Launch] {
	return func(ctx context.Context, req *GetLaunchRequest) (*Launch, error) {
		return svc.GetLaunch(ctx, req)
	}
}
