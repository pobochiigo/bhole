package update

import "context"

type Service interface {
	ListUpdates(ctx context.Context, req *ListUpdatesRequest) (*ListUpdatesResponse, error)
	GetUpdate(ctx context.Context, req *GetUpdateRequest) (*Update, error)
}
