package landing

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListLandings endpoint.Endpoint[*ListLandingsRequest, *ListLandingsResponse]
	getLanding    endpoint.Endpoint[*GetLandingRequest, *Landing]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListLandings: makeListListLandingsEndpoint(svc),
		getLanding:    makeGetLandingEndpoint(svc),
	}
}

func makeListListLandingsEndpoint(svc Service) endpoint.Endpoint[*ListLandingsRequest, *ListLandingsResponse] {
	return func(ctx context.Context, req *ListLandingsRequest) (*ListLandingsResponse, error) {
		return svc.ListLandings(ctx, req)
	}
}

func makeGetLandingEndpoint(svc Service) endpoint.Endpoint[*GetLandingRequest, *Landing] {
	return func(ctx context.Context, req *GetLandingRequest) (*Landing, error) {
		return svc.GetLanding(ctx, req)
	}
}
