package spacecraft

import (
	"context"

	"github.com/pobochiigo/bhole/internal/transport"
	spacecraftv1 "github.com/pobochiigo/bhole/proto/spacecraft/v1"
	v1connect "github.com/pobochiigo/bhole/proto/spacecraft/v1/spacecraftv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListSpacecrafts transport.Handler[spacecraftv1.ListSpacecraftsRequest, spacecraftv1.ListSpacecraftsResponse]
	getSpacecraft    transport.Handler[spacecraftv1.GetSpacecraftRequest, spacecraftv1.GetSpacecraftResponse]
}

func (s *server) ListSpacecrafts(ctx context.Context, req *connect.Request[spacecraftv1.ListSpacecraftsRequest]) (*connect.Response[spacecraftv1.ListSpacecraftsResponse], error) {
	return s.listListSpacecrafts(ctx, req)
}

func (s *server) GetSpacecraft(ctx context.Context, req *connect.Request[spacecraftv1.GetSpacecraftRequest]) (*connect.Response[spacecraftv1.GetSpacecraftResponse], error) {
	return s.getSpacecraft(ctx, req)
}

func NewSpacecraftHandler(svc Service) v1connect.SpacecraftServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListSpacecrafts: transport.NewConnectServer(
			eps.listListSpacecrafts,
			decodeListSpacecraftsRequest,
			encodeListSpacecraftsResponse,
		),
		getSpacecraft: transport.NewConnectServer(
			eps.getSpacecraft,
			decodeGetSpacecraftRequest,
			encodeGetSpacecraftResponse,
		),
	}
}

