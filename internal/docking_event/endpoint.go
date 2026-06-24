package docking_event

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListDockingEvents endpoint.Endpoint[*ListDockingEventsRequest, *ListDockingEventsResponse]
	getDockingEvent    endpoint.Endpoint[*GetDockingEventRequest, *DockingEvent]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListDockingEvents: makeListListDockingEventsEndpoint(svc),
		getDockingEvent:    makeGetDockingEventEndpoint(svc),
	}
}

func makeListListDockingEventsEndpoint(svc Service) endpoint.Endpoint[*ListDockingEventsRequest, *ListDockingEventsResponse] {
	return func(ctx context.Context, req *ListDockingEventsRequest) (*ListDockingEventsResponse, error) {
		return svc.ListDockingEvents(ctx, req)
	}
}

func makeGetDockingEventEndpoint(svc Service) endpoint.Endpoint[*GetDockingEventRequest, *DockingEvent] {
	return func(ctx context.Context, req *GetDockingEventRequest) (*DockingEvent, error) {
		return svc.GetDockingEvent(ctx, req)
	}
}
