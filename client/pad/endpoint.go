package pad

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizpad "com.gitlab/pobochiigo/bhole/internal/pad"
)

type endpoints struct {
	listListPads endpoint.Endpoint[*bizpad.ListPadsRequest, *bizpad.ListPadsResponse]
	getPad    endpoint.Endpoint[*bizpad.GetPadRequest, *bizpad.Pad]
}

func (c *endpoints) ListPads(ctx context.Context, req *bizpad.ListPadsRequest) (*bizpad.ListPadsResponse, error) {
	return c.listListPads(ctx, req)
}

func (c *endpoints) GetPad(ctx context.Context, req *bizpad.GetPadRequest) (*bizpad.Pad, error) {
	return c.getPad(ctx, req)
}
