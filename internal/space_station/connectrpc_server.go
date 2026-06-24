package space_station

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	space_stationv1 "com.gitlab/pobochiigo/bhole/proto/space_station/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/space_station/v1/space_stationv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListSpaceStations transport.Handler[space_stationv1.ListSpaceStationsRequest, space_stationv1.ListSpaceStationsResponse]
	getSpaceStation    transport.Handler[space_stationv1.GetSpaceStationRequest, space_stationv1.GetSpaceStationResponse]
}

func (s *server) ListSpaceStations(ctx context.Context, req *connect.Request[space_stationv1.ListSpaceStationsRequest]) (*connect.Response[space_stationv1.ListSpaceStationsResponse], error) {
	return s.listListSpaceStations(ctx, req)
}

func (s *server) GetSpaceStation(ctx context.Context, req *connect.Request[space_stationv1.GetSpaceStationRequest]) (*connect.Response[space_stationv1.GetSpaceStationResponse], error) {
	return s.getSpaceStation(ctx, req)
}

func NewSpaceStationHandler(svc Service) v1connect.SpaceStationServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListSpaceStations: transport.NewConnectServer(
			eps.listListSpaceStations,
			decodeListSpaceStationsRequest,
			encodeListSpaceStationsResponse,
		),
		getSpaceStation: transport.NewConnectServer(
			eps.getSpaceStation,
			decodeGetSpaceStationRequest,
			encodeGetSpaceStationResponse,
		),
	}
}

