package landing

import (
	"context"

	bizlanding "github.com/pobochiigo/bhole/landing"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListLandings endpoint.Endpoint[*bizlanding.ListLandingsRequest, *bizlanding.ListLandingsResponse]
	getLanding       endpoint.Endpoint[*bizlanding.GetLandingRequest, *bizlanding.Landing]
}

func (c *endpoints) ListLandings(ctx context.Context, req *bizlanding.ListLandingsRequest) (*bizlanding.ListLandingsResponse, error) {
	return c.listListLandings(ctx, req)
}

func (c *endpoints) GetLanding(ctx context.Context, req *bizlanding.GetLandingRequest) (*bizlanding.Landing, error) {
	return c.getLanding(ctx, req)
}
