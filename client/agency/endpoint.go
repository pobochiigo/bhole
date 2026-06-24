package agency

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizagency "com.gitlab/pobochiigo/bhole/internal/agency"
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
