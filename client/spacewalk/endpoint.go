package spacewalk

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
	bizspacewalk "com.gitlab/pobochiigo/bhole/internal/spacewalk"
)

type endpoints struct {
	listListSpacewalks endpoint.Endpoint[*bizspacewalk.ListSpacewalksRequest, *bizspacewalk.ListSpacewalksResponse]
	getSpacewalk    endpoint.Endpoint[*bizspacewalk.GetSpacewalkRequest, *bizspacewalk.Spacewalk]
}

func (c *endpoints) ListSpacewalks(ctx context.Context, req *bizspacewalk.ListSpacewalksRequest) (*bizspacewalk.ListSpacewalksResponse, error) {
	return c.listListSpacewalks(ctx, req)
}

func (c *endpoints) GetSpacewalk(ctx context.Context, req *bizspacewalk.GetSpacewalkRequest) (*bizspacewalk.Spacewalk, error) {
	return c.getSpacewalk(ctx, req)
}
