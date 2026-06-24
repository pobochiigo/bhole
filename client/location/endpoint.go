package location

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizlocation "com.gitlab/pobochiigo/bhole/internal/location"
)

type endpoints struct {
	listListLocations endpoint.Endpoint[*bizlocation.ListLocationsRequest, *bizlocation.ListLocationsResponse]
	getLocation    endpoint.Endpoint[*bizlocation.GetLocationRequest, *bizlocation.Location]
}

func (c *endpoints) ListLocations(ctx context.Context, req *bizlocation.ListLocationsRequest) (*bizlocation.ListLocationsResponse, error) {
	return c.listListLocations(ctx, req)
}

func (c *endpoints) GetLocation(ctx context.Context, req *bizlocation.GetLocationRequest) (*bizlocation.Location, error) {
	return c.getLocation(ctx, req)
}
