package docking_event

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
	bizdocking_event "github.com/pobochiigo/bhole/internal/docking_event"
)

type endpoints struct {
	listListDockingEvents endpoint.Endpoint[*bizdocking_event.ListDockingEventsRequest, *bizdocking_event.ListDockingEventsResponse]
	getDockingEvent    endpoint.Endpoint[*bizdocking_event.GetDockingEventRequest, *bizdocking_event.DockingEvent]
}

func (c *endpoints) ListDockingEvents(ctx context.Context, req *bizdocking_event.ListDockingEventsRequest) (*bizdocking_event.ListDockingEventsResponse, error) {
	return c.listListDockingEvents(ctx, req)
}

func (c *endpoints) GetDockingEvent(ctx context.Context, req *bizdocking_event.GetDockingEventRequest) (*bizdocking_event.DockingEvent, error) {
	return c.getDockingEvent(ctx, req)
}
