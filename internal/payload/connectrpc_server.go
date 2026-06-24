package payload

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	payloadv1 "com.gitlab/pobochiigo/bhole/proto/payload/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/payload/v1/payloadv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListPayloads transport.Handler[payloadv1.ListPayloadsRequest, payloadv1.ListPayloadsResponse]
	getPayload    transport.Handler[payloadv1.GetPayloadRequest, payloadv1.GetPayloadResponse]
}

func (s *server) ListPayloads(ctx context.Context, req *connect.Request[payloadv1.ListPayloadsRequest]) (*connect.Response[payloadv1.ListPayloadsResponse], error) {
	return s.listListPayloads(ctx, req)
}

func (s *server) GetPayload(ctx context.Context, req *connect.Request[payloadv1.GetPayloadRequest]) (*connect.Response[payloadv1.GetPayloadResponse], error) {
	return s.getPayload(ctx, req)
}

func NewPayloadHandler(svc Service) v1connect.PayloadServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListPayloads: transport.NewConnectServer(
			eps.listListPayloads,
			decodeListPayloadsRequest,
			encodeListPayloadsResponse,
		),
		getPayload: transport.NewConnectServer(
			eps.getPayload,
			decodeGetPayloadRequest,
			encodeGetPayloadResponse,
		),
	}
}

func decodeListPayloadsRequest(_ context.Context, req *payloadv1.ListPayloadsRequest) (*ListPayloadsRequest, error) {
	return &ListPayloadsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListPayloadsResponse(ctx context.Context, resp *ListPayloadsResponse) (*payloadv1.ListPayloadsResponse, error) {
	results := make([]*payloadv1.Payload, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoPayload(&resp.Results[i])
	}
	return &payloadv1.ListPayloadsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetPayloadRequest(_ context.Context, req *payloadv1.GetPayloadRequest) (*GetPayloadRequest, error) {
	return &GetPayloadRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetPayloadResponse(ctx context.Context, resp *Payload) (*payloadv1.GetPayloadResponse, error) {
	return &payloadv1.GetPayloadResponse{
		Payload: mapBizToProtoPayload(resp),
	}, nil
}

func mapBizToProtoAgencyDetailed(r *AgencyDetailed) *payloadv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &payloadv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*payloadv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*payloadv1.Country, len(r.Country))
			for i := range r.Country {
				res[i] = mapBizToProtoCountry(&r.Country[i])
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapBizToProtoImage(r.SocialLogo),
		SocialMediaLinks: func() []*payloadv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*payloadv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i := range r.SocialMediaLinks {
				res[i] = mapBizToProtoSocialMediaLink(&r.SocialMediaLinks[i])
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoAgencyMini(r *AgencyMini) *payloadv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &payloadv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyType(r *AgencyType) *payloadv1.AgencyType {
	if r == nil {
		return nil
	}
	return &payloadv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *payloadv1.Country {
	if r == nil {
		return nil
	}
	return &payloadv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoImage(r *Image) *payloadv1.Image {
	if r == nil {
		return nil
	}
	return &payloadv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*payloadv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*payloadv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *payloadv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &payloadv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *payloadv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &payloadv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *payloadv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &payloadv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *payloadv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &payloadv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoPayload(r *Payload) *payloadv1.Payload {
	if r == nil {
		return nil
	}
	return &payloadv1.Payload{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapBizToProtoAgencyDetailed(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapBizToProtoAgencyDetailed(r.Operator),
		Program: func() []*payloadv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*payloadv1.ProgramNormal, len(r.Program))
			for i := range r.Program {
				res[i] = mapBizToProtoProgramNormal(&r.Program[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoPayloadType(r.TypeVal),
		WikiLink: r.WikiLink,
	}
}

func mapBizToProtoPayloadType(r *PayloadType) *payloadv1.PayloadType {
	if r == nil {
		return nil
	}
	return &payloadv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoProgramNormal(r *ProgramNormal) *payloadv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &payloadv1.ProgramNormal{
		Agencies: func() []*payloadv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*payloadv1.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []*payloadv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*payloadv1.MissionPatch, len(r.MissionPatches))
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

func mapBizToProtoProgramType(r *ProgramType) *payloadv1.ProgramType {
	if r == nil {
		return nil
	}
	return &payloadv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *payloadv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &payloadv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *payloadv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &payloadv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

