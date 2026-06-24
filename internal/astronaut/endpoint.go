package astronaut

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListAstronauts endpoint.Endpoint[*ListAstronautsRequest, *ListAstronautsResponse]
	getAstronaut    endpoint.Endpoint[*GetAstronautRequest, *Astronaut]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListAstronauts: makeListListAstronautsEndpoint(svc),
		getAstronaut:    makeGetAstronautEndpoint(svc),
	}
}

func makeListListAstronautsEndpoint(svc Service) endpoint.Endpoint[*ListAstronautsRequest, *ListAstronautsResponse] {
	return func(ctx context.Context, req *ListAstronautsRequest) (*ListAstronautsResponse, error) {
		return svc.ListAstronauts(ctx, req)
	}
}

func makeGetAstronautEndpoint(svc Service) endpoint.Endpoint[*GetAstronautRequest, *Astronaut] {
	return func(ctx context.Context, req *GetAstronautRequest) (*Astronaut, error) {
		return svc.GetAstronaut(ctx, req)
	}
}
