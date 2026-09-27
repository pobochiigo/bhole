package program

import "context"

type Service interface {
	ListPrograms(ctx context.Context, req *ListProgramsRequest) (*ListProgramsResponse, error)
	GetProgram(ctx context.Context, req *GetProgramRequest) (*Program, error)
}
