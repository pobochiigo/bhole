package expedition

import (
	"context"

	"github.com/pobochiigo/bhole/internal/transport"
	expeditionv1 "github.com/pobochiigo/bhole/proto/expedition/v1"
	v1connect "github.com/pobochiigo/bhole/proto/expedition/v1/expeditionv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListExpeditions transport.Handler[expeditionv1.ListExpeditionsRequest, expeditionv1.ListExpeditionsResponse]
	getExpedition    transport.Handler[expeditionv1.GetExpeditionRequest, expeditionv1.GetExpeditionResponse]
}

func (s *server) ListExpeditions(ctx context.Context, req *connect.Request[expeditionv1.ListExpeditionsRequest]) (*connect.Response[expeditionv1.ListExpeditionsResponse], error) {
	return s.listListExpeditions(ctx, req)
}

func (s *server) GetExpedition(ctx context.Context, req *connect.Request[expeditionv1.GetExpeditionRequest]) (*connect.Response[expeditionv1.GetExpeditionResponse], error) {
	return s.getExpedition(ctx, req)
}

func NewExpeditionHandler(svc Service) v1connect.ExpeditionServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListExpeditions: transport.NewConnectServer(
			eps.listListExpeditions,
			decodeListExpeditionsRequest,
			encodeListExpeditionsResponse,
		),
		getExpedition: transport.NewConnectServer(
			eps.getExpedition,
			decodeGetExpeditionRequest,
			encodeGetExpeditionResponse,
		),
	}
}

