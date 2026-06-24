package spacewalk

import (
	"context"

	"github.com/pobochiigo/bhole/internal/transport"
	spacewalkv1 "github.com/pobochiigo/bhole/proto/spacewalk/v1"
	v1connect "github.com/pobochiigo/bhole/proto/spacewalk/v1/spacewalkv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListSpacewalks transport.Handler[spacewalkv1.ListSpacewalksRequest, spacewalkv1.ListSpacewalksResponse]
	getSpacewalk    transport.Handler[spacewalkv1.GetSpacewalkRequest, spacewalkv1.GetSpacewalkResponse]
}

func (s *server) ListSpacewalks(ctx context.Context, req *connect.Request[spacewalkv1.ListSpacewalksRequest]) (*connect.Response[spacewalkv1.ListSpacewalksResponse], error) {
	return s.listListSpacewalks(ctx, req)
}

func (s *server) GetSpacewalk(ctx context.Context, req *connect.Request[spacewalkv1.GetSpacewalkRequest]) (*connect.Response[spacewalkv1.GetSpacewalkResponse], error) {
	return s.getSpacewalk(ctx, req)
}

func NewSpacewalkHandler(svc Service) v1connect.SpacewalkServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListSpacewalks: transport.NewConnectServer(
			eps.listListSpacewalks,
			decodeListSpacewalksRequest,
			encodeListSpacewalksResponse,
		),
		getSpacewalk: transport.NewConnectServer(
			eps.getSpacewalk,
			decodeGetSpacewalkRequest,
			encodeGetSpacewalkResponse,
		),
	}
}

