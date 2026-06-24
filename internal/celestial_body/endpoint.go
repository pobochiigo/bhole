package celestial_body

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListCelestialBodies endpoint.Endpoint[*ListCelestialBodiesRequest, *ListCelestialBodiesResponse]
	getCelestialBody    endpoint.Endpoint[*GetCelestialBodyRequest, *CelestialBody]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListCelestialBodies: makeListListCelestialBodiesEndpoint(svc),
		getCelestialBody:    makeGetCelestialBodyEndpoint(svc),
	}
}

func makeListListCelestialBodiesEndpoint(svc Service) endpoint.Endpoint[*ListCelestialBodiesRequest, *ListCelestialBodiesResponse] {
	return func(ctx context.Context, req *ListCelestialBodiesRequest) (*ListCelestialBodiesResponse, error) {
		return svc.ListCelestialBodies(ctx, req)
	}
}

func makeGetCelestialBodyEndpoint(svc Service) endpoint.Endpoint[*GetCelestialBodyRequest, *CelestialBody] {
	return func(ctx context.Context, req *GetCelestialBodyRequest) (*CelestialBody, error) {
		return svc.GetCelestialBody(ctx, req)
	}
}
