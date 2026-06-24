package pad

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizpad "github.com/pobochiigo/bhole/internal/pad"
	padv1 "github.com/pobochiigo/bhole/proto/pad/v1"
	v1connect "github.com/pobochiigo/bhole/proto/pad/v1/padv1connect"
	"connectrpc.com/connect"
)

func NewPadClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizpad.Service {
	connectClient := v1connect.NewPadServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListPads: transport.NewConnectClient(
			connectClient.ListPads,
			encodeListPadsRequest,
			decodeListPadsResponse,
		),
		getPad: transport.NewConnectClient(
			connectClient.GetPad,
			encodeGetPadRequest,
			decodeGetPadResponse,
		),
	}
}

func encodeListPadsRequest(_ context.Context, req *bizpad.ListPadsRequest) (*padv1.ListPadsRequest, error) {
	return &padv1.ListPadsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetPadRequest(_ context.Context, req *bizpad.GetPadRequest) (*padv1.GetPadRequest, error) {
	return &padv1.GetPadRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListPadsResponse(ctx context.Context, resp *padv1.ListPadsResponse) (*bizpad.ListPadsResponse, error) {
	results := make([]bizpad.Pad, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizPad(r)
	}
	return &bizpad.ListPadsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetPadResponse(ctx context.Context, resp *padv1.GetPadResponse) (*bizpad.Pad, error) {
	if resp.Pad == nil {
		return nil, nil
	}
	return mapProtoToBizPad(resp.Pad), nil
}

func mapProtoToBizAgencyNormal(r *padv1.AgencyNormal) *bizpad.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizpad.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizpad.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizpad.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *padv1.AgencyType) *bizpad.AgencyType {
	if r == nil {
		return nil
	}
	return &bizpad.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCelestialBodyDetailed(r *padv1.CelestialBodyDetailed) *bizpad.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &bizpad.CelestialBodyDetailed{
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

func mapProtoToBizCelestialBodyType(r *padv1.CelestialBodyType) *bizpad.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizpad.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *padv1.Country) *bizpad.Country {
	if r == nil {
		return nil
	}
	return &bizpad.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizImage(r *padv1.Image) *bizpad.Image {
	if r == nil {
		return nil
	}
	return &bizpad.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizpad.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizpad.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *padv1.ImageLicense) *bizpad.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizpad.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *padv1.ImageVariant) *bizpad.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizpad.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *padv1.ImageVariantType) *bizpad.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizpad.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLocation(r *padv1.Location) *bizpad.Location {
	if r == nil {
		return nil
	}
	return &bizpad.Location{
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
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
}

func mapProtoToBizPad(r *padv1.Pad) *bizpad.Pad {
	if r == nil {
		return nil
	}
	return &bizpad.Pad{
		Active: r.Active,
		Agencies: func() []bizpad.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizpad.AgencyNormal, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyNormal(v)
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
		Location: mapProtoToBizLocation(r.Location),
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

