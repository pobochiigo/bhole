package spacecraft

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizspacecraft "com.gitlab/pobochiigo/bhole/internal/spacecraft"
)

type endpoints struct {
	listListSpacecrafts endpoint.Endpoint[*bizspacecraft.ListSpacecraftsRequest, *bizspacecraft.ListSpacecraftsResponse]
	getSpacecraft    endpoint.Endpoint[*bizspacecraft.GetSpacecraftRequest, *bizspacecraft.Spacecraft]
}

func (c *endpoints) ListSpacecrafts(ctx context.Context, req *bizspacecraft.ListSpacecraftsRequest) (*bizspacecraft.ListSpacecraftsResponse, error) {
	return c.listListSpacecrafts(ctx, req)
}

func (c *endpoints) GetSpacecraft(ctx context.Context, req *bizspacecraft.GetSpacecraftRequest) (*bizspacecraft.Spacecraft, error) {
	return c.getSpacecraft(ctx, req)
}
