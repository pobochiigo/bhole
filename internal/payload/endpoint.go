package payload

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListPayloads endpoint.Endpoint[*ListPayloadsRequest, *ListPayloadsResponse]
	getPayload    endpoint.Endpoint[*GetPayloadRequest, *Payload]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListPayloads: makeListListPayloadsEndpoint(svc),
		getPayload:    makeGetPayloadEndpoint(svc),
	}
}

func makeListListPayloadsEndpoint(svc Service) endpoint.Endpoint[*ListPayloadsRequest, *ListPayloadsResponse] {
	return func(ctx context.Context, req *ListPayloadsRequest) (*ListPayloadsResponse, error) {
		return svc.ListPayloads(ctx, req)
	}
}

func makeGetPayloadEndpoint(svc Service) endpoint.Endpoint[*GetPayloadRequest, *Payload] {
	return func(ctx context.Context, req *GetPayloadRequest) (*Payload, error) {
		return svc.GetPayload(ctx, req)
	}
}
