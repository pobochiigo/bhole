package docking_event

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	docking_eventv1 "com.gitlab/pobochiigo/bhole/proto/docking_event/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/docking_event/v1/docking_eventv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListDockingEvents transport.Handler[docking_eventv1.ListDockingEventsRequest, docking_eventv1.ListDockingEventsResponse]
	getDockingEvent    transport.Handler[docking_eventv1.GetDockingEventRequest, docking_eventv1.GetDockingEventResponse]
}

func (s *server) ListDockingEvents(ctx context.Context, req *connect.Request[docking_eventv1.ListDockingEventsRequest]) (*connect.Response[docking_eventv1.ListDockingEventsResponse], error) {
	return s.listListDockingEvents(ctx, req)
}

func (s *server) GetDockingEvent(ctx context.Context, req *connect.Request[docking_eventv1.GetDockingEventRequest]) (*connect.Response[docking_eventv1.GetDockingEventResponse], error) {
	return s.getDockingEvent(ctx, req)
}

func NewDockingEventHandler(svc Service) v1connect.DockingEventServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListDockingEvents: transport.NewConnectServer(
			eps.listListDockingEvents,
			decodeListDockingEventsRequest,
			encodeListDockingEventsResponse,
		),
		getDockingEvent: transport.NewConnectServer(
			eps.getDockingEvent,
			decodeGetDockingEventRequest,
			encodeGetDockingEventResponse,
		),
	}
}

