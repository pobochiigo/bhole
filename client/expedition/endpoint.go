package expedition

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizexpedition "com.gitlab/pobochiigo/bhole/internal/expedition"
)

type endpoints struct {
	listListExpeditions endpoint.Endpoint[*bizexpedition.ListExpeditionsRequest, *bizexpedition.ListExpeditionsResponse]
	getExpedition    endpoint.Endpoint[*bizexpedition.GetExpeditionRequest, *bizexpedition.Expedition]
}

func (c *endpoints) ListExpeditions(ctx context.Context, req *bizexpedition.ListExpeditionsRequest) (*bizexpedition.ListExpeditionsResponse, error) {
	return c.listListExpeditions(ctx, req)
}

func (c *endpoints) GetExpedition(ctx context.Context, req *bizexpedition.GetExpeditionRequest) (*bizexpedition.Expedition, error) {
	return c.getExpedition(ctx, req)
}