func decodeListSpacecraftsRequest(_ context.Context, req *spacecraftv1.ListSpacecraftsRequest) (*ListSpacecraftsRequest, error) {
	return &ListSpacecraftsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListSpacecraftsResponse(ctx context.Context, resp *ListSpacecraftsResponse) (*spacecraftv1.ListSpacecraftsResponse, error) {
	results := make([]*spacecraftv1.Spacecraft, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoSpacecraft(&resp.Results[i])
	}
	return &spacecraftv1.ListSpacecraftsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetSpacecraftRequest(_ context.Context, req *spacecraftv1.GetSpacecraftRequest) (*GetSpacecraftRequest, error) {
	return &GetSpacecraftRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetSpacecraftResponse(ctx context.Context, resp *Spacecraft) (*spacecraftv1.GetSpacecraftResponse, error) {
	return &spacecraftv1.GetSpacecraftResponse{
		Spacecraft: mapBizToProtoSpacecraft(resp),
	}, nil
}

func mapBizToProtoAgencyDetailed(r *AgencyDetailed) *spacecraftv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &spacecraftv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*spacecraftv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*spacecraftv1.Country, len(r.Country))
			for i := range r.Country {
				res[i] = mapBizToProtoCountry(&r.Country[i])
			}
			return res
		}(),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLandingsPayload: r.FailedLandingsPayload,
		FailedLandingsSpacecraft: r.FailedLandingsSpacecraft,
		FailedLaunches: r.FailedLaunches,
		Featured: r.Featured,
		FoundingYear: r.FoundingYear,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapBizToProtoImage(r.SocialLogo),
		SocialMediaLinks: func() []*spacecraftv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*spacecraftv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i := range r.SocialMediaLinks {
				res[i] = mapBizToProtoSocialMediaLink(&r.SocialMediaLinks[i])
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoAgencyMini(r *AgencyMini) *spacecraftv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &spacecraftv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *spacecraftv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*spacecraftv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*spacecraftv1.Country, len(r.Country))
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

func mapBizToProtoAgencyType(r *AgencyType) *spacecraftv1.AgencyType {
	if r == nil {
		return nil
	}
	return &spacecraftv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCelestialBodyDetailed(r *CelestialBodyDetailed) *spacecraftv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &spacecraftv1.CelestialBodyDetailed{
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

func mapBizToProtoCelestialBodyMini(r *CelestialBodyMini) *spacecraftv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &spacecraftv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoCelestialBodyNormal(r *CelestialBodyNormal) *spacecraftv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.CelestialBodyNormal{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoCelestialBodyType(r.TypeVal),
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoCelestialBodyType(r *CelestialBodyType) *spacecraftv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &spacecraftv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *spacecraftv1.Country {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoImage(r *Image) *spacecraftv1.Image {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*spacecraftv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*spacecraftv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *spacecraftv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &spacecraftv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *spacecraftv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &spacecraftv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *spacecraftv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &spacecraftv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoInfoURL(r *InfoURL) *spacecraftv1.InfoURL {
	if r == nil {
		return nil
	}
	return &spacecraftv1.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapBizToProtoLanguage(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		Type: mapBizToProtoInfoURLType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoInfoURLType(r *InfoURLType) *spacecraftv1.InfoURLType {
	if r == nil {
		return nil
	}
	return &spacecraftv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanding(r *Landing) *spacecraftv1.Landing {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Id: r.Id,
		LandingLocation: mapBizToProtoLandingLocation(r.LandingLocation),
		Success: r.Success,
		Type: mapBizToProtoLandingType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoLandingLocation(r *LandingLocation) *spacecraftv1.LandingLocation {
	if r == nil {
		return nil
	}
	return &spacecraftv1.LandingLocation{
		Abbrev: r.Abbrev,
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		CelestialBody: mapBizToProtoCelestialBodyNormal(r.CelestialBody),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Latitude: r.Latitude,
		Location: mapBizToProtoLocationSerializerNoCelestialBody(r.Location),
		Longitude: r.Longitude,
		Name: r.Name,
		SuccessfulLandings: r.SuccessfulLandings,
	}
}

func mapBizToProtoLandingType(r *LandingType) *spacecraftv1.LandingType {
	if r == nil {
		return nil
	}
	return &spacecraftv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanguage(r *Language) *spacecraftv1.Language {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLaunchNormal(r *LaunchNormal) *spacecraftv1.LaunchNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.LaunchNormal{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapBizToProtoAgencyMini(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapBizToProtoMission(r.Mission),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapBizToProtoNetPrecision(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapBizToProtoPad(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		Probability: r.Probability,
		Program: func() []*spacecraftv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*spacecraftv1.ProgramNormal, len(r.Program))
			for i := range r.Program {
				res[i] = mapBizToProtoProgramNormal(&r.Program[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapBizToProtoRocketNormal(r.Rocket),
		Slug: r.Slug,
		Status: mapBizToProtoLaunchStatus(r.Status),
		Url: r.Url,
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
}

func mapBizToProtoLaunchStatus(r *LaunchStatus) *spacecraftv1.LaunchStatus {
	if r == nil {
		return nil
	}
	return &spacecraftv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLauncherConfigFamilyMini(r *LauncherConfigFamilyMini) *spacecraftv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &spacecraftv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigList(r *LauncherConfigList) *spacecraftv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &spacecraftv1.LauncherConfigList{
		Families: func() []*spacecraftv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*spacecraftv1.LauncherConfigFamilyMini, len(r.Families))
			for i := range r.Families {
				res[i] = mapBizToProtoLauncherConfigFamilyMini(&r.Families[i])
			}
			return res
		}(),
		FullName: r.FullName,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		Variant: r.Variant,
	}
}

func mapBizToProtoLocation(r *Location) *spacecraftv1.Location {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Location{
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

func mapBizToProtoLocationSerializerNoCelestialBody(r *LocationSerializerNoCelestialBody) *spacecraftv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &spacecraftv1.LocationSerializerNoCelestialBody{
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

func mapBizToProtoMission(r *Mission) *spacecraftv1.Mission {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Mission{
		Agencies: func() []*spacecraftv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacecraftv1.AgencyDetailed, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyDetailed(&r.Agencies[i])
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrls: func() []*spacecraftv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*spacecraftv1.InfoURL, len(r.InfoUrls))
			for i := range r.InfoUrls {
				res[i] = mapBizToProtoInfoURL(&r.InfoUrls[i])
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapBizToProtoOrbit(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*spacecraftv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*spacecraftv1.VidURL, len(r.VidUrls))
			for i := range r.VidUrls {
				res[i] = mapBizToProtoVidURL(&r.VidUrls[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *spacecraftv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &spacecraftv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoNetPrecision(r *NetPrecision) *spacecraftv1.NetPrecision {
	if r == nil {
		return nil
	}
	return &spacecraftv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoOrbit(r *Orbit) *spacecraftv1.Orbit {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapBizToProtoCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoPad(r *Pad) *spacecraftv1.Pad {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Pad{
		Active: r.Active,
		Agencies: func() []*spacecraftv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacecraftv1.AgencyNormal, len(r.Agencies))
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

func mapBizToProtoProgramNormal(r *ProgramNormal) *spacecraftv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.ProgramNormal{
		Agencies: func() []*spacecraftv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacecraftv1.AgencyMini, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyMini(&r.Agencies[i])
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []*spacecraftv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*spacecraftv1.MissionPatch, len(r.MissionPatches))
			for i := range r.MissionPatches {
				res[i] = mapBizToProtoMissionPatch(&r.MissionPatches[i])
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		Type: mapBizToProtoProgramType(r.TypeVal),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoProgramType(r *ProgramType) *spacecraftv1.ProgramType {
	if r == nil {
		return nil
	}
	return &spacecraftv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoRocketNormal(r *RocketNormal) *spacecraftv1.RocketNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.RocketNormal{
		Configuration: mapBizToProtoLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *spacecraftv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *spacecraftv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftConfigDetailed(r *SpacecraftConfigDetailed) *spacecraftv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftConfigDetailed{
		Agency: mapBizToProtoAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*spacecraftv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*spacecraftv1.SpacecraftConfigFamilyDetailed, len(r.Family))
			for i := range r.Family {
				res[i] = mapBizToProtoSpacecraftConfigFamilyDetailed(&r.Family[i])
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FlightLife: r.FlightLife,
		Height: r.Height,
		History: r.History,
		HumanRated: r.HumanRated,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InUse: r.InUse,
		InfoLink: r.InfoLink,
		MaidenFlight: r.MaidenFlight,
		Name: r.Name,
		PayloadCapacity: r.PayloadCapacity,
		PayloadReturnCapacity: r.PayloadReturnCapacity,
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		Type: mapBizToProtoSpacecraftConfigType(r.TypeVal),
		Url: r.Url,
		WikiLink: r.WikiLink,
	}
}

func mapBizToProtoSpacecraftConfigFamilyDetailed(r *SpacecraftConfigFamilyDetailed) *spacecraftv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftConfigFamilyDetailed{
		AttemptedLandings: r.AttemptedLandings,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapBizToProtoAgencyNormal(r.Manufacturer),
		Name: r.Name,
		Parent: mapBizToProtoSpacecraftConfigFamilyNormal(r.Parent),
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
}

func mapBizToProtoSpacecraftConfigFamilyMini(r *SpacecraftConfigFamilyMini) *spacecraftv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigFamilyNormal(r *SpacecraftConfigFamilyNormal) *spacecraftv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapBizToProtoSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigNormal(r *SpacecraftConfigNormal) *spacecraftv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftConfigNormal{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Family: func() []*spacecraftv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*spacecraftv1.SpacecraftConfigFamilyNormal, len(r.Family))
			for i := range r.Family {
				res[i] = mapBizToProtoSpacecraftConfigFamilyNormal(&r.Family[i])
			}
			return res
		}(),
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InUse: r.InUse,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoSpacecraftConfigType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftConfigType(r *SpacecraftConfigType) *spacecraftv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraft(r *Spacecraft) *spacecraftv1.Spacecraft {
	if r == nil {
		return nil
	}
	return &spacecraftv1.Spacecraft{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Flights: func() []*spacecraftv1.SpacecraftFlightNormal {
			if r.Flights == nil {
				return nil
			}
			res := make([]*spacecraftv1.SpacecraftFlightNormal, len(r.Flights))
			for i := range r.Flights {
				res[i] = mapBizToProtoSpacecraftFlightNormal(&r.Flights[i])
			}
			return res
		}(),
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapBizToProtoSpacecraftConfigDetailed(r.SpacecraftConfig),
		Status: mapBizToProtoSpacecraftStatus(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftFlightNormal(r *SpacecraftFlightNormal) *spacecraftv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapBizToProtoLanding(r.Landing),
		Launch: mapBizToProtoLaunchNormal(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapBizToProtoSpacecraftNormal(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftNormal(r *SpacecraftNormal) *spacecraftv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftNormal{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapBizToProtoSpacecraftConfigNormal(r.SpacecraftConfig),
		Status: mapBizToProtoSpacecraftStatus(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftStatus(r *SpacecraftStatus) *spacecraftv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &spacecraftv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoVidURL(r *VidURL) *spacecraftv1.VidURL {
	if r == nil {
		return nil
	}
	return &spacecraftv1.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapBizToProtoLanguage(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		Type: mapBizToProtoVidURLType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoVidURLType(r *VidURLType) *spacecraftv1.VidURLType {
	if r == nil {
		return nil
	}
	return &spacecraftv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

