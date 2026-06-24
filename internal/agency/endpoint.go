package agency

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListAgencies endpoint.Endpoint[*ListAgenciesRequest, *ListAgenciesResponse]
	getAgency    endpoint.Endpoint[*GetAgencyRequest, *Agency]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListAgencies: makeListListAgenciesEndpoint(svc),
		getAgency:    makeGetAgencyEndpoint(svc),
	}
}

func makeListListAgenciesEndpoint(svc Service) endpoint.Endpoint[*ListAgenciesRequest, *ListAgenciesResponse] {
	return func(ctx context.Context, req *ListAgenciesRequest) (*ListAgenciesResponse, error) {
		return svc.ListAgencies(ctx, req)
	}
}

func makeGetAgencyEndpoint(svc Service) endpoint.Endpoint[*GetAgencyRequest, *Agency] {
	return func(ctx context.Context, req *GetAgencyRequest) (*Agency, error) {
		return svc.GetAgency(ctx, req)
	}
}
