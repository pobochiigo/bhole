package landing

import "context"

type Service interface {
	ListLandings(ctx context.Context, req *ListLandingsRequest) (*ListLandingsResponse, error)
	GetLanding(ctx context.Context, req *GetLandingRequest) (*Landing, error)
}
