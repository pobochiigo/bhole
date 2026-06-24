package event

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListEvents endpoint.Endpoint[*ListEventsRequest, *ListEventsResponse]
	getEvent    endpoint.Endpoint[*GetEventRequest, *Event]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListEvents: makeListListEventsEndpoint(svc),
		getEvent:    makeGetEventEndpoint(svc),
	}
}

func makeListListEventsEndpoint(svc Service) endpoint.Endpoint[*ListEventsRequest, *ListEventsResponse] {
	return func(ctx context.Context, req *ListEventsRequest) (*ListEventsResponse, error) {
		return svc.ListEvents(ctx, req)
	}
}

func makeGetEventEndpoint(svc Service) endpoint.Endpoint[*GetEventRequest, *Event] {
	return func(ctx context.Context, req *GetEventRequest) (*Event, error) {
		return svc.GetEvent(ctx, req)
	}
}
