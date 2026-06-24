package astronaut

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizastronaut "github.com/pobochiigo/bhole/internal/astronaut"
	astronautv1 "github.com/pobochiigo/bhole/proto/astronaut/v1"
	v1connect "github.com/pobochiigo/bhole/proto/astronaut/v1/astronautv1connect"
	"connectrpc.com/connect"
)

func NewAstronautClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizastronaut.Service {
	connectClient := v1connect.NewAstronautServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListAstronauts: transport.NewConnectClient(
			connectClient.ListAstronauts,
			encodeListAstronautsRequest,
			decodeListAstronautsResponse,
		),
		getAstronaut: transport.NewConnectClient(
			connectClient.GetAstronaut,
			encodeGetAstronautRequest,
			decodeGetAstronautResponse,
		),
	}
}

func encodeListAstronautsRequest(_ context.Context, req *bizastronaut.ListAstronautsRequest) (*astronautv1.ListAstronautsRequest, error) {
	return &astronautv1.ListAstronautsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetAstronautRequest(_ context.Context, req *bizastronaut.GetAstronautRequest) (*astronautv1.GetAstronautRequest, error) {
	return &astronautv1.GetAstronautRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListAstronautsResponse(ctx context.Context, resp *astronautv1.ListAstronautsResponse) (*bizastronaut.ListAstronautsResponse, error) {
	results := make([]bizastronaut.Astronaut, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizAstronaut(r)
	}
	return &bizastronaut.ListAstronautsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetAstronautResponse(ctx context.Context, resp *astronautv1.GetAstronautResponse) (*bizastronaut.Astronaut, error) {
	if resp.Astronaut == nil {
		return nil, nil
	}
	return mapProtoToBizAstronaut(resp.Astronaut), nil
}

func mapProtoToBizAgencyMini(r *astronautv1.AgencyMini) *bizastronaut.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizastronaut.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyType(r *astronautv1.AgencyType) *bizastronaut.AgencyType {
	if r == nil {
		return nil
	}
	return &bizastronaut.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronaut(r *astronautv1.Astronaut) *bizastronaut.Astronaut {
	if r == nil {
		return nil
	}
	return &bizastronaut.Astronaut{
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
		Nationality: func() []bizastronaut.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]bizastronaut.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = *mapProtoToBizCountry(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []bizastronaut.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizastronaut.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAstronautStatus(r *astronautv1.AstronautStatus) *bizastronaut.AstronautStatus {
	if r == nil {
		return nil
	}
	return &bizastronaut.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautType(r *astronautv1.AstronautType) *bizastronaut.AstronautType {
	if r == nil {
		return nil
	}
	return &bizastronaut.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *astronautv1.Country) *bizastronaut.Country {
	if r == nil {
		return nil
	}
	return &bizastronaut.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizImage(r *astronautv1.Image) *bizastronaut.Image {
	if r == nil {
		return nil
	}
	return &bizastronaut.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizastronaut.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizastronaut.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *astronautv1.ImageLicense) *bizastronaut.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizastronaut.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *astronautv1.ImageVariant) *bizastronaut.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizastronaut.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *astronautv1.ImageVariantType) *bizastronaut.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizastronaut.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSocialMedia(r *astronautv1.SocialMedia) *bizastronaut.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizastronaut.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *astronautv1.SocialMediaLink) *bizastronaut.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizastronaut.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

