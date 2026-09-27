package space_station

import "context"

type Service interface {
	ListSpaceStations(ctx context.Context, req *ListSpaceStationsRequest) (*ListSpaceStationsResponse, error)
	GetSpaceStation(ctx context.Context, req *GetSpaceStationRequest) (*SpaceStation, error)
}
