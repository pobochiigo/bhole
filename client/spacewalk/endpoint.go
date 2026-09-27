package spacewalk

import (
	"context"

	bizspacewalk "github.com/pobochiigo/bhole/spacewalk"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListSpacewalks endpoint.Endpoint[*bizspacewalk.ListSpacewalksRequest, *bizspacewalk.ListSpacewalksResponse]
	getSpacewalk       endpoint.Endpoint[*bizspacewalk.GetSpacewalkRequest, *bizspacewalk.Spacewalk]
}

func (c *endpoints) ListSpacewalks(ctx context.Context, req *bizspacewalk.ListSpacewalksRequest) (*bizspacewalk.ListSpacewalksResponse, error) {
	return c.listListSpacewalks(ctx, req)
}

func (c *endpoints) GetSpacewalk(ctx context.Context, req *bizspacewalk.GetSpacewalkRequest) (*bizspacewalk.Spacewalk, error) {
	return c.getSpacewalk(ctx, req)
}
