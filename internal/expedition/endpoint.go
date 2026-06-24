package expedition

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListExpeditions endpoint.Endpoint[*ListExpeditionsRequest, *ListExpeditionsResponse]
	getExpedition    endpoint.Endpoint[*GetExpeditionRequest, *Expedition]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListExpeditions: makeListListExpeditionsEndpoint(svc),
		getExpedition:    makeGetExpeditionEndpoint(svc),
	}
}

func makeListListExpeditionsEndpoint(svc Service) endpoint.Endpoint[*ListExpeditionsRequest, *ListExpeditionsResponse] {
	return func(ctx context.Context, req *ListExpeditionsRequest) (*ListExpeditionsResponse, error) {
		return svc.ListExpeditions(ctx, req)
	}
}

func makeGetExpeditionEndpoint(svc Service) endpoint.Endpoint[*GetExpeditionRequest, *Expedition] {
	return func(ctx context.Context, req *GetExpeditionRequest) (*Expedition, error) {
		return svc.GetExpedition(ctx, req)
	}
}
