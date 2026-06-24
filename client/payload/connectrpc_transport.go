package payload

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizpayload "com.gitlab/pobochiigo/bhole/internal/payload"
	payloadv1 "com.gitlab/pobochiigo/bhole/proto/payload/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/payload/v1/payloadv1connect"
	"connectrpc.com/connect"
)

func NewPayloadClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizpayload.Service {
	connectClient := v1connect.NewPayloadServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListPayloads: transport.NewConnectClient(
			connectClient.ListPayloads,
			encodeListPayloadsRequest,
			decodeListPayloadsResponse,
		),
		getPayload: transport.NewConnectClient(
			connectClient.GetPayload,
			encodeGetPayloadRequest,
			decodeGetPayloadResponse,
		),
	}
}

func encodeListPayloadsRequest(_ context.Context, req *bizpayload.ListPayloadsRequest) (*payloadv1.ListPayloadsRequest, error) {
	return &payloadv1.ListPayloadsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetPayloadRequest(_ context.Context, req *bizpayload.GetPayloadRequest) (*payloadv1.GetPayloadRequest, error) {
	return &payloadv1.GetPayloadRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListPayloadsResponse(ctx context.Context, resp *payloadv1.ListPayloadsResponse) (*bizpayload.ListPayloadsResponse, error) {
	results := make([]bizpayload.Payload, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizPayload(r)
	}
	return &bizpayload.ListPayloadsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetPayloadResponse(ctx context.Context, resp *payloadv1.GetPayloadResponse) (*bizpayload.Payload, error) {
	if resp.Payload == nil {
		return nil, nil
	}
	return mapProtoToBizPayload(resp.Payload), nil
}

func mapProtoToBizAgencyDetailed(r *payloadv1.AgencyDetailed) *bizpayload.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizpayload.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizpayload.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizpayload.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = *mapProtoToBizCountry(v)
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
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapProtoToBizImage(r.SocialLogo),
		SocialMediaLinks: func() []bizpayload.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizpayload.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = *mapProtoToBizSocialMediaLink(v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizAgencyMini(r *payloadv1.AgencyMini) *bizpayload.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizpayload.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyType(r *payloadv1.AgencyType) *bizpayload.AgencyType {
	if r == nil {
		return nil
	}
	return &bizpayload.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *payloadv1.Country) *bizpayload.Country {
	if r == nil {
		return nil
	}
	return &bizpayload.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizImage(r *payloadv1.Image) *bizpayload.Image {
	if r == nil {
		return nil
	}
	return &bizpayload.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizpayload.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizpayload.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *payloadv1.ImageLicense) *bizpayload.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizpayload.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *payloadv1.ImageVariant) *bizpayload.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizpayload.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *payloadv1.ImageVariantType) *bizpayload.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizpayload.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizMissionPatch(r *payloadv1.MissionPatch) *bizpayload.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizpayload.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizPayload(r *payloadv1.Payload) *bizpayload.Payload {
	if r == nil {
		return nil
	}
	return &bizpayload.Payload{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapProtoToBizAgencyDetailed(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapProtoToBizAgencyDetailed(r.Operator),
		Program: func() []bizpayload.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizpayload.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = *mapProtoToBizProgramNormal(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizPayloadType(r.Type),
		WikiLink: r.WikiLink,
	}
}

func mapProtoToBizPayloadType(r *payloadv1.PayloadType) *bizpayload.PayloadType {
	if r == nil {
		return nil
	}
	return &bizpayload.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizProgramNormal(r *payloadv1.ProgramNormal) *bizpayload.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizpayload.ProgramNormal{
		Agencies: func() []bizpayload.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizpayload.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizpayload.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizpayload.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *payloadv1.ProgramType) *bizpayload.ProgramType {
	if r == nil {
		return nil
	}
	return &bizpayload.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSocialMedia(r *payloadv1.SocialMedia) *bizpayload.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizpayload.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *payloadv1.SocialMediaLink) *bizpayload.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizpayload.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

