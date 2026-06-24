package pad

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	padv1 "com.gitlab/pobochiigo/bhole/proto/pad/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/pad/v1/padv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListPads transport.Handler[padv1.ListPadsRequest, padv1.ListPadsResponse]
	getPad    transport.Handler[padv1.GetPadRequest, padv1.GetPadResponse]
}

func (s *server) ListPads(ctx context.Context, req *connect.Request[padv1.ListPadsRequest]) (*connect.Response[padv1.ListPadsResponse], error) {
	return s.listListPads(ctx, req)
}

func (s *server) GetPad(ctx context.Context, req *connect.Request[padv1.GetPadRequest]) (*connect.Response[padv1.GetPadResponse], error) {
	return s.getPad(ctx, req)
}

func NewPadHandler(svc Service) v1connect.PadServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListPads: transport.NewConnectServer(
			eps.listListPads,
			decodeListPadsRequest,
			encodeListPadsResponse,
		),
		getPad: transport.NewConnectServer(
			eps.getPad,
			decodeGetPadRequest,
			encodeGetPadResponse,
		),
	}
}

func decodeListPadsRequest(_ context.Context, req *padv1.ListPadsRequest) (*ListPadsRequest, error) {
	return &ListPadsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListPadsResponse(ctx context.Context, resp *ListPadsResponse) (*padv1.ListPadsResponse, error) {
	results := make([]*padv1.Pad, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoPad(&resp.Results[i])
	}
	return &padv1.ListPadsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetPadRequest(_ context.Context, req *padv1.GetPadRequest) (*GetPadRequest, error) {
	return &GetPadRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetPadResponse(ctx context.Context, resp *Pad) (*padv1.GetPadResponse, error) {
	return &padv1.GetPadResponse{
		Pad: mapBizToProtoPad(resp),
	}, nil
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *padv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &padv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*padv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*padv1.Country, len(r.Country))
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

func mapBizToProtoAgencyType(r *AgencyType) *padv1.AgencyType {
	if r == nil {
		return nil
	}
	return &padv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCelestialBodyDetailed(r *CelestialBodyDetailed) *padv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &padv1.CelestialBodyDetailed{
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

func mapBizToProtoCelestialBodyType(r *CelestialBodyType) *padv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &padv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *padv1.Country {
	if r == nil {
		return nil
	}
	return &padv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoImage(r *Image) *padv1.Image {
	if r == nil {
		return nil
	}
	return &padv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*padv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*padv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *padv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &padv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *padv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &padv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *padv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &padv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLocation(r *Location) *padv1.Location {
	if r == nil {
		return nil
	}
	return &padv1.Location{
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
		ResponseMode: r.ResponseMode,
		TimezoneName: r.TimezoneName,
		TotalLandingCount: r.TotalLandingCount,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
	}
}

func mapBizToProtoPad(r *Pad) *padv1.Pad {
	if r == nil {
		return nil
	}
	return &padv1.Pad{
		Active: r.Active,
		Agencies: func() []*padv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*padv1.AgencyNormal, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyNormal(&r.Agencies[i])
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
		Location: mapBizToProtoLocation(r.Location),
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

