package celestial_body

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	celestial_bodyv1 "com.gitlab/pobochiigo/bhole/proto/celestial_body/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/celestial_body/v1/celestial_bodyv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListCelestialBodies transport.Handler[celestial_bodyv1.ListCelestialBodiesRequest, celestial_bodyv1.ListCelestialBodiesResponse]
	getCelestialBody    transport.Handler[celestial_bodyv1.GetCelestialBodyRequest, celestial_bodyv1.GetCelestialBodyResponse]
}

func (s *server) ListCelestialBodies(ctx context.Context, req *connect.Request[celestial_bodyv1.ListCelestialBodiesRequest]) (*connect.Response[celestial_bodyv1.ListCelestialBodiesResponse], error) {
	return s.listListCelestialBodies(ctx, req)
}

func (s *server) GetCelestialBody(ctx context.Context, req *connect.Request[celestial_bodyv1.GetCelestialBodyRequest]) (*connect.Response[celestial_bodyv1.GetCelestialBodyResponse], error) {
	return s.getCelestialBody(ctx, req)
}

func NewCelestialBodyHandler(svc Service) v1connect.CelestialBodyServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListCelestialBodies: transport.NewConnectServer(
			eps.listListCelestialBodies,
			decodeListCelestialBodiesRequest,
			encodeListCelestialBodiesResponse,
		),
		getCelestialBody: transport.NewConnectServer(
			eps.getCelestialBody,
			decodeGetCelestialBodyRequest,
			encodeGetCelestialBodyResponse,
		),
	}
}

func decodeListCelestialBodiesRequest(_ context.Context, req *celestial_bodyv1.ListCelestialBodiesRequest) (*ListCelestialBodiesRequest, error) {
	return &ListCelestialBodiesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListCelestialBodiesResponse(ctx context.Context, resp *ListCelestialBodiesResponse) (*celestial_bodyv1.ListCelestialBodiesResponse, error) {
	results := make([]*celestial_bodyv1.CelestialBody, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoCelestialBody(&resp.Results[i])
	}
	return &celestial_bodyv1.ListCelestialBodiesResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetCelestialBodyRequest(_ context.Context, req *celestial_bodyv1.GetCelestialBodyRequest) (*GetCelestialBodyRequest, error) {
	return &GetCelestialBodyRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetCelestialBodyResponse(ctx context.Context, resp *CelestialBody) (*celestial_bodyv1.GetCelestialBodyResponse, error) {
	return &celestial_bodyv1.GetCelestialBodyResponse{
		CelestialBody: mapBizToProtoCelestialBody(resp),
	}, nil
}

func mapBizToProtoCelestialBody(r *CelestialBody) *celestial_bodyv1.CelestialBody {
	if r == nil {
		return nil
	}
	return &celestial_bodyv1.CelestialBody{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		LengthOfDay: r.LengthOfDay,
		Locations: func() []*celestial_bodyv1.LocationSerializerNoCelestialBody {
			if r.Locations == nil {
				return nil
			}
			res := make([]*celestial_bodyv1.LocationSerializerNoCelestialBody, len(r.Locations))
			for i := range r.Locations {
				res[i] = mapBizToProtoLocationSerializerNoCelestialBody(&r.Locations[i])
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
		Type: mapBizToProtoCelestialBodyType(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoCelestialBodyType(r *CelestialBodyType) *celestial_bodyv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &celestial_bodyv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *celestial_bodyv1.Country {
	if r == nil {
		return nil
	}
	return &celestial_bodyv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoImage(r *Image) *celestial_bodyv1.Image {
	if r == nil {
		return nil
	}
	return &celestial_bodyv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*celestial_bodyv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*celestial_bodyv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *celestial_bodyv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &celestial_bodyv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *celestial_bodyv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &celestial_bodyv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *celestial_bodyv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &celestial_bodyv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLocationSerializerNoCelestialBody(r *LocationSerializerNoCelestialBody) *celestial_bodyv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &celestial_bodyv1.LocationSerializerNoCelestialBody{
		Active: r.Active,
		Country: mapBizToProtoCountry(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
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

