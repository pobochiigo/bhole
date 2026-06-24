package space_station

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizspace_station "com.gitlab/pobochiigo/bhole/internal/space_station"
)

type endpoints struct {
	listListSpaceStations endpoint.Endpoint[*bizspace_station.ListSpaceStationsRequest, *bizspace_station.ListSpaceStationsResponse]
	getSpaceStation    endpoint.Endpoint[*bizspace_station.GetSpaceStationRequest, *bizspace_station.SpaceStation]
}

func (c *endpoints) ListSpaceStations(ctx context.Context, req *bizspace_station.ListSpaceStationsRequest) (*bizspace_station.ListSpaceStationsResponse, error) {
	return c.listListSpaceStations(ctx, req)
}

func (c *endpoints) GetSpaceStation(ctx context.Context, req *bizspace_station.GetSpaceStationRequest) (*bizspace_station.SpaceStation, error) {
	return c.getSpaceStation(ctx, req)
}
