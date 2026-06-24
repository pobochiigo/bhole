package program

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/endpoint"
)

type Endpoints struct {
	listListPrograms endpoint.Endpoint[*ListProgramsRequest, *ListProgramsResponse]
	getProgram    endpoint.Endpoint[*GetProgramRequest, *Program]
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		listListPrograms: makeListListProgramsEndpoint(svc),
		getProgram:    makeGetProgramEndpoint(svc),
	}
}

func makeListListProgramsEndpoint(svc Service) endpoint.Endpoint[*ListProgramsRequest, *ListProgramsResponse] {
	return func(ctx context.Context, req *ListProgramsRequest) (*ListProgramsResponse, error) {
		return svc.ListPrograms(ctx, req)
	}
}

func makeGetProgramEndpoint(svc Service) endpoint.Endpoint[*GetProgramRequest, *Program] {
	return func(ctx context.Context, req *GetProgramRequest) (*Program, error) {
		return svc.GetProgram(ctx, req)
	}
}
