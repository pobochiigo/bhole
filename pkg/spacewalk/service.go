package spacewalk

import "context"

type Service interface {
	ListSpacewalks(ctx context.Context, req *ListSpacewalksRequest) (*ListSpacewalksResponse, error)
	GetSpacewalk(ctx context.Context, req *GetSpacewalkRequest) (*Spacewalk, error)
}
