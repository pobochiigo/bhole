package docking_event

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizdocking_event "com.gitlab/pobochiigo/bhole/internal/docking_event"
	docking_eventv1 "com.gitlab/pobochiigo/bhole/proto/docking_event/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/docking_event/v1/docking_eventv1connect"
	"connectrpc.com/connect"
)

func NewDockingEventClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizdocking_event.Service {
	connectClient := v1connect.NewDockingEventServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListDockingEvents: transport.NewConnectClient(
			connectClient.ListDockingEvents,
			encodeListDockingEventsRequest,
			decodeListDockingEventsResponse,
		),
		getDockingEvent: transport.NewConnectClient(
			connectClient.GetDockingEvent,
			encodeGetDockingEventRequest,
			decodeGetDockingEventResponse,
		),
	}
}

func encodeListDockingEventsRequest(_ context.Context, req *bizdocking_event.ListDockingEventsRequest) (*docking_eventv1.ListDockingEventsRequest, error) {
	return &docking_eventv1.ListDockingEventsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetDockingEventRequest(_ context.Context, req *bizdocking_event.GetDockingEventRequest) (*docking_eventv1.GetDockingEventRequest, error) {
	return &docking_eventv1.GetDockingEventRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListDockingEventsResponse(ctx context.Context, resp *docking_eventv1.ListDockingEventsResponse) (*bizdocking_event.ListDockingEventsResponse, error) {
	results := make([]bizdocking_event.DockingEvent, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizDockingEvent(r)
	}
	return &bizdocking_event.ListDockingEventsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetDockingEventResponse(ctx context.Context, resp *docking_eventv1.GetDockingEventResponse) (*bizdocking_event.DockingEvent, error) {
	if resp.DockingEvent == nil {
		return nil, nil
	}
	return mapProtoToBizDockingEvent(resp.DockingEvent), nil
}

func mapProtoToBizAgencyDetailed(r *docking_eventv1.AgencyDetailed) *bizdocking_event.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizdocking_event.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizdocking_event.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizdocking_event.Country, len(r.Country))
			for i, v := range r.Country {
				res[i] = *mapProtoToBizCountry(v)
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
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		Launchers: r.Launchers,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapProtoToBizImage(r.SocialLogo),
		SocialMediaLinks: func() []bizdocking_event.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizdocking_event.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = *mapProtoToBizSocialMediaLink(v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLandingsPayload: r.SuccessfulLandingsPayload,
		SuccessfulLandingsSpacecraft: r.SuccessfulLandingsSpacecraft,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizAgencyMini(r *docking_eventv1.AgencyMini) *bizdocking_event.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *docking_eventv1.AgencyNormal) *bizdocking_event.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizdocking_event.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizdocking_event.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *docking_eventv1.AgencyType) *bizdocking_event.AgencyType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCelestialBodyDetailed(r *docking_eventv1.CelestialBodyDetailed) *bizdocking_event.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &bizdocking_event.CelestialBodyDetailed{
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

func mapProtoToBizCelestialBodyMini(r *docking_eventv1.CelestialBodyMini) *bizdocking_event.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizCelestialBodyNormal(r *docking_eventv1.CelestialBodyNormal) *bizdocking_event.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.CelestialBodyNormal{
		Atmosphere: r.Atmosphere,
		Description: r.Description,
		Diameter: r.Diameter,
		Gravity: r.Gravity,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		LengthOfDay: r.LengthOfDay,
		Mass: r.Mass,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizCelestialBodyType(r.Type),
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizCelestialBodyType(r *docking_eventv1.CelestialBodyType) *bizdocking_event.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *docking_eventv1.Country) *bizdocking_event.Country {
	if r == nil {
		return nil
	}
	return &bizdocking_event.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizDockingEvent(r *docking_eventv1.DockingEvent) *bizdocking_event.DockingEvent {
	if r == nil {
		return nil
	}
	return &bizdocking_event.DockingEvent{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapProtoToBizDockingLocation(r.DockingLocation),
		FlightVehicleChaser: mapProtoToBizSpacecraftFlightNormal(r.FlightVehicleChaser),
		FlightVehicleTarget: mapProtoToBizSpacecraftFlightMini(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightChaser: mapProtoToBizPayloadFlightNormal(r.PayloadFlightChaser),
		PayloadFlightTarget: mapProtoToBizPayloadFlightMini(r.PayloadFlightTarget),
		ResponseMode: r.ResponseMode,
		SpaceStationChaser: mapProtoToBizSpaceStationNormal(r.SpaceStationChaser),
		SpaceStationTarget: mapProtoToBizSpaceStationMini(r.SpaceStationTarget),
		Url: r.Url,
	}
}

func mapProtoToBizDockingLocation(r *docking_eventv1.DockingLocation) *bizdocking_event.DockingLocation {
	if r == nil {
		return nil
	}
	return &bizdocking_event.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapProtoToBizPayloadMini(r.Payload),
		Spacecraft: mapProtoToBizSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapProtoToBizSpaceStationMini(r.Spacestation),
	}
}

func mapProtoToBizImage(r *docking_eventv1.Image) *bizdocking_event.Image {
	if r == nil {
		return nil
	}
	return &bizdocking_event.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizdocking_event.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizdocking_event.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *docking_eventv1.ImageLicense) *bizdocking_event.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizdocking_event.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *docking_eventv1.ImageVariant) *bizdocking_event.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizdocking_event.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *docking_eventv1.ImageVariantType) *bizdocking_event.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizInfoURL(r *docking_eventv1.InfoURL) *bizdocking_event.InfoURL {
	if r == nil {
		return nil
	}
	return &bizdocking_event.InfoURL{
		Description: r.Description,
		FeatureImage: r.FeatureImage,
		Language: mapProtoToBizLanguage(r.Language),
		Priority: r.Priority,
		Source: r.Source,
		Title: r.Title,
		TypeVal: mapProtoToBizInfoURLType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizInfoURLType(r *docking_eventv1.InfoURLType) *bizdocking_event.InfoURLType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanding(r *docking_eventv1.Landing) *bizdocking_event.Landing {
	if r == nil {
		return nil
	}
	return &bizdocking_event.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Id: r.Id,
		LandingLocation: mapProtoToBizLandingLocation(r.LandingLocation),
		Success: r.Success,
		TypeVal: mapProtoToBizLandingType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizLandingLocation(r *docking_eventv1.LandingLocation) *bizdocking_event.LandingLocation {
	if r == nil {
		return nil
	}
	return &bizdocking_event.LandingLocation{
		Abbrev: r.Abbrev,
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		CelestialBody: mapProtoToBizCelestialBodyNormal(r.CelestialBody),
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Latitude: r.Latitude,
		Location: mapProtoToBizLocationSerializerNoCelestialBody(r.Location),
		Longitude: r.Longitude,
		Name: r.Name,
		SuccessfulLandings: r.SuccessfulLandings,
	}
}

func mapProtoToBizLandingType(r *docking_eventv1.LandingType) *bizdocking_event.LandingType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanguage(r *docking_eventv1.Language) *bizdocking_event.Language {
	if r == nil {
		return nil
	}
	return &bizdocking_event.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLaunchMini(r *docking_eventv1.LaunchMini) *bizdocking_event.LaunchMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.LaunchMini{
		Id: r.Id,
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizLaunchNormal(r *docking_eventv1.LaunchNormal) *bizdocking_event.LaunchNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.LaunchNormal{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapProtoToBizAgencyMini(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapProtoToBizMission(r.Mission),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapProtoToBizNetPrecision(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapProtoToBizPad(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		Probability: r.Probability,
		Program: func() []bizdocking_event.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizdocking_event.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = *mapProtoToBizProgramNormal(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapProtoToBizRocketNormal(r.Rocket),
		Slug: r.Slug,
		Status: mapProtoToBizLaunchStatus(r.Status),
		Url: r.Url,
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
}

func mapProtoToBizLaunchStatus(r *docking_eventv1.LaunchStatus) *bizdocking_event.LaunchStatus {
	if r == nil {
		return nil
	}
	return &bizdocking_event.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfigFamilyMini(r *docking_eventv1.LauncherConfigFamilyMini) *bizdocking_event.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigList(r *docking_eventv1.LauncherConfigList) *bizdocking_event.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &bizdocking_event.LauncherConfigList{
		Families: func() []bizdocking_event.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]bizdocking_event.LauncherConfigFamilyMini, len(r.Families))
			for i, v := range r.Families {
				res[i] = *mapProtoToBizLauncherConfigFamilyMini(v)
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

func mapProtoToBizLocation(r *docking_eventv1.Location) *bizdocking_event.Location {
	if r == nil {
		return nil
	}
	return &bizdocking_event.Location{
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

func mapProtoToBizLocationSerializerNoCelestialBody(r *docking_eventv1.LocationSerializerNoCelestialBody) *bizdocking_event.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &bizdocking_event.LocationSerializerNoCelestialBody{
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

func mapProtoToBizMission(r *docking_eventv1.Mission) *bizdocking_event.Mission {
	if r == nil {
		return nil
	}
	return &bizdocking_event.Mission{
		Agencies: func() []bizdocking_event.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizdocking_event.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyDetailed(v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizdocking_event.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizdocking_event.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapProtoToBizOrbit(r.Orbit),
		TypeVal: r.Type,
		VidUrls: func() []bizdocking_event.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizdocking_event.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizMissionPatch(r *docking_eventv1.MissionPatch) *bizdocking_event.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizdocking_event.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizNetPrecision(r *docking_eventv1.NetPrecision) *bizdocking_event.NetPrecision {
	if r == nil {
		return nil
	}
	return &bizdocking_event.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizOrbit(r *docking_eventv1.Orbit) *bizdocking_event.Orbit {
	if r == nil {
		return nil
	}
	return &bizdocking_event.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapProtoToBizCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizPad(r *docking_eventv1.Pad) *bizdocking_event.Pad {
	if r == nil {
		return nil
	}
	return &bizdocking_event.Pad{
		Active: r.Active,
		Agencies: func() []bizdocking_event.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizdocking_event.AgencyNormal, len(r.Agencies))
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

func mapProtoToBizPayloadFlightMini(r *docking_eventv1.PayloadFlightMini) *bizdocking_event.PayloadFlightMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.PayloadFlightMini{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapProtoToBizLanding(r.Landing),
		Launch: mapProtoToBizLaunchMini(r.Launch),
		Payload: mapProtoToBizPayloadMini(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
}

func mapProtoToBizPayloadFlightNormal(r *docking_eventv1.PayloadFlightNormal) *bizdocking_event.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.PayloadFlightNormal{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapProtoToBizLanding(r.Landing),
		Launch: mapProtoToBizLaunchNormal(r.Launch),
		Payload: mapProtoToBizPayloadNormal(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
}

func mapProtoToBizPayloadMini(r *docking_eventv1.PayloadMini) *bizdocking_event.PayloadMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.PayloadMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapProtoToBizAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizPayloadType(r.Type),
	}
}

func mapProtoToBizPayloadNormal(r *docking_eventv1.PayloadNormal) *bizdocking_event.PayloadNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapProtoToBizAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapProtoToBizAgencyNormal(r.Operator),
		Program: func() []bizdocking_event.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]bizdocking_event.ProgramMini, len(r.Program))
			for i, v := range r.Program {
				res[i] = *mapProtoToBizProgramMini(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizPayloadType(r.Type),
		WikiLink: r.WikiLink,
	}
}

func mapProtoToBizPayloadType(r *docking_eventv1.PayloadType) *bizdocking_event.PayloadType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizProgramMini(r *docking_eventv1.ProgramMini) *bizdocking_event.ProgramMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.ProgramMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizProgramNormal(r *docking_eventv1.ProgramNormal) *bizdocking_event.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.ProgramNormal{
		Agencies: func() []bizdocking_event.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizdocking_event.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyMini(v)
			}
			return res
		}(),
		Description: r.Description,
		EndDate: r.EndDate,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		MissionPatches: func() []bizdocking_event.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizdocking_event.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = *mapProtoToBizMissionPatch(v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		StartDate: r.StartDate,
		TypeVal: mapProtoToBizProgramType(r.Type),
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizProgramType(r *docking_eventv1.ProgramType) *bizdocking_event.ProgramType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizRocketNormal(r *docking_eventv1.RocketNormal) *bizdocking_event.RocketNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.RocketNormal{
		Configuration: mapProtoToBizLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapProtoToBizSocialMedia(r *docking_eventv1.SocialMedia) *bizdocking_event.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *docking_eventv1.SocialMediaLink) *bizdocking_event.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationMini(r *docking_eventv1.SpaceStationMini) *bizdocking_event.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpaceStationMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationNormal(r *docking_eventv1.SpaceStationNormal) *bizdocking_event.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpaceStationNormal{
		Deorbited: r.Deorbited,
		Description: r.Description,
		Founded: r.Founded,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Name: r.Name,
		Orbit: r.Orbit,
		Status: mapProtoToBizSpaceStationStatus(r.Status),
		TypeVal: mapProtoToBizSpaceStationType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationStatus(r *docking_eventv1.SpaceStationStatus) *bizdocking_event.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpaceStationType(r *docking_eventv1.SpaceStationType) *bizdocking_event.SpaceStationType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftConfigFamilyMini(r *docking_eventv1.SpacecraftConfigFamilyMini) *bizdocking_event.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigFamilyNormal(r *docking_eventv1.SpacecraftConfigFamilyNormal) *bizdocking_event.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapProtoToBizSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigNormal(r *docking_eventv1.SpacecraftConfigNormal) *bizdocking_event.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpacecraftConfigNormal{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Family: func() []bizdocking_event.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]bizdocking_event.SpacecraftConfigFamilyNormal, len(r.Family))
			for i, v := range r.Family {
				res[i] = *mapProtoToBizSpacecraftConfigFamilyNormal(v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InUse: r.InUse,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizSpacecraftConfigType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftConfigType(r *docking_eventv1.SpacecraftConfigType) *bizdocking_event.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftFlightMini(r *docking_eventv1.SpacecraftFlightMini) *bizdocking_event.SpacecraftFlightMini {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpacecraftFlightMini{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapProtoToBizLanding(r.Landing),
		Launch: mapProtoToBizLaunchMini(r.Launch),
		MissionEnd: r.MissionEnd,
		Spacecraft: mapProtoToBizSpacecraftNormal(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftFlightNormal(r *docking_eventv1.SpacecraftFlightNormal) *bizdocking_event.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapProtoToBizLanding(r.Landing),
		Launch: mapProtoToBizLaunchNormal(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapProtoToBizSpacecraftNormal(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftNormal(r *docking_eventv1.SpacecraftNormal) *bizdocking_event.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpacecraftNormal{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		FlightsCount: r.FlightsCount,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InSpace: r.InSpace,
		IsPlaceholder: r.IsPlaceholder,
		MissionEndsCount: r.MissionEndsCount,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		SpacecraftConfig: mapProtoToBizSpacecraftConfigNormal(r.SpacecraftConfig),
		Status: mapProtoToBizSpacecraftStatus(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftStatus(r *docking_eventv1.SpacecraftStatus) *bizdocking_event.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &bizdocking_event.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizVidURL(r *docking_eventv1.VidURL) *bizdocking_event.VidURL {
	if r == nil {
		return nil
	}
	return &bizdocking_event.VidURL{
		Description: r.Description,
		EndTime: r.EndTime,
		FeatureImage: r.FeatureImage,
		Language: mapProtoToBizLanguage(r.Language),
		Live: r.Live,
		Priority: r.Priority,
		Publisher: r.Publisher,
		Source: r.Source,
		StartTime: r.StartTime,
		Title: r.Title,
		TypeVal: mapProtoToBizVidURLType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizVidURLType(r *docking_eventv1.VidURLType) *bizdocking_event.VidURLType {
	if r == nil {
		return nil
	}
	return &bizdocking_event.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

