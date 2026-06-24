package astronaut

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizastronaut "com.gitlab/pobochiigo/bhole/internal/astronaut"
)

type endpoints struct {
	listListAstronauts endpoint.Endpoint[*bizastronaut.ListAstronautsRequest, *bizastronaut.ListAstronautsResponse]
	getAstronaut    endpoint.Endpoint[*bizastronaut.GetAstronautRequest, *bizastronaut.Astronaut]
}

func (c *endpoints) ListAstronauts(ctx context.Context, req *bizastronaut.ListAstronautsRequest) (*bizastronaut.ListAstronautsResponse, error) {
	return c.listListAstronauts(ctx, req)
}

func (c *endpoints) GetAstronaut(ctx context.Context, req *bizastronaut.GetAstronautRequest) (*bizastronaut.Astronaut, error) {
	return c.getAstronaut(ctx, req)
}
