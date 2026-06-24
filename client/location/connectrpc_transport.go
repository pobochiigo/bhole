package location

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizlocation "github.com/pobochiigo/bhole/internal/location"
	locationv1 "github.com/pobochiigo/bhole/proto/location/v1"
	v1connect "github.com/pobochiigo/bhole/proto/location/v1/locationv1connect"
	"connectrpc.com/connect"
)

func NewLocationClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizlocation.Service {
	connectClient := v1connect.NewLocationServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListLocations: transport.NewConnectClient(
			connectClient.ListLocations,
			encodeListLocationsRequest,
			decodeListLocationsResponse,
		),
		getLocation: transport.NewConnectClient(
			connectClient.GetLocation,
			encodeGetLocationRequest,
			decodeGetLocationResponse,
		),
	}
}

func encodeListLocationsRequest(_ context.Context, req *bizlocation.ListLocationsRequest) (*locationv1.ListLocationsRequest, error) {
	return &locationv1.ListLocationsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetLocationRequest(_ context.Context, req *bizlocation.GetLocationRequest) (*locationv1.GetLocationRequest, error) {
	return &locationv1.GetLocationRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListLocationsResponse(ctx context.Context, resp *locationv1.ListLocationsResponse) (*bizlocation.ListLocationsResponse, error) {
	results := make([]bizlocation.Location, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizLocation(r)
	}
	return &bizlocation.ListLocationsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLocationResponse(ctx context.Context, resp *locationv1.GetLocationResponse) (*bizlocation.Location, error) {
	if resp.Location == nil {
		return nil, nil
	}
	return mapProtoToBizLocation(resp.Location), nil
}

func mapProtoToBizAgencyMini(r *locationv1.AgencyMini) *bizlocation.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizlocation.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyType(r *locationv1.AgencyType) *bizlocation.AgencyType {
	if r == nil {
		return nil
	}
	return &bizlocation.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCelestialBodyDetailed(r *locationv1.CelestialBodyDetailed) *bizlocation.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &bizlocation.CelestialBodyDetailed{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		LengthOfDay: r.LengthOfDay,
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

func mapProtoToBizCelestialBodyType(r *locationv1.CelestialBodyType) *bizlocation.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizlocation.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *locationv1.Country) *bizlocation.Country {
	if r == nil {
		return nil
	}
	return &bizlocation.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizImage(r *locationv1.Image) *bizlocation.Image {
	if r == nil {
		return nil
	}
	return &bizlocation.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizlocation.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizlocation.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *locationv1.ImageLicense) *bizlocation.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizlocation.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *locationv1.ImageVariant) *bizlocation.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizlocation.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *locationv1.ImageVariantType) *bizlocation.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizlocation.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLocation(r *locationv1.Location) *bizlocation.Location {
	if r == nil {
		return nil
	}
	return &bizlocation.Location{
		Active: r.Active,
		CelestialBody: mapProtoToBizCelestialBodyDetailed(r.CelestialBody),
		Country: mapProtoToBizCountry(r.Country),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Latitude: r.Latitude,
		Longitude: r.Longitude,
		MapImage: r.MapImage,
		Name: r.Name,
		Pads: func() []bizlocation.PadSerializerNoLocation {
			if r.Pads == nil {
				return nil
			}
			res := make([]bizlocation.PadSerializerNoLocation, len(r.Pads))
			for i, v := range r.Pads {
				res[i] = *mapProtoToBizPadSerializerNoLocation(v)
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

func mapProtoToBizPadSerializerNoLocation(r *locationv1.PadSerializerNoLocation) *bizlocation.PadSerializerNoLocation {
	if r == nil {
		return nil
	}
	return &bizlocation.PadSerializerNoLocation{
		Active: r.Active,
		Agencies: func() []bizlocation.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizlocation.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyMini(v)
			}
			return res
		}(),
		Country: mapProtoToBizCountry(r.Country),
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
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

