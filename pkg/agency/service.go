package agency

import "context"

type Service interface {
	ListAgencies(ctx context.Context, req *ListAgenciesRequest) (*ListAgenciesResponse, error)
	GetAgency(ctx context.Context, req *GetAgencyRequest) (*Agency, error)
}
