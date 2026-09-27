package astronaut

import "context"

type Service interface {
	ListAstronauts(ctx context.Context, req *ListAstronautsRequest) (*ListAstronautsResponse, error)
	GetAstronaut(ctx context.Context, req *GetAstronautRequest) (*Astronaut, error)
}
