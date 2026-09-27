package program

import (
	"context"

	bizprogram "github.com/pobochiigo/bhole/program"
	"github.com/pobochiigo/silo/endpoint"
)

type endpoints struct {
	listListPrograms endpoint.Endpoint[*bizprogram.ListProgramsRequest, *bizprogram.ListProgramsResponse]
	getProgram       endpoint.Endpoint[*bizprogram.GetProgramRequest, *bizprogram.Program]
}

func (c *endpoints) ListPrograms(ctx context.Context, req *bizprogram.ListProgramsRequest) (*bizprogram.ListProgramsResponse, error) {
	return c.listListPrograms(ctx, req)
}

func (c *endpoints) GetProgram(ctx context.Context, req *bizprogram.GetProgramRequest) (*bizprogram.Program, error) {
	return c.getProgram(ctx, req)
}
