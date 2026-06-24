package pad

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListPads endpoint.Endpoint[*ListPadsRequest, *ListPadsResponse]
	getPad    endpoint.Endpoint[*GetPadRequest, *Pad]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListPads: makeListListPadsEndpoint(svc),
		getPad:    makeGetPadEndpoint(svc),
	}
}

func makeListListPadsEndpoint(svc Service) endpoint.Endpoint[*ListPadsRequest, *ListPadsResponse] {
	return func(ctx context.Context, req *ListPadsRequest) (*ListPadsResponse, error) {
		return svc.ListPads(ctx, req)
	}
}

func makeGetPadEndpoint(svc Service) endpoint.Endpoint[*GetPadRequest, *Pad] {
	return func(ctx context.Context, req *GetPadRequest) (*Pad, error) {
		return svc.GetPad(ctx, req)
	}
}
