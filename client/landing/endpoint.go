package landing

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
	bizlanding "github.com/pobochiigo/bhole/internal/landing"
)

type endpoints struct {
	listListLandings endpoint.Endpoint[*bizlanding.ListLandingsRequest, *bizlanding.ListLandingsResponse]
	getLanding    endpoint.Endpoint[*bizlanding.GetLandingRequest, *bizlanding.Landing]
}

func (c *endpoints) ListLandings(ctx context.Context, req *bizlanding.ListLandingsRequest) (*bizlanding.ListLandingsResponse, error) {
	return c.listListLandings(ctx, req)
}

func (c *endpoints) GetLanding(ctx context.Context, req *bizlanding.GetLandingRequest) (*bizlanding.Landing, error) {
	return c.getLanding(ctx, req)
}