func decodeListDockingEventsRequest(_ context.Context, req *docking_eventv1.ListDockingEventsRequest) (*ListDockingEventsRequest, error) {
	return &ListDockingEventsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListDockingEventsResponse(ctx context.Context, resp *ListDockingEventsResponse) (*docking_eventv1.ListDockingEventsResponse, error) {
	results := make([]*docking_eventv1.DockingEvent, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoDockingEvent(&resp.Results[i])
	}
	return &docking_eventv1.ListDockingEventsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetDockingEventRequest(_ context.Context, req *docking_eventv1.GetDockingEventRequest) (*GetDockingEventRequest, error) {
	return &GetDockingEventRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetDockingEventResponse(ctx context.Context, resp *DockingEvent) (*docking_eventv1.GetDockingEventResponse, error) {
	return &docking_eventv1.GetDockingEventResponse{
		DockingEvent: mapBizToProtoDockingEvent(resp),
	}, nil
}

func mapBizToProtoAgencyDetailed(r *AgencyDetailed) *docking_eventv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &docking_eventv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*docking_eventv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*docking_eventv1.Country, len(r.Country))
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
		SocialMediaLinks: func() []*docking_eventv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*docking_eventv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAgencyMini(r *AgencyMini) *docking_eventv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *docking_eventv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*docking_eventv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*docking_eventv1.Country, len(r.Country))
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

func mapBizToProtoAgencyType(r *AgencyType) *docking_eventv1.AgencyType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCelestialBodyDetailed(r *CelestialBodyDetailed) *docking_eventv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &docking_eventv1.CelestialBodyDetailed{
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

func mapBizToProtoCelestialBodyMini(r *CelestialBodyMini) *docking_eventv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoCelestialBodyNormal(r *CelestialBodyNormal) *docking_eventv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.CelestialBodyNormal{
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

func mapBizToProtoCelestialBodyType(r *CelestialBodyType) *docking_eventv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *docking_eventv1.Country {
	if r == nil {
		return nil
	}
	return &docking_eventv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoDockingEvent(r *DockingEvent) *docking_eventv1.DockingEvent {
	if r == nil {
		return nil
	}
	return &docking_eventv1.DockingEvent{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapBizToProtoDockingLocation(r.DockingLocation),
		FlightVehicleChaser: mapBizToProtoSpacecraftFlightNormal(r.FlightVehicleChaser),
		FlightVehicleTarget: mapBizToProtoSpacecraftFlightMini(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightChaser: mapBizToProtoPayloadFlightNormal(r.PayloadFlightChaser),
		PayloadFlightTarget: mapBizToProtoPayloadFlightMini(r.PayloadFlightTarget),
		ResponseMode: r.ResponseMode,
		SpaceStationChaser: mapBizToProtoSpaceStationNormal(r.SpaceStationChaser),
		SpaceStationTarget: mapBizToProtoSpaceStationMini(r.SpaceStationTarget),
		Url: r.Url,
	}
}

func mapBizToProtoDockingLocation(r *DockingLocation) *docking_eventv1.DockingLocation {
	if r == nil {
		return nil
	}
	return &docking_eventv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapBizToProtoPayloadMini(r.Payload),
		Spacecraft: mapBizToProtoSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapBizToProtoSpaceStationMini(r.Spacestation),
	}
}

func mapBizToProtoImage(r *Image) *docking_eventv1.Image {
	if r == nil {
		return nil
	}
	return &docking_eventv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*docking_eventv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*docking_eventv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *docking_eventv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &docking_eventv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *docking_eventv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &docking_eventv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *docking_eventv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoInfoURL(r *InfoURL) *docking_eventv1.InfoURL {
	if r == nil {
		return nil
	}
	return &docking_eventv1.InfoURL{
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

func mapBizToProtoInfoURLType(r *InfoURLType) *docking_eventv1.InfoURLType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanding(r *Landing) *docking_eventv1.Landing {
	if r == nil {
		return nil
	}
	return &docking_eventv1.Landing{
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

func mapBizToProtoLandingLocation(r *LandingLocation) *docking_eventv1.LandingLocation {
	if r == nil {
		return nil
	}
	return &docking_eventv1.LandingLocation{
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

func mapBizToProtoLandingType(r *LandingType) *docking_eventv1.LandingType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanguage(r *Language) *docking_eventv1.Language {
	if r == nil {
		return nil
	}
	return &docking_eventv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLaunchMini(r *LaunchMini) *docking_eventv1.LaunchMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.LaunchMini{
		Id: r.Id,
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoLaunchNormal(r *LaunchNormal) *docking_eventv1.LaunchNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.LaunchNormal{
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
		Program: func() []*docking_eventv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*docking_eventv1.ProgramNormal, len(r.Program))
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

func mapBizToProtoLaunchStatus(r *LaunchStatus) *docking_eventv1.LaunchStatus {
	if r == nil {
		return nil
	}
	return &docking_eventv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLauncherConfigFamilyMini(r *LauncherConfigFamilyMini) *docking_eventv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigList(r *LauncherConfigList) *docking_eventv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &docking_eventv1.LauncherConfigList{
		Families: func() []*docking_eventv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*docking_eventv1.LauncherConfigFamilyMini, len(r.Families))
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

func mapBizToProtoLocation(r *Location) *docking_eventv1.Location {
	if r == nil {
		return nil
	}
	return &docking_eventv1.Location{
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

func mapBizToProtoLocationSerializerNoCelestialBody(r *LocationSerializerNoCelestialBody) *docking_eventv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &docking_eventv1.LocationSerializerNoCelestialBody{
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

func mapBizToProtoMission(r *Mission) *docking_eventv1.Mission {
	if r == nil {
		return nil
	}
	return &docking_eventv1.Mission{
		Agencies: func() []*docking_eventv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*docking_eventv1.AgencyDetailed, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyDetailed(&r.Agencies[i])
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrls: func() []*docking_eventv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*docking_eventv1.InfoURL, len(r.InfoUrls))
			for i := range r.InfoUrls {
				res[i] = mapBizToProtoInfoURL(&r.InfoUrls[i])
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapBizToProtoOrbit(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*docking_eventv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*docking_eventv1.VidURL, len(r.VidUrls))
			for i := range r.VidUrls {
				res[i] = mapBizToProtoVidURL(&r.VidUrls[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *docking_eventv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &docking_eventv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoNetPrecision(r *NetPrecision) *docking_eventv1.NetPrecision {
	if r == nil {
		return nil
	}
	return &docking_eventv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoOrbit(r *Orbit) *docking_eventv1.Orbit {
	if r == nil {
		return nil
	}
	return &docking_eventv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapBizToProtoCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoPad(r *Pad) *docking_eventv1.Pad {
	if r == nil {
		return nil
	}
	return &docking_eventv1.Pad{
		Active: r.Active,
		Agencies: func() []*docking_eventv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*docking_eventv1.AgencyNormal, len(r.Agencies))
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

func mapBizToProtoPayloadFlightMini(r *PayloadFlightMini) *docking_eventv1.PayloadFlightMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.PayloadFlightMini{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapBizToProtoLanding(r.Landing),
		Launch: mapBizToProtoLaunchMini(r.Launch),
		Payload: mapBizToProtoPayloadMini(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
}

func mapBizToProtoPayloadFlightNormal(r *PayloadFlightNormal) *docking_eventv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.PayloadFlightNormal{
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

func mapBizToProtoPayloadMini(r *PayloadMini) *docking_eventv1.PayloadMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.PayloadMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapBizToProtoAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoPayloadType(r.TypeVal),
	}
}

func mapBizToProtoPayloadNormal(r *PayloadNormal) *docking_eventv1.PayloadNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapBizToProtoAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapBizToProtoAgencyNormal(r.Operator),
		Program: func() []*docking_eventv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*docking_eventv1.ProgramMini, len(r.Program))
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

func mapBizToProtoPayloadType(r *PayloadType) *docking_eventv1.PayloadType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoProgramMini(r *ProgramMini) *docking_eventv1.ProgramMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.ProgramMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoProgramNormal(r *ProgramNormal) *docking_eventv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.ProgramNormal{
		Agencies: func() []*docking_eventv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*docking_eventv1.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []*docking_eventv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*docking_eventv1.MissionPatch, len(r.MissionPatches))
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

func mapBizToProtoProgramType(r *ProgramType) *docking_eventv1.ProgramType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoRocketNormal(r *RocketNormal) *docking_eventv1.RocketNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.RocketNormal{
		Configuration: mapBizToProtoLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *docking_eventv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *docking_eventv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationMini(r *SpaceStationMini) *docking_eventv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpaceStationMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationNormal(r *SpaceStationNormal) *docking_eventv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpaceStationNormal{
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

func mapBizToProtoSpaceStationStatus(r *SpaceStationStatus) *docking_eventv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpaceStationType(r *SpaceStationType) *docking_eventv1.SpaceStationType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraftConfigFamilyMini(r *SpacecraftConfigFamilyMini) *docking_eventv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigFamilyNormal(r *SpacecraftConfigFamilyNormal) *docking_eventv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapBizToProtoSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigNormal(r *SpacecraftConfigNormal) *docking_eventv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpacecraftConfigNormal{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Family: func() []*docking_eventv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*docking_eventv1.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapBizToProtoSpacecraftConfigType(r *SpacecraftConfigType) *docking_eventv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraftFlightMini(r *SpacecraftFlightMini) *docking_eventv1.SpacecraftFlightMini {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpacecraftFlightMini{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapBizToProtoLanding(r.Landing),
		Launch: mapBizToProtoLaunchMini(r.Launch),
		MissionEnd: r.MissionEnd,
		Spacecraft: mapBizToProtoSpacecraftNormal(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftFlightNormal(r *SpacecraftFlightNormal) *docking_eventv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpacecraftFlightNormal{
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

func mapBizToProtoSpacecraftNormal(r *SpacecraftNormal) *docking_eventv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpacecraftNormal{
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

func mapBizToProtoSpacecraftStatus(r *SpacecraftStatus) *docking_eventv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &docking_eventv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoVidURL(r *VidURL) *docking_eventv1.VidURL {
	if r == nil {
		return nil
	}
	return &docking_eventv1.VidURL{
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

func mapBizToProtoVidURLType(r *VidURLType) *docking_eventv1.VidURLType {
	if r == nil {
		return nil
	}
	return &docking_eventv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