func decodeListSpacewalksRequest(_ context.Context, req *spacewalkv1.ListSpacewalksRequest) (*ListSpacewalksRequest, error) {
	return &ListSpacewalksRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListSpacewalksResponse(ctx context.Context, resp *ListSpacewalksResponse) (*spacewalkv1.ListSpacewalksResponse, error) {
	results := make([]*spacewalkv1.Spacewalk, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoSpacewalk(&resp.Results[i])
	}
	return &spacewalkv1.ListSpacewalksResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetSpacewalkRequest(_ context.Context, req *spacewalkv1.GetSpacewalkRequest) (*GetSpacewalkRequest, error) {
	return &GetSpacewalkRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetSpacewalkResponse(ctx context.Context, resp *Spacewalk) (*spacewalkv1.GetSpacewalkResponse, error) {
	return &spacewalkv1.GetSpacewalkResponse{
		Spacewalk: mapBizToProtoSpacewalk(resp),
	}, nil
}

func mapBizToProtoAgencyDetailed(r *AgencyDetailed) *spacewalkv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*spacewalkv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*spacewalkv1.Country, len(r.Country))
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
		SocialMediaLinks: func() []*spacewalkv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*spacewalkv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAgencyMini(r *AgencyMini) *spacewalkv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *spacewalkv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*spacewalkv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*spacewalkv1.Country, len(r.Country))
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

func mapBizToProtoAgencyType(r *AgencyType) *spacewalkv1.AgencyType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronautDetailed(r *AstronautDetailed) *spacewalkv1.AstronautDetailed {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AstronautDetailed{
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
		Nationality: func() []*spacewalkv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*spacewalkv1.Country, len(r.Nationality))
			for i := range r.Nationality {
				res[i] = mapBizToProtoCountry(&r.Nationality[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*spacewalkv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*spacewalkv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAstronautFlight(r *AstronautFlight) *spacewalkv1.AstronautFlight {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AstronautFlight{
		Astronaut: mapBizToProtoAstronautDetailed(r.Astronaut),
		Id: r.Id,
		Role: mapBizToProtoAstronautRole(r.Role),
	}
}

func mapBizToProtoAstronautRole(r *AstronautRole) *spacewalkv1.AstronautRole {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
}

func mapBizToProtoAstronautStatus(r *AstronautStatus) *spacewalkv1.AstronautStatus {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronautType(r *AstronautType) *spacewalkv1.AstronautType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCelestialBodyDetailed(r *CelestialBodyDetailed) *spacewalkv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &spacewalkv1.CelestialBodyDetailed{
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

func mapBizToProtoCelestialBodyMini(r *CelestialBodyMini) *spacewalkv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &spacewalkv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoCelestialBodyNormal(r *CelestialBodyNormal) *spacewalkv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.CelestialBodyNormal{
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

func mapBizToProtoCelestialBodyType(r *CelestialBodyType) *spacewalkv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *spacewalkv1.Country {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoDockingEventForChaserNormal(r *DockingEventForChaserNormal) *spacewalkv1.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.DockingEventForChaserNormal{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapBizToProtoDockingLocation(r.DockingLocation),
		FlightVehicleTarget: mapBizToProtoSpacecraftFlightNormal(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightTarget: mapBizToProtoPayloadFlightNormal(r.PayloadFlightTarget),
		SpaceStationTarget: mapBizToProtoSpaceStationNormal(r.SpaceStationTarget),
		Url: r.Url,
	}
}

func mapBizToProtoDockingLocation(r *DockingLocation) *spacewalkv1.DockingLocation {
	if r == nil {
		return nil
	}
	return &spacewalkv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapBizToProtoPayloadMini(r.Payload),
		Spacecraft: mapBizToProtoSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapBizToProtoSpaceStationMini(r.Spacestation),
	}
}

func mapBizToProtoEventNormal(r *EventNormal) *spacewalkv1.EventNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.EventNormal{
		Date: r.Date,
		DatePrecision: mapBizToProtoNetPrecision(r.DatePrecision),
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrls: func() []*spacewalkv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*spacewalkv1.InfoURL, len(r.InfoUrls))
			for i := range r.InfoUrls {
				res[i] = mapBizToProtoInfoURL(&r.InfoUrls[i])
			}
			return res
		}(),
		Location: r.Location,
		Name: r.Name,
		Slug: r.Slug,
		Type: mapBizToProtoEventType(r.TypeVal),
		Url: r.Url,
		VidUrls: func() []*spacewalkv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*spacewalkv1.VidURL, len(r.VidUrls))
			for i := range r.VidUrls {
				res[i] = mapBizToProtoVidURL(&r.VidUrls[i])
			}
			return res
		}(),
		WebcastLive: r.WebcastLive,
	}
}

func mapBizToProtoEventType(r *EventType) *spacewalkv1.EventType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.EventType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoExpeditionNormalSerializerForSpacewalk(r *ExpeditionNormalSerializerForSpacewalk) *spacewalkv1.ExpeditionNormalSerializerForSpacewalk {
	if r == nil {
		return nil
	}
	return &spacewalkv1.ExpeditionNormalSerializerForSpacewalk{
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []*spacewalkv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*spacewalkv1.MissionPatch, len(r.MissionPatches))
			for i := range r.MissionPatches {
				res[i] = mapBizToProtoMissionPatch(&r.MissionPatches[i])
			}
			return res
		}(),
		Name: r.Name,
		Spacestation: mapBizToProtoSpaceStationNormal(r.Spacestation),
		Start: r.Start,
		Url: r.Url,
	}
}

func mapBizToProtoImage(r *Image) *spacewalkv1.Image {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*spacewalkv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*spacewalkv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *spacewalkv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &spacewalkv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *spacewalkv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &spacewalkv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *spacewalkv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoInfoURL(r *InfoURL) *spacewalkv1.InfoURL {
	if r == nil {
		return nil
	}
	return &spacewalkv1.InfoURL{
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

func mapBizToProtoInfoURLType(r *InfoURLType) *spacewalkv1.InfoURLType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanding(r *Landing) *spacewalkv1.Landing {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Landing{
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

func mapBizToProtoLandingLocation(r *LandingLocation) *spacewalkv1.LandingLocation {
	if r == nil {
		return nil
	}
	return &spacewalkv1.LandingLocation{
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

func mapBizToProtoLandingType(r *LandingType) *spacewalkv1.LandingType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanguage(r *Language) *spacewalkv1.Language {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLaunchNormal(r *LaunchNormal) *spacewalkv1.LaunchNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.LaunchNormal{
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
		Program: func() []*spacewalkv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*spacewalkv1.ProgramNormal, len(r.Program))
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

func mapBizToProtoLaunchStatus(r *LaunchStatus) *spacewalkv1.LaunchStatus {
	if r == nil {
		return nil
	}
	return &spacewalkv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLauncherConfigFamilyMini(r *LauncherConfigFamilyMini) *spacewalkv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &spacewalkv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigList(r *LauncherConfigList) *spacewalkv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &spacewalkv1.LauncherConfigList{
		Families: func() []*spacewalkv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*spacewalkv1.LauncherConfigFamilyMini, len(r.Families))
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

func mapBizToProtoLocation(r *Location) *spacewalkv1.Location {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Location{
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

func mapBizToProtoLocationSerializerNoCelestialBody(r *LocationSerializerNoCelestialBody) *spacewalkv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &spacewalkv1.LocationSerializerNoCelestialBody{
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

func mapBizToProtoMission(r *Mission) *spacewalkv1.Mission {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Mission{
		Agencies: func() []*spacewalkv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacewalkv1.AgencyDetailed, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyDetailed(&r.Agencies[i])
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrls: func() []*spacewalkv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*spacewalkv1.InfoURL, len(r.InfoUrls))
			for i := range r.InfoUrls {
				res[i] = mapBizToProtoInfoURL(&r.InfoUrls[i])
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapBizToProtoOrbit(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*spacewalkv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*spacewalkv1.VidURL, len(r.VidUrls))
			for i := range r.VidUrls {
				res[i] = mapBizToProtoVidURL(&r.VidUrls[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *spacewalkv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &spacewalkv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoNetPrecision(r *NetPrecision) *spacewalkv1.NetPrecision {
	if r == nil {
		return nil
	}
	return &spacewalkv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoOrbit(r *Orbit) *spacewalkv1.Orbit {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapBizToProtoCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoPad(r *Pad) *spacewalkv1.Pad {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Pad{
		Active: r.Active,
		Agencies: func() []*spacewalkv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacewalkv1.AgencyNormal, len(r.Agencies))
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

func mapBizToProtoPayloadFlightNormal(r *PayloadFlightNormal) *spacewalkv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.PayloadFlightNormal{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapBizToProtoLanding(r.Landing),
		Launch: mapBizToProtoLaunchNormal(r.Launch),
		Payload: mapBizToProtoPayloadNormal(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
}

func mapBizToProtoPayloadMini(r *PayloadMini) *spacewalkv1.PayloadMini {
	if r == nil {
		return nil
	}
	return &spacewalkv1.PayloadMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapBizToProtoAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoPayloadType(r.TypeVal),
	}
}

func mapBizToProtoPayloadNormal(r *PayloadNormal) *spacewalkv1.PayloadNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapBizToProtoAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapBizToProtoAgencyNormal(r.Operator),
		Program: func() []*spacewalkv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*spacewalkv1.ProgramMini, len(r.Program))
			for i := range r.Program {
				res[i] = mapBizToProtoProgramMini(&r.Program[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoPayloadType(r.TypeVal),
		WikiLink: r.WikiLink,
	}
}

func mapBizToProtoPayloadType(r *PayloadType) *spacewalkv1.PayloadType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoProgramMini(r *ProgramMini) *spacewalkv1.ProgramMini {
	if r == nil {
		return nil
	}
	return &spacewalkv1.ProgramMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoProgramNormal(r *ProgramNormal) *spacewalkv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.ProgramNormal{
		Agencies: func() []*spacewalkv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*spacewalkv1.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []*spacewalkv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*spacewalkv1.MissionPatch, len(r.MissionPatches))
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

func mapBizToProtoProgramType(r *ProgramType) *spacewalkv1.ProgramType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoRocketNormal(r *RocketNormal) *spacewalkv1.RocketNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.RocketNormal{
		Configuration: mapBizToProtoLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *spacewalkv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *spacewalkv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationMini(r *SpaceStationMini) *spacewalkv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpaceStationMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationNormal(r *SpaceStationNormal) *spacewalkv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpaceStationNormal{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Status: mapBizToProtoSpaceStationStatus(r.Status),
		Type: mapBizToProtoSpaceStationType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationStatus(r *SpaceStationStatus) *spacewalkv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpaceStationType(r *SpaceStationType) *spacewalkv1.SpaceStationType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraftConfigDetailed(r *SpacecraftConfigDetailed) *spacewalkv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftConfigDetailed{
		Agency: mapBizToProtoAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*spacewalkv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*spacewalkv1.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapBizToProtoSpacecraftConfigFamilyDetailed(r *SpacecraftConfigFamilyDetailed) *spacewalkv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftConfigFamilyDetailed{
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

func mapBizToProtoSpacecraftConfigFamilyMini(r *SpacecraftConfigFamilyMini) *spacewalkv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigFamilyNormal(r *SpacecraftConfigFamilyNormal) *spacewalkv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapBizToProtoSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigNormal(r *SpacecraftConfigNormal) *spacewalkv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftConfigNormal{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Family: func() []*spacewalkv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*spacewalkv1.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapBizToProtoSpacecraftConfigType(r *SpacecraftConfigType) *spacewalkv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraftDetailed(r *SpacecraftDetailed) *spacewalkv1.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftDetailed{
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
		SpacecraftConfig: mapBizToProtoSpacecraftConfigDetailed(r.SpacecraftConfig),
		Status: mapBizToProtoSpacecraftStatus(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftFlightDetailed(r *SpacecraftFlightDetailed) *spacewalkv1.SpacecraftFlightDetailed {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftFlightDetailed{
		Destination: r.Destination,
		DockingEvents: func() []*spacewalkv1.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]*spacewalkv1.DockingEventForChaserNormal, len(r.DockingEvents))
			for i := range r.DockingEvents {
				res[i] = mapBizToProtoDockingEventForChaserNormal(&r.DockingEvents[i])
			}
			return res
		}(),
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapBizToProtoLanding(r.Landing),
		LandingCrew: func() []*spacewalkv1.AstronautFlight {
			if r.LandingCrew == nil {
				return nil
			}
			res := make([]*spacewalkv1.AstronautFlight, len(r.LandingCrew))
			for i := range r.LandingCrew {
				res[i] = mapBizToProtoAstronautFlight(&r.LandingCrew[i])
			}
			return res
		}(),
		Launch: mapBizToProtoLaunchNormal(r.Launch),
		LaunchCrew: func() []*spacewalkv1.AstronautFlight {
			if r.LaunchCrew == nil {
				return nil
			}
			res := make([]*spacewalkv1.AstronautFlight, len(r.LaunchCrew))
			for i := range r.LaunchCrew {
				res[i] = mapBizToProtoAstronautFlight(&r.LaunchCrew[i])
			}
			return res
		}(),
		MissionEnd: r.MissionEnd,
		OnboardCrew: func() []*spacewalkv1.AstronautFlight {
			if r.OnboardCrew == nil {
				return nil
			}
			res := make([]*spacewalkv1.AstronautFlight, len(r.OnboardCrew))
			for i := range r.OnboardCrew {
				res[i] = mapBizToProtoAstronautFlight(&r.OnboardCrew[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Spacecraft: mapBizToProtoSpacecraftDetailed(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftFlightNormal(r *SpacecraftFlightNormal) *spacewalkv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftFlightNormal{
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

func mapBizToProtoSpacecraftNormal(r *SpacecraftNormal) *spacewalkv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftNormal{
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

func mapBizToProtoSpacecraftStatus(r *SpacecraftStatus) *spacewalkv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &spacewalkv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacewalk(r *Spacewalk) *spacewalkv1.Spacewalk {
	if r == nil {
		return nil
	}
	return &spacewalkv1.Spacewalk{
		Crew: func() []*spacewalkv1.AstronautFlight {
			if r.Crew == nil {
				return nil
			}
			res := make([]*spacewalkv1.AstronautFlight, len(r.Crew))
			for i := range r.Crew {
				res[i] = mapBizToProtoAstronautFlight(&r.Crew[i])
			}
			return res
		}(),
		Duration: r.Duration,
		End: r.End,
		Event: mapBizToProtoEventNormal(r.Event),
		Expedition: mapBizToProtoExpeditionNormalSerializerForSpacewalk(r.Expedition),
		Id: r.Id,
		Location: r.Location,
		Name: r.Name,
		Program: func() []*spacewalkv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*spacewalkv1.ProgramNormal, len(r.Program))
			for i := range r.Program {
				res[i] = mapBizToProtoProgramNormal(&r.Program[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SpacecraftFlight: mapBizToProtoSpacecraftFlightDetailed(r.SpacecraftFlight),
		Spacestation: mapBizToProtoSpaceStationNormal(r.Spacestation),
		Start: r.Start,
		Url: r.Url,
	}
}

func mapBizToProtoVidURL(r *VidURL) *spacewalkv1.VidURL {
	if r == nil {
		return nil
	}
	return &spacewalkv1.VidURL{
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

func mapBizToProtoVidURLType(r *VidURLType) *spacewalkv1.VidURLType {
	if r == nil {
		return nil
	}
	return &spacewalkv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

