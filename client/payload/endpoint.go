package payload

import (
	"context"

	bizpayload "github.com/pobochiigo/bhole/pkg/payload"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListPayloads endpoint.Endpoint[*bizpayload.ListPayloadsRequest, *bizpayload.ListPayloadsResponse]
	getPayload       endpoint.Endpoint[*bizpayload.GetPayloadRequest, *bizpayload.Payload]
}

func (c *endpoints) ListPayloads(ctx context.Context, req *bizpayload.ListPayloadsRequest) (*bizpayload.ListPayloadsResponse, error) {
	return c.listListPayloads(ctx, req)
}

func (c *endpoints) GetPayload(ctx context.Context, req *bizpayload.GetPayloadRequest) (*bizpayload.Payload, error) {
	return c.getPayload(ctx, req)
}
