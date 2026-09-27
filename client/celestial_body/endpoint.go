package celestial_body

import (
	"context"

	bizcelestial_body "github.com/pobochiigo/bhole/pkg/celestial_body"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListCelestialBodies endpoint.Endpoint[*bizcelestial_body.ListCelestialBodiesRequest, *bizcelestial_body.ListCelestialBodiesResponse]
	getCelestialBody        endpoint.Endpoint[*bizcelestial_body.GetCelestialBodyRequest, *bizcelestial_body.CelestialBody]
}

func (c *endpoints) ListCelestialBodies(ctx context.Context, req *bizcelestial_body.ListCelestialBodiesRequest) (*bizcelestial_body.ListCelestialBodiesResponse, error) {
	return c.listListCelestialBodies(ctx, req)
}

func (c *endpoints) GetCelestialBody(ctx context.Context, req *bizcelestial_body.GetCelestialBodyRequest) (*bizcelestial_body.CelestialBody, error) {
	return c.getCelestialBody(ctx, req)
}
