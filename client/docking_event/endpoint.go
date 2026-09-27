package docking_event

import (
	"context"

	bizdocking_event "github.com/pobochiigo/bhole/pkg/docking_event"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListDockingEvents endpoint.Endpoint[*bizdocking_event.ListDockingEventsRequest, *bizdocking_event.ListDockingEventsResponse]
	getDockingEvent       endpoint.Endpoint[*bizdocking_event.GetDockingEventRequest, *bizdocking_event.DockingEvent]
}

func (c *endpoints) ListDockingEvents(ctx context.Context, req *bizdocking_event.ListDockingEventsRequest) (*bizdocking_event.ListDockingEventsResponse, error) {
	return c.listListDockingEvents(ctx, req)
}

func (c *endpoints) GetDockingEvent(ctx context.Context, req *bizdocking_event.GetDockingEventRequest) (*bizdocking_event.DockingEvent, error) {
	return c.getDockingEvent(ctx, req)
}
