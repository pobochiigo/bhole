package celestial_body

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizcelestial_body "github.com/pobochiigo/bhole/internal/celestial_body"
	celestial_bodyv1 "github.com/pobochiigo/bhole/proto/celestial_body/v1"
	v1connect "github.com/pobochiigo/bhole/proto/celestial_body/v1/celestial_bodyv1connect"
	"connectrpc.com/connect"
)

func NewCelestialBodyClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizcelestial_body.Service {
	connectClient := v1connect.NewCelestialBodyServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListCelestialBodies: transport.NewConnectClient(
			connectClient.ListCelestialBodies,
			encodeListCelestialBodiesRequest,
			decodeListCelestialBodiesResponse,
		),
		getCelestialBody: transport.NewConnectClient(
			connectClient.GetCelestialBody,
			encodeGetCelestialBodyRequest,
			decodeGetCelestialBodyResponse,
		),
	}
}

func encodeListCelestialBodiesRequest(_ context.Context, req *bizcelestial_body.ListCelestialBodiesRequest) (*celestial_bodyv1.ListCelestialBodiesRequest, error) {
	return &celestial_bodyv1.ListCelestialBodiesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetCelestialBodyRequest(_ context.Context, req *bizcelestial_body.GetCelestialBodyRequest) (*celestial_bodyv1.GetCelestialBodyRequest, error) {
	return &celestial_bodyv1.GetCelestialBodyRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListCelestialBodiesResponse(ctx context.Context, resp *celestial_bodyv1.ListCelestialBodiesResponse) (*bizcelestial_body.ListCelestialBodiesResponse, error) {
	results := make([]bizcelestial_body.CelestialBody, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizCelestialBody(r)
	}
	return &bizcelestial_body.ListCelestialBodiesResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetCelestialBodyResponse(ctx context.Context, resp *celestial_bodyv1.GetCelestialBodyResponse) (*bizcelestial_body.CelestialBody, error) {
	if resp.CelestialBody == nil {
		return nil, nil
	}
	return mapProtoToBizCelestialBody(resp.CelestialBody), nil
}

func mapProtoToBizCelestialBody(r *celestial_bodyv1.CelestialBody) *bizcelestial_body.CelestialBody {
	if r == nil {
		return nil
	}
	return &bizcelestial_body.CelestialBody{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		LengthOfDay: r.LengthOfDay,
		Locations: func() []bizcelestial_body.LocationSerializerNoCelestialBody {
			if r.Locations == nil {
				return nil
			}
			res := make([]bizcelestial_body.LocationSerializerNoCelestialBody, len(r.Locations))
			for i, v := range r.Locations {
				res[i] = *mapProtoToBizLocationSerializerNoCelestialBody(v)
			}
			return res
		}(),
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalAttemptedLandings: r.TotalAttemptedLandings,
		TotalAttemptedLaunches: r.TotalAttemptedLaunches,
		TypeVal: mapProtoToBizCelestialBodyType(r.Type),
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizCelestialBodyType(r *celestial_bodyv1.CelestialBodyType) *bizcelestial_body.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizcelestial_body.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *celestial_bodyv1.Country) *bizcelestial_body.Country {
	if r == nil {
		return nil
	}
	return &bizcelestial_body.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizImage(r *celestial_bodyv1.Image) *bizcelestial_body.Image {
	if r == nil {
		return nil
	}
	return &bizcelestial_body.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizcelestial_body.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizcelestial_body.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *celestial_bodyv1.ImageLicense) *bizcelestial_body.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizcelestial_body.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *celestial_bodyv1.ImageVariant) *bizcelestial_body.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizcelestial_body.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *celestial_bodyv1.ImageVariantType) *bizcelestial_body.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizcelestial_body.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLocationSerializerNoCelestialBody(r *celestial_bodyv1.LocationSerializerNoCelestialBody) *bizcelestial_body.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &bizcelestial_body.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapProtoToBizCountry(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
}