func decodeListSpaceStationsRequest(_ context.Context, req *space_stationv1.ListSpaceStationsRequest) (*ListSpaceStationsRequest, error) {
	return &ListSpaceStationsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListSpaceStationsResponse(ctx context.Context, resp *ListSpaceStationsResponse) (*space_stationv1.ListSpaceStationsResponse, error) {
	results := make([]*space_stationv1.SpaceStation, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoSpaceStation(&resp.Results[i])
	}
	return &space_stationv1.ListSpaceStationsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetSpaceStationRequest(_ context.Context, req *space_stationv1.GetSpaceStationRequest) (*GetSpaceStationRequest, error) {
	return &GetSpaceStationRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetSpaceStationResponse(ctx context.Context, resp *SpaceStation) (*space_stationv1.GetSpaceStationResponse, error) {
	return &space_stationv1.GetSpaceStationResponse{
		SpaceStation: mapBizToProtoSpaceStation(resp),
	}, nil
}

func mapBizToProtoAgencyDetailed(r *AgencyDetailed) *space_stationv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &space_stationv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*space_stationv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*space_stationv1.Country, len(r.Country))
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
		SocialMediaLinks: func() []*space_stationv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*space_stationv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAgencyMini(r *AgencyMini) *space_stationv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &space_stationv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *space_stationv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*space_stationv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*space_stationv1.Country, len(r.Country))
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

func mapBizToProtoAgencyType(r *AgencyType) *space_stationv1.AgencyType {
	if r == nil {
		return nil
	}
	return &space_stationv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCelestialBodyDetailed(r *CelestialBodyDetailed) *space_stationv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &space_stationv1.CelestialBodyDetailed{
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

func mapBizToProtoCelestialBodyMini(r *CelestialBodyMini) *space_stationv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &space_stationv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoCelestialBodyNormal(r *CelestialBodyNormal) *space_stationv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.CelestialBodyNormal{
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

func mapBizToProtoCelestialBodyType(r *CelestialBodyType) *space_stationv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &space_stationv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *space_stationv1.Country {
	if r == nil {
		return nil
	}
	return &space_stationv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoDockingEventDetailedSerializerForSpacestation(r *DockingEventDetailedSerializerForSpacestation) *space_stationv1.DockingEventDetailedSerializerForSpacestation {
	if r == nil {
		return nil
	}
	return &space_stationv1.DockingEventDetailedSerializerForSpacestation{
		Departure: r.Departure,
		Docking: r.Docking,
		FlightVehicleChaser: mapBizToProtoSpacecraftFlightForDockingEvent(r.FlightVehicleChaser),
		Id: r.Id,
		PayloadFlightChaser: mapBizToProtoPayloadFlightNormal(r.PayloadFlightChaser),
		SpaceStationChaser: mapBizToProtoSpaceStationNormal(r.SpaceStationChaser),
		Url: r.Url,
	}
}

func mapBizToProtoDockingEventForChaserNormal(r *DockingEventForChaserNormal) *space_stationv1.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.DockingEventForChaserNormal{
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

func mapBizToProtoDockingLocation(r *DockingLocation) *space_stationv1.DockingLocation {
	if r == nil {
		return nil
	}
	return &space_stationv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapBizToProtoPayloadMini(r.Payload),
		Spacecraft: mapBizToProtoSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapBizToProtoSpaceStationMini(r.Spacestation),
	}
}

func mapBizToProtoDockingLocationSerializerForSpacestation(r *DockingLocationSerializerForSpacestation) *space_stationv1.DockingLocationSerializerForSpacestation {
	if r == nil {
		return nil
	}
	return &space_stationv1.DockingLocationSerializerForSpacestation{
		CurrentlyDocked: mapBizToProtoDockingEventDetailedSerializerForSpacestation(r.CurrentlyDocked),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoExpeditionMini(r *ExpeditionMini) *space_stationv1.ExpeditionMini {
	if r == nil {
		return nil
	}
	return &space_stationv1.ExpeditionMini{
		End: r.End,
		Id: r.Id,
		Name: r.Name,
		Start: r.Start,
		Url: r.Url,
	}
}

func mapBizToProtoImage(r *Image) *space_stationv1.Image {
	if r == nil {
		return nil
	}
	return &space_stationv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*space_stationv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*space_stationv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *space_stationv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &space_stationv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *space_stationv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &space_stationv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *space_stationv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &space_stationv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoInfoURL(r *InfoURL) *space_stationv1.InfoURL {
	if r == nil {
		return nil
	}
	return &space_stationv1.InfoURL{
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

func mapBizToProtoInfoURLType(r *InfoURLType) *space_stationv1.InfoURLType {
	if r == nil {
		return nil
	}
	return &space_stationv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanding(r *Landing) *space_stationv1.Landing {
	if r == nil {
		return nil
	}
	return &space_stationv1.Landing{
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

func mapBizToProtoLandingLocation(r *LandingLocation) *space_stationv1.LandingLocation {
	if r == nil {
		return nil
	}
	return &space_stationv1.LandingLocation{
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

func mapBizToProtoLandingType(r *LandingType) *space_stationv1.LandingType {
	if r == nil {
		return nil
	}
	return &space_stationv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanguage(r *Language) *space_stationv1.Language {
	if r == nil {
		return nil
	}
	return &space_stationv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLaunchNormal(r *LaunchNormal) *space_stationv1.LaunchNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.LaunchNormal{
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
		Program: func() []*space_stationv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*space_stationv1.ProgramNormal, len(r.Program))
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

func mapBizToProtoLaunchStatus(r *LaunchStatus) *space_stationv1.LaunchStatus {
	if r == nil {
		return nil
	}
	return &space_stationv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLauncherConfigFamilyMini(r *LauncherConfigFamilyMini) *space_stationv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &space_stationv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigList(r *LauncherConfigList) *space_stationv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &space_stationv1.LauncherConfigList{
		Families: func() []*space_stationv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*space_stationv1.LauncherConfigFamilyMini, len(r.Families))
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

func mapBizToProtoLocation(r *Location) *space_stationv1.Location {
	if r == nil {
		return nil
	}
	return &space_stationv1.Location{
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

func mapBizToProtoLocationSerializerNoCelestialBody(r *LocationSerializerNoCelestialBody) *space_stationv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &space_stationv1.LocationSerializerNoCelestialBody{
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

func mapBizToProtoMission(r *Mission) *space_stationv1.Mission {
	if r == nil {
		return nil
	}
	return &space_stationv1.Mission{
		Agencies: func() []*space_stationv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*space_stationv1.AgencyDetailed, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyDetailed(&r.Agencies[i])
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrls: func() []*space_stationv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*space_stationv1.InfoURL, len(r.InfoUrls))
			for i := range r.InfoUrls {
				res[i] = mapBizToProtoInfoURL(&r.InfoUrls[i])
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapBizToProtoOrbit(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*space_stationv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*space_stationv1.VidURL, len(r.VidUrls))
			for i := range r.VidUrls {
				res[i] = mapBizToProtoVidURL(&r.VidUrls[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *space_stationv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &space_stationv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoNetPrecision(r *NetPrecision) *space_stationv1.NetPrecision {
	if r == nil {
		return nil
	}
	return &space_stationv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoOrbit(r *Orbit) *space_stationv1.Orbit {
	if r == nil {
		return nil
	}
	return &space_stationv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapBizToProtoCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoPad(r *Pad) *space_stationv1.Pad {
	if r == nil {
		return nil
	}
	return &space_stationv1.Pad{
		Active: r.Active,
		Agencies: func() []*space_stationv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*space_stationv1.AgencyNormal, len(r.Agencies))
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

func mapBizToProtoPayloadFlightNormal(r *PayloadFlightNormal) *space_stationv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.PayloadFlightNormal{
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

func mapBizToProtoPayloadMini(r *PayloadMini) *space_stationv1.PayloadMini {
	if r == nil {
		return nil
	}
	return &space_stationv1.PayloadMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapBizToProtoAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoPayloadType(r.TypeVal),
	}
}

func mapBizToProtoPayloadNormal(r *PayloadNormal) *space_stationv1.PayloadNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapBizToProtoAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapBizToProtoAgencyNormal(r.Operator),
		Program: func() []*space_stationv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*space_stationv1.ProgramMini, len(r.Program))
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

func mapBizToProtoPayloadType(r *PayloadType) *space_stationv1.PayloadType {
	if r == nil {
		return nil
	}
	return &space_stationv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoProgramMini(r *ProgramMini) *space_stationv1.ProgramMini {
	if r == nil {
		return nil
	}
	return &space_stationv1.ProgramMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoProgramNormal(r *ProgramNormal) *space_stationv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.ProgramNormal{
		Agencies: func() []*space_stationv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*space_stationv1.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []*space_stationv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*space_stationv1.MissionPatch, len(r.MissionPatches))
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

func mapBizToProtoProgramType(r *ProgramType) *space_stationv1.ProgramType {
	if r == nil {
		return nil
	}
	return &space_stationv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoRocketNormal(r *RocketNormal) *space_stationv1.RocketNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.RocketNormal{
		Configuration: mapBizToProtoLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *space_stationv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &space_stationv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *space_stationv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &space_stationv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStation(r *SpaceStation) *space_stationv1.SpaceStation {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpaceStation{
		ActiveDockingEvents: func() []*space_stationv1.DockingEventForChaserNormal {
			if r.ActiveDockingEvents == nil {
				return nil
			}
			res := make([]*space_stationv1.DockingEventForChaserNormal, len(r.ActiveDockingEvents))
			for i := range r.ActiveDockingEvents {
				res[i] = mapBizToProtoDockingEventForChaserNormal(&r.ActiveDockingEvents[i])
			}
			return res
		}(),
		ActiveExpeditions: func() []*space_stationv1.ExpeditionMini {
			if r.ActiveExpeditions == nil {
				return nil
			}
			res := make([]*space_stationv1.ExpeditionMini, len(r.ActiveExpeditions))
			for i := range r.ActiveExpeditions {
				res[i] = mapBizToProtoExpeditionMini(&r.ActiveExpeditions[i])
			}
			return res
		}(),
		Deorbited: r.Deorbited,
		Description: r.Description,
		DockedVehicles: r.DockedVehicles,
		DockingLocation: func() []*space_stationv1.DockingLocationSerializerForSpacestation {
			if r.DockingLocation == nil {
				return nil
			}
			res := make([]*space_stationv1.DockingLocationSerializerForSpacestation, len(r.DockingLocation))
			for i := range r.DockingLocation {
				res[i] = mapBizToProtoDockingLocationSerializerForSpacestation(&r.DockingLocation[i])
			}
			return res
		}(),
		Founded: r.Founded,
		Height: r.Height,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Mass: r.Mass,
		Name: r.Name,
		OnboardCrew: r.OnboardCrew,
		Orbit: r.Orbit,
		Owners: func() []*space_stationv1.AgencyNormal {
			if r.Owners == nil {
				return nil
			}
			res := make([]*space_stationv1.AgencyNormal, len(r.Owners))
			for i := range r.Owners {
				res[i] = mapBizToProtoAgencyNormal(&r.Owners[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Status: mapBizToProtoSpaceStationStatus(r.Status),
		Type: mapBizToProtoSpaceStationType(r.TypeVal),
		Url: r.Url,
		Volume: r.Volume,
		Width: r.Width,
	}
}

func mapBizToProtoSpaceStationMini(r *SpaceStationMini) *space_stationv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpaceStationMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationNormal(r *SpaceStationNormal) *space_stationv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpaceStationNormal{
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

func mapBizToProtoSpaceStationStatus(r *SpaceStationStatus) *space_stationv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpaceStationType(r *SpaceStationType) *space_stationv1.SpaceStationType {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraftConfigDetailed(r *SpacecraftConfigDetailed) *space_stationv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftConfigDetailed{
		Agency: mapBizToProtoAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*space_stationv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*space_stationv1.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapBizToProtoSpacecraftConfigFamilyDetailed(r *SpacecraftConfigFamilyDetailed) *space_stationv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftConfigFamilyDetailed{
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

func mapBizToProtoSpacecraftConfigFamilyMini(r *SpacecraftConfigFamilyMini) *space_stationv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigFamilyNormal(r *SpacecraftConfigFamilyNormal) *space_stationv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapBizToProtoSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigNormal(r *SpacecraftConfigNormal) *space_stationv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftConfigNormal{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Family: func() []*space_stationv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*space_stationv1.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapBizToProtoSpacecraftConfigType(r *SpacecraftConfigType) *space_stationv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraftDetailed(r *SpacecraftDetailed) *space_stationv1.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftDetailed{
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

func mapBizToProtoSpacecraftFlightForDockingEvent(r *SpacecraftFlightForDockingEvent) *space_stationv1.SpacecraftFlightForDockingEvent {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftFlightForDockingEvent{
		Id: r.Id,
		Launch: mapBizToProtoLaunchNormal(r.Launch),
		Spacecraft: mapBizToProtoSpacecraftDetailed(r.Spacecraft),
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftFlightNormal(r *SpacecraftFlightNormal) *space_stationv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftFlightNormal{
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

func mapBizToProtoSpacecraftNormal(r *SpacecraftNormal) *space_stationv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftNormal{
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

func mapBizToProtoSpacecraftStatus(r *SpacecraftStatus) *space_stationv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &space_stationv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoVidURL(r *VidURL) *space_stationv1.VidURL {
	if r == nil {
		return nil
	}
	return &space_stationv1.VidURL{
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

func mapBizToProtoVidURLType(r *VidURLType) *space_stationv1.VidURLType {
	if r == nil {
		return nil
	}
	return &space_stationv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

