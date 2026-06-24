package spacewalk

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListSpacewalks endpoint.Endpoint[*ListSpacewalksRequest, *ListSpacewalksResponse]
	getSpacewalk    endpoint.Endpoint[*GetSpacewalkRequest, *Spacewalk]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListSpacewalks: makeListListSpacewalksEndpoint(svc),
		getSpacewalk:    makeGetSpacewalkEndpoint(svc),
	}
}

func makeListListSpacewalksEndpoint(svc Service) endpoint.Endpoint[*ListSpacewalksRequest, *ListSpacewalksResponse] {
	return func(ctx context.Context, req *ListSpacewalksRequest) (*ListSpacewalksResponse, error) {
		return svc.ListSpacewalks(ctx, req)
	}
}

func makeGetSpacewalkEndpoint(svc Service) endpoint.Endpoint[*GetSpacewalkRequest, *Spacewalk] {
	return func(ctx context.Context, req *GetSpacewalkRequest) (*Spacewalk, error) {
		return svc.GetSpacewalk(ctx, req)
	}
}
