package location

import (
	"context"

	bizlocation "github.com/pobochiigo/bhole/location"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListLocations endpoint.Endpoint[*bizlocation.ListLocationsRequest, *bizlocation.ListLocationsResponse]
	getLocation       endpoint.Endpoint[*bizlocation.GetLocationRequest, *bizlocation.Location]
}

func (c *endpoints) ListLocations(ctx context.Context, req *bizlocation.ListLocationsRequest) (*bizlocation.ListLocationsResponse, error) {
	return c.listListLocations(ctx, req)
}

func (c *endpoints) GetLocation(ctx context.Context, req *bizlocation.GetLocationRequest) (*bizlocation.Location, error) {
	return c.getLocation(ctx, req)
}
