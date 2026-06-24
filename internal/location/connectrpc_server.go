package location

import (
	"context"

	"github.com/pobochiigo/bhole/internal/transport"
	locationv1 "github.com/pobochiigo/bhole/proto/location/v1"
	v1connect "github.com/pobochiigo/bhole/proto/location/v1/locationv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListLocations transport.Handler[locationv1.ListLocationsRequest, locationv1.ListLocationsResponse]
	getLocation    transport.Handler[locationv1.GetLocationRequest, locationv1.GetLocationResponse]
}

func (s *server) ListLocations(ctx context.Context, req *connect.Request[locationv1.ListLocationsRequest]) (*connect.Response[locationv1.ListLocationsResponse], error) {
	return s.listListLocations(ctx, req)
}

func (s *server) GetLocation(ctx context.Context, req *connect.Request[locationv1.GetLocationRequest]) (*connect.Response[locationv1.GetLocationResponse], error) {
	return s.getLocation(ctx, req)
}

func NewLocationHandler(svc Service) v1connect.LocationServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListLocations: transport.NewConnectServer(
			eps.listListLocations,
			decodeListLocationsRequest,
			encodeListLocationsResponse,
		),
		getLocation: transport.NewConnectServer(
			eps.getLocation,
			decodeGetLocationRequest,
			encodeGetLocationResponse,
		),
	}
}

func decodeListLocationsRequest(_ context.Context, req *locationv1.ListLocationsRequest) (*ListLocationsRequest, error) {
	return &ListLocationsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListLocationsResponse(ctx context.Context, resp *ListLocationsResponse) (*locationv1.ListLocationsResponse, error) {
	results := make([]*locationv1.Location, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoLocation(&resp.Results[i])
	}
	return &locationv1.ListLocationsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLocationRequest(_ context.Context, req *locationv1.GetLocationRequest) (*GetLocationRequest, error) {
	return &GetLocationRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetLocationResponse(ctx context.Context, resp *Location) (*locationv1.GetLocationResponse, error) {
	return &locationv1.GetLocationResponse{
		Location: mapBizToProtoLocation(resp),
	}, nil
}

func mapBizToProtoAgencyMini(r *AgencyMini) *locationv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &locationv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyType(r *AgencyType) *locationv1.AgencyType {
	if r == nil {
		return nil
	}
	return &locationv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCelestialBodyDetailed(r *CelestialBodyDetailed) *locationv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &locationv1.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		LengthOfDay: r.LengthOfDay,
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

func mapBizToProtoCelestialBodyType(r *CelestialBodyType) *locationv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &locationv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *locationv1.Country {
	if r == nil {
		return nil
	}
	return &locationv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoImage(r *Image) *locationv1.Image {
	if r == nil {
		return nil
	}
	return &locationv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*locationv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*locationv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *locationv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &locationv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *locationv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &locationv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *locationv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &locationv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLocation(r *Location) *locationv1.Location {
	if r == nil {
		return nil
	}
	return &locationv1.Location{
		Active: r.Active,
		CelestialBody: mapBizToProtoCelestialBodyDetailed(r.CelestialBody),
		Country: mapBizToProtoCountry(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		Pads: func() []*locationv1.PadSerializerNoLocation {
			if r.Pads == nil {
				return nil
			}
			res := make([]*locationv1.PadSerializerNoLocation, len(r.Pads))
			for i := range r.Pads {
				res[i] = mapBizToProtoPadSerializerNoLocation(&r.Pads[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
}

func mapBizToProtoPadSerializerNoLocation(r *PadSerializerNoLocation) *locationv1.PadSerializerNoLocation {
	if r == nil {
		return nil
	}
	return &locationv1.PadSerializerNoLocation{
		Active: r.Active,
		Agencies: func() []*locationv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*locationv1.AgencyMini, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyMini(&r.Agencies[i])
			}
			return res
		}(),
		Country: mapBizToProtoCountry(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		MapUrl: r.MapUrl,
		Name: r.Name,
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

