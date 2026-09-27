package event

import "context"

type Service interface {
	ListEvents(ctx context.Context, req *ListEventsRequest) (*ListEventsResponse, error)
	GetEvent(ctx context.Context, req *GetEventRequest) (*Event, error)
}
