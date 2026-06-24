package location

import "context"

type Service interface {
	ListLocations(ctx context.Context, req *ListLocationsRequest) (*ListLocationsResponse, error)
	GetLocation(ctx context.Context, req *GetLocationRequest) (*Location, error)
}
