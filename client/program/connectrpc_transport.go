package program

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizprogram "github.com/pobochiigo/bhole/internal/program"
	programv1 "github.com/pobochiigo/bhole/proto/program/v1"
	v1connect "github.com/pobochiigo/bhole/proto/program/v1/programv1connect"
	"connectrpc.com/connect"
)

func NewProgramClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizprogram.Service {
	connectClient := v1connect.NewProgramServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListPrograms: transport.NewConnectClient(
			connectClient.ListPrograms,
			encodeListProgramsRequest,
			decodeListProgramsResponse,
		),
		getProgram: transport.NewConnectClient(
			connectClient.GetProgram,
			encodeGetProgramRequest,
			decodeGetProgramResponse,
		),
	}
}

func encodeListProgramsRequest(_ context.Context, req *bizprogram.ListProgramsRequest) (*programv1.ListProgramsRequest, error) {
	return &programv1.ListProgramsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetProgramRequest(_ context.Context, req *bizprogram.GetProgramRequest) (*programv1.GetProgramRequest, error) {
	return &programv1.GetProgramRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListProgramsResponse(ctx context.Context, resp *programv1.ListProgramsResponse) (*bizprogram.ListProgramsResponse, error) {
	results := make([]bizprogram.Program, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizProgram(r)
	}
	return &bizprogram.ListProgramsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetProgramResponse(ctx context.Context, resp *programv1.GetProgramResponse) (*bizprogram.Program, error) {
	if resp.Program == nil {
		return nil, nil
	}
	return mapProtoToBizProgram(resp.Program), nil
}

func mapProtoToBizAgencyMini(r *programv1.AgencyMini) *bizprogram.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizprogram.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyType(r *programv1.AgencyType) *bizprogram.AgencyType {
	if r == nil {
		return nil
	}
	return &bizprogram.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizImage(r *programv1.Image) *bizprogram.Image {
	if r == nil {
		return nil
	}
	return &bizprogram.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizprogram.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizprogram.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *programv1.ImageLicense) *bizprogram.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizprogram.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *programv1.ImageVariant) *bizprogram.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizprogram.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *programv1.ImageVariantType) *bizprogram.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizprogram.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizMissionPatch(r *programv1.MissionPatch) *bizprogram.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizprogram.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizProgram(r *programv1.Program) *bizprogram.Program {
	if r == nil {
		return nil
	}
	return &bizprogram.Program{
		Agencies: func() []bizprogram.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizprogram.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyMini(v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []bizprogram.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizprogram.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = *mapProtoToBizMissionPatch(v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		TypeVal: mapProtoToBizProgramType(r.Type),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizProgramType(r *programv1.ProgramType) *bizprogram.ProgramType {
	if r == nil {
		return nil
	}
	return &bizprogram.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

