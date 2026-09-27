package expedition

import "context"

type Service interface {
	ListExpeditions(ctx context.Context, req *ListExpeditionsRequest) (*ListExpeditionsResponse, error)
	GetExpedition(ctx context.Context, req *GetExpeditionRequest) (*Expedition, error)
}
