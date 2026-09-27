package expedition

import (
	"context"

	bizexpedition "github.com/pobochiigo/bhole/pkg/expedition"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListExpeditions endpoint.Endpoint[*bizexpedition.ListExpeditionsRequest, *bizexpedition.ListExpeditionsResponse]
	getExpedition       endpoint.Endpoint[*bizexpedition.GetExpeditionRequest, *bizexpedition.Expedition]
}

func (c *endpoints) ListExpeditions(ctx context.Context, req *bizexpedition.ListExpeditionsRequest) (*bizexpedition.ListExpeditionsResponse, error) {
	return c.listListExpeditions(ctx, req)
}

func (c *endpoints) GetExpedition(ctx context.Context, req *bizexpedition.GetExpeditionRequest) (*bizexpedition.Expedition, error) {
	return c.getExpedition(ctx, req)
}
