package pad

import "context"

type Service interface {
	ListPads(ctx context.Context, req *ListPadsRequest) (*ListPadsResponse, error)
	GetPad(ctx context.Context, req *GetPadRequest) (*Pad, error)
}
