package agency

import (
	"context"

	bizagency "github.com/pobochiigo/bhole/pkg/agency"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListAgencies endpoint.Endpoint[*bizagency.ListAgenciesRequest, *bizagency.ListAgenciesResponse]
	getAgency        endpoint.Endpoint[*bizagency.GetAgencyRequest, *bizagency.Agency]
}

func (c *endpoints) ListAgencies(ctx context.Context, req *bizagency.ListAgenciesRequest) (*bizagency.ListAgenciesResponse, error) {
	return c.listListAgencies(ctx, req)
}

func (c *endpoints) GetAgency(ctx context.Context, req *bizagency.GetAgencyRequest) (*bizagency.Agency, error) {
	return c.getAgency(ctx, req)
}
