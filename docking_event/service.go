package docking_event

import "context"

type Service interface {
	ListDockingEvents(ctx context.Context, req *ListDockingEventsRequest) (*ListDockingEventsResponse, error)
	GetDockingEvent(ctx context.Context, req *GetDockingEventRequest) (*DockingEvent, error)
}
