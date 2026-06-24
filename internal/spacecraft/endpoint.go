package spacecraft

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListSpacecrafts endpoint.Endpoint[*ListSpacecraftsRequest, *ListSpacecraftsResponse]
	getSpacecraft    endpoint.Endpoint[*GetSpacecraftRequest, *Spacecraft]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListSpacecrafts: makeListListSpacecraftsEndpoint(svc),
		getSpacecraft:    makeGetSpacecraftEndpoint(svc),
	}
}

func makeListListSpacecraftsEndpoint(svc Service) endpoint.Endpoint[*ListSpacecraftsRequest, *ListSpacecraftsResponse] {
	return func(ctx context.Context, req *ListSpacecraftsRequest) (*ListSpacecraftsResponse, error) {
		return svc.ListSpacecrafts(ctx, req)
	}
}

func makeGetSpacecraftEndpoint(svc Service) endpoint.Endpoint[*GetSpacecraftRequest, *Spacecraft] {
	return func(ctx context.Context, req *GetSpacecraftRequest) (*Spacecraft, error) {
		return svc.GetSpacecraft(ctx, req)
	}
}
