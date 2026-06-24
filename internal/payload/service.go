package payload

import "context"

type Service interface {
	ListPayloads(ctx context.Context, req *ListPayloadsRequest) (*ListPayloadsResponse, error)
	GetPayload(ctx context.Context, req *GetPayloadRequest) (*Payload, error)
}
