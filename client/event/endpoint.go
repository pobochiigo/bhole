package event

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
	bizevent "github.com/pobochiigo/bhole/internal/event"
)

type endpoints struct {
	listListEvents endpoint.Endpoint[*bizevent.ListEventsRequest, *bizevent.ListEventsResponse]
	getEvent    endpoint.Endpoint[*bizevent.GetEventRequest, *bizevent.Event]
}

func (c *endpoints) ListEvents(ctx context.Context, req *bizevent.ListEventsRequest) (*bizevent.ListEventsResponse, error) {
	return c.listListEvents(ctx, req)
}

func (c *endpoints) GetEvent(ctx context.Context, req *bizevent.GetEventRequest) (*bizevent.Event, error) {
	return c.getEvent(ctx, req)
}
