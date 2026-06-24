package update

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListUpdates endpoint.Endpoint[*ListUpdatesRequest, *ListUpdatesResponse]
	getUpdate    endpoint.Endpoint[*GetUpdateRequest, *Update]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListUpdates: makeListListUpdatesEndpoint(svc),
		getUpdate:    makeGetUpdateEndpoint(svc),
	}
}

func makeListListUpdatesEndpoint(svc Service) endpoint.Endpoint[*ListUpdatesRequest, *ListUpdatesResponse] {
	return func(ctx context.Context, req *ListUpdatesRequest) (*ListUpdatesResponse, error) {
		return svc.ListUpdates(ctx, req)
	}
}

func makeGetUpdateEndpoint(svc Service) endpoint.Endpoint[*GetUpdateRequest, *Update] {
	return func(ctx context.Context, req *GetUpdateRequest) (*Update, error) {
		return svc.GetUpdate(ctx, req)
	}
}