func decodeListExpeditionsRequest(_ context.Context, req *expeditionv1.ListExpeditionsRequest) (*ListExpeditionsRequest, error) {
	return &ListExpeditionsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListExpeditionsResponse(ctx context.Context, resp *ListExpeditionsResponse) (*expeditionv1.ListExpeditionsResponse, error) {
	results := make([]*expeditionv1.Expedition, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoExpedition(&resp.Results[i])
	}
	return &expeditionv1.ListExpeditionsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetExpeditionRequest(_ context.Context, req *expeditionv1.GetExpeditionRequest) (*GetExpeditionRequest, error) {
	return &GetExpeditionRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetExpeditionResponse(ctx context.Context, resp *Expedition) (*expeditionv1.GetExpeditionResponse, error) {
	return &expeditionv1.GetExpeditionResponse{
		Expedition: mapBizToProtoExpedition(resp),
	}, nil
}

func mapBizToProtoAgencyMini(r *AgencyMini) *expeditionv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &expeditionv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *expeditionv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &expeditionv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*expeditionv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*expeditionv1.Country, len(r.Country))
			for i := range r.Country {
				res[i] = mapBizToProtoCountry(&r.Country[i])
			}
			return res
		}(),
		Description: r.Description,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Launchers: r.Launchers,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapBizToProtoImage(r.SocialLogo),
		Spacecraft: r.Spacecraft,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyType(r *AgencyType) *expeditionv1.AgencyType {
	if r == nil {
		return nil
	}
	return &expeditionv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronautDetailed(r *AstronautDetailed) *expeditionv1.AstronautDetailed {
	if r == nil {
		return nil
	}
	return &expeditionv1.AstronautDetailed{
		Age: r.Age,
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Bio: r.Bio,
		DateOfBirth: r.DateOfBirth,
		DateOfDeath: r.DateOfDeath,
		EvaTime: r.EvaTime,
		FirstFlight: r.FirstFlight,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InSpace: r.InSpace,
		LastFlight: r.LastFlight,
		Name: r.Name,
		Nationality: func() []*expeditionv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*expeditionv1.Country, len(r.Nationality))
			for i := range r.Nationality {
				res[i] = mapBizToProtoCountry(&r.Nationality[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*expeditionv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*expeditionv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i := range r.SocialMediaLinks {
				res[i] = mapBizToProtoSocialMediaLink(&r.SocialMediaLinks[i])
			}
			return res
		}(),
		Status: mapBizToProtoAstronautStatus(r.Status),
		TimeInSpace: r.TimeInSpace,
		Type: mapBizToProtoAstronautType(r.TypeVal),
		Url: r.Url,
		Wiki: r.Wiki,
	}
}

func mapBizToProtoAstronautFlight(r *AstronautFlight) *expeditionv1.AstronautFlight {
	if r == nil {
		return nil
	}
	return &expeditionv1.AstronautFlight{
		Astronaut: mapBizToProtoAstronautDetailed(r.Astronaut),
		Id: r.Id,
		Role: mapBizToProtoAstronautRole(r.Role),
	}
}

func mapBizToProtoAstronautRole(r *AstronautRole) *expeditionv1.AstronautRole {
	if r == nil {
		return nil
	}
	return &expeditionv1.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
}

func mapBizToProtoAstronautStatus(r *AstronautStatus) *expeditionv1.AstronautStatus {
	if r == nil {
		return nil
	}
	return &expeditionv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronautType(r *AstronautType) *expeditionv1.AstronautType {
	if r == nil {
		return nil
	}
	return &expeditionv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *expeditionv1.Country {
	if r == nil {
		return nil
	}
	return &expeditionv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoExpedition(r *Expedition) *expeditionv1.Expedition {
	if r == nil {
		return nil
	}
	return &expeditionv1.Expedition{
		Crew: func() []*expeditionv1.AstronautFlight {
			if r.Crew == nil {
				return nil
			}
			res := make([]*expeditionv1.AstronautFlight, len(r.Crew))
			for i := range r.Crew {
				res[i] = mapBizToProtoAstronautFlight(&r.Crew[i])
			}
			return res
		}(),
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []*expeditionv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*expeditionv1.MissionPatch, len(r.MissionPatches))
			for i := range r.MissionPatches {
				res[i] = mapBizToProtoMissionPatch(&r.MissionPatches[i])
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Spacestation: mapBizToProtoSpaceStationDetailed(r.Spacestation),
		Spacewalks: func() []*expeditionv1.SpacewalkList {
			if r.Spacewalks == nil {
				return nil
			}
			res := make([]*expeditionv1.SpacewalkList, len(r.Spacewalks))
			for i := range r.Spacewalks {
				res[i] = mapBizToProtoSpacewalkList(&r.Spacewalks[i])
			}
			return res
		}(),
		Start: r.Start,
		Url: r.Url,
	}
}

func mapBizToProtoImage(r *Image) *expeditionv1.Image {
	if r == nil {
		return nil
	}
	return &expeditionv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*expeditionv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*expeditionv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *expeditionv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &expeditionv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *expeditionv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &expeditionv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *expeditionv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &expeditionv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *expeditionv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &expeditionv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *expeditionv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &expeditionv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *expeditionv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &expeditionv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationDetailed(r *SpaceStationDetailed) *expeditionv1.SpaceStationDetailed {
	if r == nil {
		return nil
	}
	return &expeditionv1.SpaceStationDetailed{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Owners: func() []*expeditionv1.AgencyNormal {
			if r.Owners == nil {
				return nil
			}
			res := make([]*expeditionv1.AgencyNormal, len(r.Owners))
			for i := range r.Owners {
				res[i] = mapBizToProtoAgencyNormal(&r.Owners[i])
			}
			return res
		}(),
		Status: mapBizToProtoSpaceStationStatus(r.Status),
		Type: mapBizToProtoSpaceStationType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationStatus(r *SpaceStationStatus) *expeditionv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &expeditionv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpaceStationType(r *SpaceStationType) *expeditionv1.SpaceStationType {
	if r == nil {
		return nil
	}
	return &expeditionv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacewalkList(r *SpacewalkList) *expeditionv1.SpacewalkList {
	if r == nil {
		return nil
	}
	return &expeditionv1.SpacewalkList{
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

