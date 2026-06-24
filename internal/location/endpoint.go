package location

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListLocations endpoint.Endpoint[*ListLocationsRequest, *ListLocationsResponse]
	getLocation    endpoint.Endpoint[*GetLocationRequest, *Location]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListLocations: makeListListLocationsEndpoint(svc),
		getLocation:    makeGetLocationEndpoint(svc),
	}
}

func makeListListLocationsEndpoint(svc Service) endpoint.Endpoint[*ListLocationsRequest, *ListLocationsResponse] {
	return func(ctx context.Context, req *ListLocationsRequest) (*ListLocationsResponse, error) {
		return svc.ListLocations(ctx, req)
	}
}

func makeGetLocationEndpoint(svc Service) endpoint.Endpoint[*GetLocationRequest, *Location] {
	return func(ctx context.Context, req *GetLocationRequest) (*Location, error) {
		return svc.GetLocation(ctx, req)
	}
}
