package program

import (
	"context"

	"github.com/pobochiigo/bhole/client/endpoint"
	bizprogram "github.com/pobochiigo/bhole/internal/program"
)

type endpoints struct {
	listListPrograms endpoint.Endpoint[*bizprogram.ListProgramsRequest, *bizprogram.ListProgramsResponse]
	getProgram    endpoint.Endpoint[*bizprogram.GetProgramRequest, *bizprogram.Program]
}

func (c *endpoints) ListPrograms(ctx context.Context, req *bizprogram.ListProgramsRequest) (*bizprogram.ListProgramsResponse, error) {
	return c.listListPrograms(ctx, req)
}

func (c *endpoints) GetProgram(ctx context.Context, req *bizprogram.GetProgramRequest) (*bizprogram.Program, error) {
	return c.getProgram(ctx, req)
}
