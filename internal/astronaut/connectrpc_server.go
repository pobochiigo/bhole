package astronaut

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	astronautv1 "com.gitlab/pobochiigo/bhole/proto/astronaut/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/astronaut/v1/astronautv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListAstronauts transport.Handler[astronautv1.ListAstronautsRequest, astronautv1.ListAstronautsResponse]
	getAstronaut    transport.Handler[astronautv1.GetAstronautRequest, astronautv1.GetAstronautResponse]
}

func (s *server) ListAstronauts(ctx context.Context, req *connect.Request[astronautv1.ListAstronautsRequest]) (*connect.Response[astronautv1.ListAstronautsResponse], error) {
	return s.listListAstronauts(ctx, req)
}

func (s *server) GetAstronaut(ctx context.Context, req *connect.Request[astronautv1.GetAstronautRequest]) (*connect.Response[astronautv1.GetAstronautResponse], error) {
	return s.getAstronaut(ctx, req)
}

func NewAstronautHandler(svc Service) v1connect.AstronautServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListAstronauts: transport.NewConnectServer(
			eps.listListAstronauts,
			decodeListAstronautsRequest,
			encodeListAstronautsResponse,
		),
		getAstronaut: transport.NewConnectServer(
			eps.getAstronaut,
			decodeGetAstronautRequest,
			encodeGetAstronautResponse,
		),
	}
}

func decodeListAstronautsRequest(_ context.Context, req *astronautv1.ListAstronautsRequest) (*ListAstronautsRequest, error) {
	return &ListAstronautsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListAstronautsResponse(ctx context.Context, resp *ListAstronautsResponse) (*astronautv1.ListAstronautsResponse, error) {
	results := make([]*astronautv1.Astronaut, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoAstronaut(&resp.Results[i])
	}
	return &astronautv1.ListAstronautsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetAstronautRequest(_ context.Context, req *astronautv1.GetAstronautRequest) (*GetAstronautRequest, error) {
	return &GetAstronautRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetAstronautResponse(ctx context.Context, resp *Astronaut) (*astronautv1.GetAstronautResponse, error) {
	return &astronautv1.GetAstronautResponse{
		Astronaut: mapBizToProtoAstronaut(resp),
	}, nil
}

func mapBizToProtoAgencyMini(r *AgencyMini) *astronautv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &astronautv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyType(r *AgencyType) *astronautv1.AgencyType {
	if r == nil {
		return nil
	}
	return &astronautv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronaut(r *Astronaut) *astronautv1.Astronaut {
	if r == nil {
		return nil
	}
	return &astronautv1.Astronaut{
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
		Nationality: func() []*astronautv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*astronautv1.Country, len(r.Nationality))
			for i := range r.Nationality {
				res[i] = mapBizToProtoCountry(&r.Nationality[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*astronautv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*astronautv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAstronautStatus(r *AstronautStatus) *astronautv1.AstronautStatus {
	if r == nil {
		return nil
	}
	return &astronautv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronautType(r *AstronautType) *astronautv1.AstronautType {
	if r == nil {
		return nil
	}
	return &astronautv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *astronautv1.Country {
	if r == nil {
		return nil
	}
	return &astronautv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoImage(r *Image) *astronautv1.Image {
	if r == nil {
		return nil
	}
	return &astronautv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*astronautv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*astronautv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *astronautv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &astronautv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *astronautv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &astronautv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *astronautv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &astronautv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *astronautv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &astronautv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *astronautv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &astronautv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

