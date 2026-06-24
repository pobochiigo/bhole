package space_station

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListSpaceStations endpoint.Endpoint[*ListSpaceStationsRequest, *ListSpaceStationsResponse]
	getSpaceStation    endpoint.Endpoint[*GetSpaceStationRequest, *SpaceStation]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListSpaceStations: makeListListSpaceStationsEndpoint(svc),
		getSpaceStation:    makeGetSpaceStationEndpoint(svc),
	}
}

func makeListListSpaceStationsEndpoint(svc Service) endpoint.Endpoint[*ListSpaceStationsRequest, *ListSpaceStationsResponse] {
	return func(ctx context.Context, req *ListSpaceStationsRequest) (*ListSpaceStationsResponse, error) {
		return svc.ListSpaceStations(ctx, req)
	}
}

func makeGetSpaceStationEndpoint(svc Service) endpoint.Endpoint[*GetSpaceStationRequest, *SpaceStation] {
	return func(ctx context.Context, req *GetSpaceStationRequest) (*SpaceStation, error) {
		return svc.GetSpaceStation(ctx, req)
	}
}
