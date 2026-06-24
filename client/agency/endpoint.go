package agency

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
	bizagency "github.com/pobochiigo/bhole/internal/agency"
)

type endpoints struct {
	listListAgencies endpoint.Endpoint[*bizagency.ListAgenciesRequest, *bizagency.ListAgenciesResponse]
	getAgency    endpoint.Endpoint[*bizagency.GetAgencyRequest, *bizagency.Agency]
}

func (c *endpoints) ListAgencies(ctx context.Context, req *bizagency.ListAgenciesRequest) (*bizagency.ListAgenciesResponse, error) {
	return c.listListAgencies(ctx, req)
}

func (c *endpoints) GetAgency(ctx context.Context, req *bizagency.GetAgencyRequest) (*bizagency.Agency, error) {
	return c.getAgency(ctx, req)
}
