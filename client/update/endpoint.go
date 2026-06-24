package update

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizupdate "com.gitlab/pobochiigo/bhole/internal/update"
)

type endpoints struct {
	listListUpdates endpoint.Endpoint[*bizupdate.ListUpdatesRequest, *bizupdate.ListUpdatesResponse]
	getUpdate    endpoint.Endpoint[*bizupdate.GetUpdateRequest, *bizupdate.Update]
}

func (c *endpoints) ListUpdates(ctx context.Context, req *bizupdate.ListUpdatesRequest) (*bizupdate.ListUpdatesResponse, error) {
	return c.listListUpdates(ctx, req)
}

func (c *endpoints) GetUpdate(ctx context.Context, req *bizupdate.GetUpdateRequest) (*bizupdate.Update, error) {
	return c.getUpdate(ctx, req)
}
