package program

import (
	"context"

	"github.com/pobochiigo/bhole/internal/transport"
	programv1 "github.com/pobochiigo/bhole/proto/program/v1"
	v1connect "github.com/pobochiigo/bhole/proto/program/v1/programv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListPrograms transport.Handler[programv1.ListProgramsRequest, programv1.ListProgramsResponse]
	getProgram    transport.Handler[programv1.GetProgramRequest, programv1.GetProgramResponse]
}

func (s *server) ListPrograms(ctx context.Context, req *connect.Request[programv1.ListProgramsRequest]) (*connect.Response[programv1.ListProgramsResponse], error) {
	return s.listListPrograms(ctx, req)
}

func (s *server) GetProgram(ctx context.Context, req *connect.Request[programv1.GetProgramRequest]) (*connect.Response[programv1.GetProgramResponse], error) {
	return s.getProgram(ctx, req)
}

func NewProgramHandler(svc Service) v1connect.ProgramServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListPrograms: transport.NewConnectServer(
			eps.listListPrograms,
			decodeListProgramsRequest,
			encodeListProgramsResponse,
		),
		getProgram: transport.NewConnectServer(
			eps.getProgram,
			decodeGetProgramRequest,
			encodeGetProgramResponse,
		),
	}
}

func decodeListProgramsRequest(_ context.Context, req *programv1.ListProgramsRequest) (*ListProgramsRequest, error) {
	return &ListProgramsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListProgramsResponse(ctx context.Context, resp *ListProgramsResponse) (*programv1.ListProgramsResponse, error) {
	results := make([]*programv1.Program, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoProgram(&resp.Results[i])
	}
	return &programv1.ListProgramsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetProgramRequest(_ context.Context, req *programv1.GetProgramRequest) (*GetProgramRequest, error) {
	return &GetProgramRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetProgramResponse(ctx context.Context, resp *Program) (*programv1.GetProgramResponse, error) {
	return &programv1.GetProgramResponse{
		Program: mapBizToProtoProgram(resp),
	}, nil
}

func mapBizToProtoAgencyMini(r *AgencyMini) *programv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &programv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyType(r *AgencyType) *programv1.AgencyType {
	if r == nil {
		return nil
	}
	return &programv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoImage(r *Image) *programv1.Image {
	if r == nil {
		return nil
	}
	return &programv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*programv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*programv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *programv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &programv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *programv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &programv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *programv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &programv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *programv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &programv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoProgram(r *Program) *programv1.Program {
	if r == nil {
		return nil
	}
	return &programv1.Program{
		Agencies: func() []*programv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*programv1.AgencyMini, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyMini(&r.Agencies[i])
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*programv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*programv1.MissionPatch, len(r.MissionPatches))
			for i := range r.MissionPatches {
				res[i] = mapBizToProtoMissionPatch(&r.MissionPatches[i])
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapBizToProtoProgramType(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoProgramType(r *ProgramType) *programv1.ProgramType {
	if r == nil {
		return nil
	}
	return &programv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

