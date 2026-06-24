package celestial_body

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizcelestial_body "com.gitlab/pobochiigo/bhole/internal/celestial_body"
)

type endpoints struct {
	listListCelestialBodies endpoint.Endpoint[*bizcelestial_body.ListCelestialBodiesRequest, *bizcelestial_body.ListCelestialBodiesResponse]
	getCelestialBody    endpoint.Endpoint[*bizcelestial_body.GetCelestialBodyRequest, *bizcelestial_body.CelestialBody]
}

func (c *endpoints) ListCelestialBodies(ctx context.Context, req *bizcelestial_body.ListCelestialBodiesRequest) (*bizcelestial_body.ListCelestialBodiesResponse, error) {
	return c.listListCelestialBodies(ctx, req)
}

func (c *endpoints) GetCelestialBody(ctx context.Context, req *bizcelestial_body.GetCelestialBodyRequest) (*bizcelestial_body.CelestialBody, error) {
	return c.getCelestialBody(ctx, req)
}
