package spacecraft

import "context"

type Service interface {
	ListSpacecrafts(ctx context.Context, req *ListSpacecraftsRequest) (*ListSpacecraftsResponse, error)
	GetSpacecraft(ctx context.Context, req *GetSpacecraftRequest) (*Spacecraft, error)
}
