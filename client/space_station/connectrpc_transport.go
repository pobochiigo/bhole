package space_station

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizspace_station "github.com/pobochiigo/bhole/internal/space_station"
	space_stationv1 "github.com/pobochiigo/bhole/proto/space_station/v1"
	v1connect "github.com/pobochiigo/bhole/proto/space_station/v1/space_stationv1connect"
	"connectrpc.com/connect"
)

func NewSpaceStationClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizspace_station.Service {
	connectClient := v1connect.NewSpaceStationServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListSpaceStations: transport.NewConnectClient(
			connectClient.ListSpaceStations,
			encodeListSpaceStationsRequest,
			decodeListSpaceStationsResponse,
		),
		getSpaceStation: transport.NewConnectClient(
			connectClient.GetSpaceStation,
			encodeGetSpaceStationRequest,
			decodeGetSpaceStationResponse,
		),
	}
}

func encodeListSpaceStationsRequest(_ context.Context, req *bizspace_station.ListSpaceStationsRequest) (*space_stationv1.ListSpaceStationsRequest, error) {
	return &space_stationv1.ListSpaceStationsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetSpaceStationRequest(_ context.Context, req *bizspace_station.GetSpaceStationRequest) (*space_stationv1.GetSpaceStationRequest, error) {
	return &space_stationv1.GetSpaceStationRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListSpaceStationsResponse(ctx context.Context, resp *space_stationv1.ListSpaceStationsResponse) (*bizspace_station.ListSpaceStationsResponse, error) {
	results := make([]bizspace_station.SpaceStation, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizSpaceStation(r)
	}
	return &bizspace_station.ListSpaceStationsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetSpaceStationResponse(ctx context.Context, resp *space_stationv1.GetSpaceStationResponse) (*bizspace_station.SpaceStation, error) {
	if resp.SpaceStation == nil {
		return nil, nil
	}
	return mapProtoToBizSpaceStation(resp.SpaceStation), nil
}

func mapProtoToBizAgencyDetailed(r *space_stationv1.AgencyDetailed) *bizspace_station.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizspace_station.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizspace_station.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizspace_station.Country, len(r.Country))
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
		SocialMediaLinks: func() []bizspace_station.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizspace_station.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAgencyMini(r *space_stationv1.AgencyMini) *bizspace_station.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizspace_station.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *space_stationv1.AgencyNormal) *bizspace_station.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizspace_station.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizspace_station.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *space_stationv1.AgencyType) *bizspace_station.AgencyType {
	if r == nil {
		return nil
	}
	return &bizspace_station.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCelestialBodyDetailed(r *space_stationv1.CelestialBodyDetailed) *bizspace_station.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &bizspace_station.CelestialBodyDetailed{
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

func mapProtoToBizCelestialBodyMini(r *space_stationv1.CelestialBodyMini) *bizspace_station.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &bizspace_station.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizCelestialBodyNormal(r *space_stationv1.CelestialBodyNormal) *bizspace_station.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.CelestialBodyNormal{
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

func mapProtoToBizCelestialBodyType(r *space_stationv1.CelestialBodyType) *bizspace_station.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizspace_station.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *space_stationv1.Country) *bizspace_station.Country {
	if r == nil {
		return nil
	}
	return &bizspace_station.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizDockingEventDetailedSerializerForSpacestation(r *space_stationv1.DockingEventDetailedSerializerForSpacestation) *bizspace_station.DockingEventDetailedSerializerForSpacestation {
	if r == nil {
		return nil
	}
	return &bizspace_station.DockingEventDetailedSerializerForSpacestation{
		Departure: r.Departure,
		Docking: r.Docking,
		FlightVehicleChaser: mapProtoToBizSpacecraftFlightForDockingEvent(r.FlightVehicleChaser),
		Id: r.Id,
		PayloadFlightChaser: mapProtoToBizPayloadFlightNormal(r.PayloadFlightChaser),
		SpaceStationChaser: mapProtoToBizSpaceStationNormal(r.SpaceStationChaser),
		Url: r.Url,
	}
}

func mapProtoToBizDockingEventForChaserNormal(r *space_stationv1.DockingEventForChaserNormal) *bizspace_station.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.DockingEventForChaserNormal{
		Departure: r.Departure,
		Docking: r.Docking,
		DockingLocation: mapProtoToBizDockingLocation(r.DockingLocation),
		FlightVehicleTarget: mapProtoToBizSpacecraftFlightNormal(r.FlightVehicleTarget),
		Id: r.Id,
		PayloadFlightTarget: mapProtoToBizPayloadFlightNormal(r.PayloadFlightTarget),
		SpaceStationTarget: mapProtoToBizSpaceStationNormal(r.SpaceStationTarget),
		Url: r.Url,
	}
}

func mapProtoToBizDockingLocation(r *space_stationv1.DockingLocation) *bizspace_station.DockingLocation {
	if r == nil {
		return nil
	}
	return &bizspace_station.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapProtoToBizPayloadMini(r.Payload),
		Spacecraft: mapProtoToBizSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapProtoToBizSpaceStationMini(r.Spacestation),
	}
}

func mapProtoToBizDockingLocationSerializerForSpacestation(r *space_stationv1.DockingLocationSerializerForSpacestation) *bizspace_station.DockingLocationSerializerForSpacestation {
	if r == nil {
		return nil
	}
	return &bizspace_station.DockingLocationSerializerForSpacestation{
		CurrentlyDocked: mapProtoToBizDockingEventDetailedSerializerForSpacestation(r.CurrentlyDocked),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizExpeditionMini(r *space_stationv1.ExpeditionMini) *bizspace_station.ExpeditionMini {
	if r == nil {
		return nil
	}
	return &bizspace_station.ExpeditionMini{
		End: r.End,
		Id: r.Id,
		Name: r.Name,
		Start: r.Start,
		Url: r.Url,
	}
}

func mapProtoToBizImage(r *space_stationv1.Image) *bizspace_station.Image {
	if r == nil {
		return nil
	}
	return &bizspace_station.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizspace_station.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizspace_station.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *space_stationv1.ImageLicense) *bizspace_station.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizspace_station.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *space_stationv1.ImageVariant) *bizspace_station.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizspace_station.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *space_stationv1.ImageVariantType) *bizspace_station.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizspace_station.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizInfoURL(r *space_stationv1.InfoURL) *bizspace_station.InfoURL {
	if r == nil {
		return nil
	}
	return &bizspace_station.InfoURL{
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

func mapProtoToBizInfoURLType(r *space_stationv1.InfoURLType) *bizspace_station.InfoURLType {
	if r == nil {
		return nil
	}
	return &bizspace_station.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanding(r *space_stationv1.Landing) *bizspace_station.Landing {
	if r == nil {
		return nil
	}
	return &bizspace_station.Landing{
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

func mapProtoToBizLandingLocation(r *space_stationv1.LandingLocation) *bizspace_station.LandingLocation {
	if r == nil {
		return nil
	}
	return &bizspace_station.LandingLocation{
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

func mapProtoToBizLandingType(r *space_stationv1.LandingType) *bizspace_station.LandingType {
	if r == nil {
		return nil
	}
	return &bizspace_station.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanguage(r *space_stationv1.Language) *bizspace_station.Language {
	if r == nil {
		return nil
	}
	return &bizspace_station.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLaunchNormal(r *space_stationv1.LaunchNormal) *bizspace_station.LaunchNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.LaunchNormal{
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
		Program: func() []bizspace_station.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizspace_station.ProgramNormal, len(r.Program))
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

func mapProtoToBizLaunchStatus(r *space_stationv1.LaunchStatus) *bizspace_station.LaunchStatus {
	if r == nil {
		return nil
	}
	return &bizspace_station.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfigFamilyMini(r *space_stationv1.LauncherConfigFamilyMini) *bizspace_station.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizspace_station.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigList(r *space_stationv1.LauncherConfigList) *bizspace_station.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &bizspace_station.LauncherConfigList{
		Families: func() []bizspace_station.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]bizspace_station.LauncherConfigFamilyMini, len(r.Families))
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

func mapProtoToBizLocation(r *space_stationv1.Location) *bizspace_station.Location {
	if r == nil {
		return nil
	}
	return &bizspace_station.Location{
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

func mapProtoToBizLocationSerializerNoCelestialBody(r *space_stationv1.LocationSerializerNoCelestialBody) *bizspace_station.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &bizspace_station.LocationSerializerNoCelestialBody{
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

func mapProtoToBizMission(r *space_stationv1.Mission) *bizspace_station.Mission {
	if r == nil {
		return nil
	}
	return &bizspace_station.Mission{
		Agencies: func() []bizspace_station.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspace_station.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyDetailed(v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizspace_station.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizspace_station.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapProtoToBizOrbit(r.Orbit),
		TypeVal: r.Type,
		VidUrls: func() []bizspace_station.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizspace_station.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizMissionPatch(r *space_stationv1.MissionPatch) *bizspace_station.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizspace_station.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizNetPrecision(r *space_stationv1.NetPrecision) *bizspace_station.NetPrecision {
	if r == nil {
		return nil
	}
	return &bizspace_station.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizOrbit(r *space_stationv1.Orbit) *bizspace_station.Orbit {
	if r == nil {
		return nil
	}
	return &bizspace_station.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapProtoToBizCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizPad(r *space_stationv1.Pad) *bizspace_station.Pad {
	if r == nil {
		return nil
	}
	return &bizspace_station.Pad{
		Active: r.Active,
		Agencies: func() []bizspace_station.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspace_station.AgencyNormal, len(r.Agencies))
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

func mapProtoToBizPayloadFlightNormal(r *space_stationv1.PayloadFlightNormal) *bizspace_station.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.PayloadFlightNormal{
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

func mapProtoToBizPayloadMini(r *space_stationv1.PayloadMini) *bizspace_station.PayloadMini {
	if r == nil {
		return nil
	}
	return &bizspace_station.PayloadMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapProtoToBizAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizPayloadType(r.Type),
	}
}

func mapProtoToBizPayloadNormal(r *space_stationv1.PayloadNormal) *bizspace_station.PayloadNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapProtoToBizAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapProtoToBizAgencyNormal(r.Operator),
		Program: func() []bizspace_station.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]bizspace_station.ProgramMini, len(r.Program))
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

func mapProtoToBizPayloadType(r *space_stationv1.PayloadType) *bizspace_station.PayloadType {
	if r == nil {
		return nil
	}
	return &bizspace_station.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizProgramMini(r *space_stationv1.ProgramMini) *bizspace_station.ProgramMini {
	if r == nil {
		return nil
	}
	return &bizspace_station.ProgramMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizProgramNormal(r *space_stationv1.ProgramNormal) *bizspace_station.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.ProgramNormal{
		Agencies: func() []bizspace_station.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspace_station.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizspace_station.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizspace_station.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *space_stationv1.ProgramType) *bizspace_station.ProgramType {
	if r == nil {
		return nil
	}
	return &bizspace_station.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizRocketNormal(r *space_stationv1.RocketNormal) *bizspace_station.RocketNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.RocketNormal{
		Configuration: mapProtoToBizLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapProtoToBizSocialMedia(r *space_stationv1.SocialMedia) *bizspace_station.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizspace_station.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *space_stationv1.SocialMediaLink) *bizspace_station.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizspace_station.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStation(r *space_stationv1.SpaceStation) *bizspace_station.SpaceStation {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpaceStation{
		ActiveDockingEvents: func() []bizspace_station.DockingEventForChaserNormal {
			if r.ActiveDockingEvents == nil {
				return nil
			}
			res := make([]bizspace_station.DockingEventForChaserNormal, len(r.ActiveDockingEvents))
			for i, v := range r.ActiveDockingEvents {
				res[i] = *mapProtoToBizDockingEventForChaserNormal(v)
			}
			return res
		}(),
		ActiveExpeditions: func() []bizspace_station.ExpeditionMini {
			if r.ActiveExpeditions == nil {
				return nil
			}
			res := make([]bizspace_station.ExpeditionMini, len(r.ActiveExpeditions))
			for i, v := range r.ActiveExpeditions {
				res[i] = *mapProtoToBizExpeditionMini(v)
			}
			return res
		}(),
		Deorbited: r.Deorbited,
		Description: r.Description,
		DockedVehicles: r.DockedVehicles,
		DockingLocation: func() []bizspace_station.DockingLocationSerializerForSpacestation {
			if r.DockingLocation == nil {
				return nil
			}
			res := make([]bizspace_station.DockingLocationSerializerForSpacestation, len(r.DockingLocation))
			for i, v := range r.DockingLocation {
				res[i] = *mapProtoToBizDockingLocationSerializerForSpacestation(v)
			}
			return res
		}(),
		Founded: r.Founded,
		Height: r.Height,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Mass: r.Mass,
		Name: r.Name,
		OnboardCrew: r.OnboardCrew,
		Orbit: r.Orbit,
		Owners: func() []bizspace_station.AgencyNormal {
			if r.Owners == nil {
				return nil
			}
			res := make([]bizspace_station.AgencyNormal, len(r.Owners))
			for i, v := range r.Owners {
				res[i] = *mapProtoToBizAgencyNormal(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Status: mapProtoToBizSpaceStationStatus(r.Status),
		TypeVal: mapProtoToBizSpaceStationType(r.Type),
		Url: r.Url,
		Volume: r.Volume,
		Width: r.Width,
	}
}

func mapProtoToBizSpaceStationMini(r *space_stationv1.SpaceStationMini) *bizspace_station.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpaceStationMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationNormal(r *space_stationv1.SpaceStationNormal) *bizspace_station.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpaceStationNormal{
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

func mapProtoToBizSpaceStationStatus(r *space_stationv1.SpaceStationStatus) *bizspace_station.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpaceStationType(r *space_stationv1.SpaceStationType) *bizspace_station.SpaceStationType {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftConfigDetailed(r *space_stationv1.SpacecraftConfigDetailed) *bizspace_station.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftConfigDetailed{
		Agency: mapProtoToBizAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []bizspace_station.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]bizspace_station.SpacecraftConfigFamilyDetailed, len(r.Family))
			for i, v := range r.Family {
				res[i] = *mapProtoToBizSpacecraftConfigFamilyDetailed(v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FlightLife: r.FlightLife,
		Height: r.Height,
		History: r.History,
		HumanRated: r.HumanRated,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
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
		TypeVal: mapProtoToBizSpacecraftConfigType(r.Type),
		Url: r.Url,
		WikiLink: r.WikiLink,
	}
}

func mapProtoToBizSpacecraftConfigFamilyDetailed(r *space_stationv1.SpacecraftConfigFamilyDetailed) *bizspace_station.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftConfigFamilyDetailed{
		AttemptedLandings: r.AttemptedLandings,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyNormal(r.Manufacturer),
		Name: r.Name,
		Parent: mapProtoToBizSpacecraftConfigFamilyNormal(r.Parent),
		ResponseMode: r.ResponseMode,
		SpacecraftFlown: r.SpacecraftFlown,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
}

func mapProtoToBizSpacecraftConfigFamilyMini(r *space_stationv1.SpacecraftConfigFamilyMini) *bizspace_station.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigFamilyNormal(r *space_stationv1.SpacecraftConfigFamilyNormal) *bizspace_station.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapProtoToBizSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigNormal(r *space_stationv1.SpacecraftConfigNormal) *bizspace_station.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftConfigNormal{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Family: func() []bizspace_station.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]bizspace_station.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapProtoToBizSpacecraftConfigType(r *space_stationv1.SpacecraftConfigType) *bizspace_station.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftDetailed(r *space_stationv1.SpacecraftDetailed) *bizspace_station.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftDetailed{
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
		SpacecraftConfig: mapProtoToBizSpacecraftConfigDetailed(r.SpacecraftConfig),
		Status: mapProtoToBizSpacecraftStatus(r.Status),
		TimeDocked: r.TimeDocked,
		TimeInSpace: r.TimeInSpace,
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftFlightForDockingEvent(r *space_stationv1.SpacecraftFlightForDockingEvent) *bizspace_station.SpacecraftFlightForDockingEvent {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftFlightForDockingEvent{
		Id: r.Id,
		Launch: mapProtoToBizLaunchNormal(r.Launch),
		Spacecraft: mapProtoToBizSpacecraftDetailed(r.Spacecraft),
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftFlightNormal(r *space_stationv1.SpacecraftFlightNormal) *bizspace_station.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftFlightNormal{
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

func mapProtoToBizSpacecraftNormal(r *space_stationv1.SpacecraftNormal) *bizspace_station.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftNormal{
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

func mapProtoToBizSpacecraftStatus(r *space_stationv1.SpacecraftStatus) *bizspace_station.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &bizspace_station.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizVidURL(r *space_stationv1.VidURL) *bizspace_station.VidURL {
	if r == nil {
		return nil
	}
	return &bizspace_station.VidURL{
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

func mapProtoToBizVidURLType(r *space_stationv1.VidURLType) *bizspace_station.VidURLType {
	if r == nil {
		return nil
	}
	return &bizspace_station.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

