package expedition

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizexpedition "com.gitlab/pobochiigo/bhole/internal/expedition"
	expeditionv1 "com.gitlab/pobochiigo/bhole/proto/expedition/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/expedition/v1/expeditionv1connect"
	"connectrpc.com/connect"
)

func NewExpeditionClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizexpedition.Service {
	connectClient := v1connect.NewExpeditionServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListExpeditions: transport.NewConnectClient(
			connectClient.ListExpeditions,
			encodeListExpeditionsRequest,
			decodeListExpeditionsResponse,
		),
		getExpedition: transport.NewConnectClient(
			connectClient.GetExpedition,
			encodeGetExpeditionRequest,
			decodeGetExpeditionResponse,
		),
	}
}

func encodeListExpeditionsRequest(_ context.Context, req *bizexpedition.ListExpeditionsRequest) (*expeditionv1.ListExpeditionsRequest, error) {
	return &expeditionv1.ListExpeditionsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetExpeditionRequest(_ context.Context, req *bizexpedition.GetExpeditionRequest) (*expeditionv1.GetExpeditionRequest, error) {
	return &expeditionv1.GetExpeditionRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListExpeditionsResponse(ctx context.Context, resp *expeditionv1.ListExpeditionsResponse) (*bizexpedition.ListExpeditionsResponse, error) {
	results := make([]bizexpedition.Expedition, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizExpedition(r)
	}
	return &bizexpedition.ListExpeditionsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetExpeditionResponse(ctx context.Context, resp *expeditionv1.GetExpeditionResponse) (*bizexpedition.Expedition, error) {
	if resp.Expedition == nil {
		return nil, nil
	}
	return mapProtoToBizExpedition(resp.Expedition), nil
}

func mapProtoToBizAgencyMini(r *expeditionv1.AgencyMini) *bizexpedition.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizexpedition.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *expeditionv1.AgencyNormal) *bizexpedition.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizexpedition.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizexpedition.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizexpedition.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = *mapProtoToBizCountry(v)
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Launchers: r.Launchers,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapProtoToBizImage(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyType(r *expeditionv1.AgencyType) *bizexpedition.AgencyType {
	if r == nil {
		return nil
	}
	return &bizexpedition.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautDetailed(r *expeditionv1.AstronautDetailed) *bizexpedition.AstronautDetailed {
	if r == nil {
		return nil
	}
	return &bizexpedition.AstronautDetailed{
		Age: r.Age,
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Bio: r.Bio,
		DateOfBirth: r.DateOfBirth,
		DateOfDeath: r.DateOfDeath,
		EvaTime: r.EvaTime,
		FirstFlight: r.FirstFlight,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InSpace: r.InSpace,
		LastFlight: r.LastFlight,
		Name: r.Name,
		Nationality: func() []bizexpedition.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]bizexpedition.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = *mapProtoToBizCountry(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []bizexpedition.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizexpedition.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = *mapProtoToBizSocialMediaLink(v)
			}
			return res
		}(),
		Status: mapProtoToBizAstronautStatus(r.Status),
		TimeInSpace: r.TimeInSpace,
		TypeVal: mapProtoToBizAstronautType(r.Type),
		Url: r.Url,
		Wiki: r.Wiki,
	}
}

func mapProtoToBizAstronautFlight(r *expeditionv1.AstronautFlight) *bizexpedition.AstronautFlight {
	if r == nil {
		return nil
	}
	return &bizexpedition.AstronautFlight{
		Astronaut: mapProtoToBizAstronautDetailed(r.Astronaut),
		Id: r.Id,
		Role: mapProtoToBizAstronautRole(r.Role),
	}
}

func mapProtoToBizAstronautRole(r *expeditionv1.AstronautRole) *bizexpedition.AstronautRole {
	if r == nil {
		return nil
	}
	return &bizexpedition.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
}

func mapProtoToBizAstronautStatus(r *expeditionv1.AstronautStatus) *bizexpedition.AstronautStatus {
	if r == nil {
		return nil
	}
	return &bizexpedition.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautType(r *expeditionv1.AstronautType) *bizexpedition.AstronautType {
	if r == nil {
		return nil
	}
	return &bizexpedition.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *expeditionv1.Country) *bizexpedition.Country {
	if r == nil {
		return nil
	}
	return &bizexpedition.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizExpedition(r *expeditionv1.Expedition) *bizexpedition.Expedition {
	if r == nil {
		return nil
	}
	return &bizexpedition.Expedition{
		Crew: func() []bizexpedition.AstronautFlight {
			if r.Crew == nil {
				return nil
			}
			res := make([]bizexpedition.AstronautFlight, len(r.Crew))
			for i, v := range r.Crew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []bizexpedition.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizexpedition.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = *mapProtoToBizMissionPatch(v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Spacestation: mapProtoToBizSpaceStationDetailed(r.Spacestation),
		Spacewalks: func() []bizexpedition.SpacewalkList {
			if r.Spacewalks == nil {
				return nil
			}
			res := make([]bizexpedition.SpacewalkList, len(r.Spacewalks))
			for i, v := range r.Spacewalks {
				res[i] = *mapProtoToBizSpacewalkList(v)
			}
			return res
		}(),
		Start: r.Start,
		Url: r.Url,
	}
}

func mapProtoToBizImage(r *expeditionv1.Image) *bizexpedition.Image {
	if r == nil {
		return nil
	}
	return &bizexpedition.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizexpedition.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizexpedition.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *expeditionv1.ImageLicense) *bizexpedition.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizexpedition.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *expeditionv1.ImageVariant) *bizexpedition.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizexpedition.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *expeditionv1.ImageVariantType) *bizexpedition.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizexpedition.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizMissionPatch(r *expeditionv1.MissionPatch) *bizexpedition.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizexpedition.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSocialMedia(r *expeditionv1.SocialMedia) *bizexpedition.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizexpedition.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *expeditionv1.SocialMediaLink) *bizexpedition.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizexpedition.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationDetailed(r *expeditionv1.SpaceStationDetailed) *bizexpedition.SpaceStationDetailed {
	if r == nil {
		return nil
	}
	return &bizexpedition.SpaceStationDetailed{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Owners: func() []bizexpedition.AgencyNormal {
			if r.Owners == nil {
				return nil
			}
			res := make([]bizexpedition.AgencyNormal, len(r.Owners))
			for i, v := range r.Owners {
				res[i] = *mapProtoToBizAgencyNormal(v)
			}
			return res
		}(),
		Status: mapProtoToBizSpaceStationStatus(r.Status),
		TypeVal: mapProtoToBizSpaceStationType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationStatus(r *expeditionv1.SpaceStationStatus) *bizexpedition.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &bizexpedition.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpaceStationType(r *expeditionv1.SpaceStationType) *bizexpedition.SpaceStationType {
	if r == nil {
		return nil
	}
	return &bizexpedition.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacewalkList(r *expeditionv1.SpacewalkList) *bizexpedition.SpacewalkList {
	if r == nil {
		return nil
	}
	return &bizexpedition.SpacewalkList{
		Duration: r.Duration,
		End: r.End,
		Id: r.Id,
		Location: r.Location,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Start: r.Start,
		Url: r.Url,
	}
}

