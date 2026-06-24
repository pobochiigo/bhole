package landing

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizlanding "com.gitlab/pobochiigo/bhole/internal/landing"
	landingv1 "com.gitlab/pobochiigo/bhole/proto/landing/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/landing/v1/landingv1connect"
	"connectrpc.com/connect"
)

func NewLandingClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizlanding.Service {
	connectClient := v1connect.NewLandingServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListLandings: transport.NewConnectClient(
			connectClient.ListLandings,
			encodeListLandingsRequest,
			decodeListLandingsResponse,
		),
		getLanding: transport.NewConnectClient(
			connectClient.GetLanding,
			encodeGetLandingRequest,
			decodeGetLandingResponse,
		),
	}
}

func encodeListLandingsRequest(_ context.Context, req *bizlanding.ListLandingsRequest) (*landingv1.ListLandingsRequest, error) {
	return &landingv1.ListLandingsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetLandingRequest(_ context.Context, req *bizlanding.GetLandingRequest) (*landingv1.GetLandingRequest, error) {
	return &landingv1.GetLandingRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListLandingsResponse(ctx context.Context, resp *landingv1.ListLandingsResponse) (*bizlanding.ListLandingsResponse, error) {
	results := make([]bizlanding.Landing, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizLanding(r)
	}
	return &bizlanding.ListLandingsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLandingResponse(ctx context.Context, resp *landingv1.GetLandingResponse) (*bizlanding.Landing, error) {
	if resp.Landing == nil {
		return nil, nil
	}
	return mapProtoToBizLanding(resp.Landing), nil
}

func mapProtoToBizAgencyDetailed(r *landingv1.AgencyDetailed) *bizlanding.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizlanding.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizlanding.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizlanding.Country, len(r.Country))
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
		SocialMediaLinks: func() []bizlanding.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizlanding.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAgencyMini(r *landingv1.AgencyMini) *bizlanding.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizlanding.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *landingv1.AgencyNormal) *bizlanding.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizlanding.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizlanding.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *landingv1.AgencyType) *bizlanding.AgencyType {
	if r == nil {
		return nil
	}
	return &bizlanding.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautDetailed(r *landingv1.AstronautDetailed) *bizlanding.AstronautDetailed {
	if r == nil {
		return nil
	}
	return &bizlanding.AstronautDetailed{
		Age: r.Age,
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Bio: r.Bio,
		DateOfBirth: r.DateOfBirth,
		DateOfDeath: r.DateOfDeath,
		EvaTime: r.EvaTime,
		FirstFlight: r.FirstFlight,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InSpace: r.InSpace,
		LastFlight: r.LastFlight,
		Name: r.Name,
		Nationality: func() []bizlanding.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]bizlanding.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = *mapProtoToBizCountry(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []bizlanding.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizlanding.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = *mapProtoToBizSocialMediaLink(v)
			}
			return res
		}(),
		Status: mapProtoToBizAstronautStatus(r.Status),
		TimeInSpace: r.TimeInSpace,
		TypeVal: mapProtoToBizAstronautType(r.Type),
		Url: r.Url,
		Wiki: r.Wiki,
	}
}

func mapProtoToBizAstronautFlight(r *landingv1.AstronautFlight) *bizlanding.AstronautFlight {
	if r == nil {
		return nil
	}
	return &bizlanding.AstronautFlight{
		Astronaut: mapProtoToBizAstronautDetailed(r.Astronaut),
		Id: r.Id,
		Role: mapProtoToBizAstronautRole(r.Role),
	}
}

func mapProtoToBizAstronautRole(r *landingv1.AstronautRole) *bizlanding.AstronautRole {
	if r == nil {
		return nil
	}
	return &bizlanding.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
}

func mapProtoToBizAstronautStatus(r *landingv1.AstronautStatus) *bizlanding.AstronautStatus {
	if r == nil {
		return nil
	}
	return &bizlanding.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautType(r *landingv1.AstronautType) *bizlanding.AstronautType {
	if r == nil {
		return nil
	}
	return &bizlanding.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCelestialBodyDetailed(r *landingv1.CelestialBodyDetailed) *bizlanding.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &bizlanding.CelestialBodyDetailed{
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

func mapProtoToBizCelestialBodyMini(r *landingv1.CelestialBodyMini) *bizlanding.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &bizlanding.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizCelestialBodyNormal(r *landingv1.CelestialBodyNormal) *bizlanding.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.CelestialBodyNormal{
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

func mapProtoToBizCelestialBodyType(r *landingv1.CelestialBodyType) *bizlanding.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizlanding.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *landingv1.Country) *bizlanding.Country {
	if r == nil {
		return nil
	}
	return &bizlanding.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizDockingEventForChaserNormal(r *landingv1.DockingEventForChaserNormal) *bizlanding.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.DockingEventForChaserNormal{
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

func mapProtoToBizDockingLocation(r *landingv1.DockingLocation) *bizlanding.DockingLocation {
	if r == nil {
		return nil
	}
	return &bizlanding.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapProtoToBizPayloadMini(r.Payload),
		Spacecraft: mapProtoToBizSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapProtoToBizSpaceStationMini(r.Spacestation),
	}
}

func mapProtoToBizFirstStageDetailedSerializerNoLanding(r *landingv1.FirstStageDetailedSerializerNoLanding) *bizlanding.FirstStageDetailedSerializerNoLanding {
	if r == nil {
		return nil
	}
	return &bizlanding.FirstStageDetailedSerializerNoLanding{
		Id: r.Id,
		Launcher: mapProtoToBizLauncherNormal(r.Launcher),
		LauncherFlightNumber: r.LauncherFlightNumber,
		PreviousFlight: mapProtoToBizLaunchNormal(r.PreviousFlight),
		PreviousFlightDate: r.PreviousFlightDate,
		Reused: r.Reused,
		TurnAroundTime: r.TurnAroundTime,
		TypeVal: r.Type,
	}
}

func mapProtoToBizImage(r *landingv1.Image) *bizlanding.Image {
	if r == nil {
		return nil
	}
	return &bizlanding.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizlanding.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizlanding.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *landingv1.ImageLicense) *bizlanding.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizlanding.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *landingv1.ImageVariant) *bizlanding.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizlanding.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *landingv1.ImageVariantType) *bizlanding.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizlanding.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizInfoURL(r *landingv1.InfoURL) *bizlanding.InfoURL {
	if r == nil {
		return nil
	}
	return &bizlanding.InfoURL{
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

func mapProtoToBizInfoURLType(r *landingv1.InfoURLType) *bizlanding.InfoURLType {
	if r == nil {
		return nil
	}
	return &bizlanding.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLandingRecord(r *landingv1.LandingRecord) *bizlanding.LandingRecord {
	if r == nil {
		return nil
	}
	return &bizlanding.LandingRecord{
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

func mapProtoToBizLanding(r *landingv1.Landing) *bizlanding.Landing {
	if r == nil {
		return nil
	}
	return &bizlanding.Landing{
		Attempt: r.Attempt,
		Description: r.Description,
		DownrangeDistance: r.DownrangeDistance,
		Firststage: mapProtoToBizFirstStageDetailedSerializerNoLanding(r.Firststage),
		Id: r.Id,
		LandingLocation: mapProtoToBizLandingLocation(r.LandingLocation),
		Payloadflight: mapProtoToBizPayloadFlightDetailedSerializerNoLanding(r.Payloadflight),
		ResponseMode: r.ResponseMode,
		Spacecraftflight: mapProtoToBizSpacecraftFlightDetailedSerializerNoLanding(r.Spacecraftflight),
		Success: r.Success,
		TypeVal: mapProtoToBizLandingType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizLandingLocation(r *landingv1.LandingLocation) *bizlanding.LandingLocation {
	if r == nil {
		return nil
	}
	return &bizlanding.LandingLocation{
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

func mapProtoToBizLandingType(r *landingv1.LandingType) *bizlanding.LandingType {
	if r == nil {
		return nil
	}
	return &bizlanding.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanguage(r *landingv1.Language) *bizlanding.Language {
	if r == nil {
		return nil
	}
	return &bizlanding.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLaunchNormal(r *landingv1.LaunchNormal) *bizlanding.LaunchNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.LaunchNormal{
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
		Program: func() []bizlanding.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlanding.ProgramNormal, len(r.Program))
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

func mapProtoToBizLaunchStatus(r *landingv1.LaunchStatus) *bizlanding.LaunchStatus {
	if r == nil {
		return nil
	}
	return &bizlanding.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfigFamilyMini(r *landingv1.LauncherConfigFamilyMini) *bizlanding.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizlanding.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigList(r *landingv1.LauncherConfigList) *bizlanding.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &bizlanding.LauncherConfigList{
		Families: func() []bizlanding.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]bizlanding.LauncherConfigFamilyMini, len(r.Families))
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

func mapProtoToBizLauncherNormal(r *landingv1.LauncherNormal) *bizlanding.LauncherNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.LauncherNormal{
		AttemptedLandings: r.AttemptedLandings,
		Details: r.Details,
		FastestTurnaround: r.FastestTurnaround,
		FirstLaunchDate: r.FirstLaunchDate,
		FlightProven: r.FlightProven,
		Flights: r.Flights,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		IsPlaceholder: r.IsPlaceholder,
		LastLaunchDate: r.LastLaunchDate,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		Status: mapProtoToBizLauncherStatus(r.Status),
		SuccessfulLandings: r.SuccessfulLandings,
		Url: r.Url,
	}
}

func mapProtoToBizLauncherStatus(r *landingv1.LauncherStatus) *bizlanding.LauncherStatus {
	if r == nil {
		return nil
	}
	return &bizlanding.LauncherStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLocation(r *landingv1.Location) *bizlanding.Location {
	if r == nil {
		return nil
	}
	return &bizlanding.Location{
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

func mapProtoToBizLocationSerializerNoCelestialBody(r *landingv1.LocationSerializerNoCelestialBody) *bizlanding.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &bizlanding.LocationSerializerNoCelestialBody{
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

func mapProtoToBizMission(r *landingv1.Mission) *bizlanding.Mission {
	if r == nil {
		return nil
	}
	return &bizlanding.Mission{
		Agencies: func() []bizlanding.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizlanding.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyDetailed(v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizlanding.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizlanding.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapProtoToBizOrbit(r.Orbit),
		TypeVal: r.Type,
		VidUrls: func() []bizlanding.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizlanding.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizMissionPatch(r *landingv1.MissionPatch) *bizlanding.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizlanding.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizNetPrecision(r *landingv1.NetPrecision) *bizlanding.NetPrecision {
	if r == nil {
		return nil
	}
	return &bizlanding.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizOrbit(r *landingv1.Orbit) *bizlanding.Orbit {
	if r == nil {
		return nil
	}
	return &bizlanding.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapProtoToBizCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizPad(r *landingv1.Pad) *bizlanding.Pad {
	if r == nil {
		return nil
	}
	return &bizlanding.Pad{
		Active: r.Active,
		Agencies: func() []bizlanding.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizlanding.AgencyNormal, len(r.Agencies))
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

func mapProtoToBizPayloadDetailed(r *landingv1.PayloadDetailed) *bizlanding.PayloadDetailed {
	if r == nil {
		return nil
	}
	return &bizlanding.PayloadDetailed{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapProtoToBizAgencyDetailed(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapProtoToBizAgencyDetailed(r.Operator),
		Program: func() []bizlanding.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlanding.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = *mapProtoToBizProgramNormal(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizPayloadType(r.Type),
		WikiLink: r.WikiLink,
	}
}

func mapProtoToBizPayloadFlightDetailedSerializerNoLanding(r *landingv1.PayloadFlightDetailedSerializerNoLanding) *bizlanding.PayloadFlightDetailedSerializerNoLanding {
	if r == nil {
		return nil
	}
	return &bizlanding.PayloadFlightDetailedSerializerNoLanding{
		Amount: r.Amount,
		Destination: r.Destination,
		DockingEvents: func() []bizlanding.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]bizlanding.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = *mapProtoToBizDockingEventForChaserNormal(v)
			}
			return res
		}(),
		Id: r.Id,
		Launch: mapProtoToBizLaunchNormal(r.Launch),
		Payload: mapProtoToBizPayloadDetailed(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
}

func mapProtoToBizPayloadFlightNormal(r *landingv1.PayloadFlightNormal) *bizlanding.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.PayloadFlightNormal{
		Amount: r.Amount,
		Destination: r.Destination,
		Id: r.Id,
		Landing: mapProtoToBizLandingRecord(r.Landing),
		Launch: mapProtoToBizLaunchNormal(r.Launch),
		Payload: mapProtoToBizPayloadNormal(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
}

func mapProtoToBizPayloadMini(r *landingv1.PayloadMini) *bizlanding.PayloadMini {
	if r == nil {
		return nil
	}
	return &bizlanding.PayloadMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapProtoToBizAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizPayloadType(r.Type),
	}
}

func mapProtoToBizPayloadNormal(r *landingv1.PayloadNormal) *bizlanding.PayloadNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapProtoToBizAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapProtoToBizAgencyNormal(r.Operator),
		Program: func() []bizlanding.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlanding.ProgramMini, len(r.Program))
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

func mapProtoToBizPayloadType(r *landingv1.PayloadType) *bizlanding.PayloadType {
	if r == nil {
		return nil
	}
	return &bizlanding.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizProgramMini(r *landingv1.ProgramMini) *bizlanding.ProgramMini {
	if r == nil {
		return nil
	}
	return &bizlanding.ProgramMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizProgramNormal(r *landingv1.ProgramNormal) *bizlanding.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.ProgramNormal{
		Agencies: func() []bizlanding.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizlanding.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizlanding.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizlanding.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *landingv1.ProgramType) *bizlanding.ProgramType {
	if r == nil {
		return nil
	}
	return &bizlanding.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizRocketNormal(r *landingv1.RocketNormal) *bizlanding.RocketNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.RocketNormal{
		Configuration: mapProtoToBizLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapProtoToBizSocialMedia(r *landingv1.SocialMedia) *bizlanding.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizlanding.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *landingv1.SocialMediaLink) *bizlanding.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizlanding.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationMini(r *landingv1.SpaceStationMini) *bizlanding.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &bizlanding.SpaceStationMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationNormal(r *landingv1.SpaceStationNormal) *bizlanding.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.SpaceStationNormal{
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

func mapProtoToBizSpaceStationStatus(r *landingv1.SpaceStationStatus) *bizlanding.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &bizlanding.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpaceStationType(r *landingv1.SpaceStationType) *bizlanding.SpaceStationType {
	if r == nil {
		return nil
	}
	return &bizlanding.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftConfigDetailed(r *landingv1.SpacecraftConfigDetailed) *bizlanding.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftConfigDetailed{
		Agency: mapProtoToBizAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []bizlanding.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]bizlanding.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapProtoToBizSpacecraftConfigFamilyDetailed(r *landingv1.SpacecraftConfigFamilyDetailed) *bizlanding.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftConfigFamilyDetailed{
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

func mapProtoToBizSpacecraftConfigFamilyMini(r *landingv1.SpacecraftConfigFamilyMini) *bizlanding.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigFamilyNormal(r *landingv1.SpacecraftConfigFamilyNormal) *bizlanding.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapProtoToBizSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigNormal(r *landingv1.SpacecraftConfigNormal) *bizlanding.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftConfigNormal{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Family: func() []bizlanding.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]bizlanding.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapProtoToBizSpacecraftConfigType(r *landingv1.SpacecraftConfigType) *bizlanding.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftDetailed(r *landingv1.SpacecraftDetailed) *bizlanding.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftDetailed{
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

func mapProtoToBizSpacecraftFlightDetailedSerializerNoLanding(r *landingv1.SpacecraftFlightDetailedSerializerNoLanding) *bizlanding.SpacecraftFlightDetailedSerializerNoLanding {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftFlightDetailedSerializerNoLanding{
		Destination: r.Destination,
		DockingEvents: func() []bizlanding.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]bizlanding.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = *mapProtoToBizDockingEventForChaserNormal(v)
			}
			return res
		}(),
		Duration: r.Duration,
		Id: r.Id,
		LandingCrew: func() []bizlanding.AstronautFlight {
			if r.LandingCrew == nil {
				return nil
			}
			res := make([]bizlanding.AstronautFlight, len(r.LandingCrew))
			for i, v := range r.LandingCrew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		Launch: mapProtoToBizLaunchNormal(r.Launch),
		LaunchCrew: func() []bizlanding.AstronautFlight {
			if r.LaunchCrew == nil {
				return nil
			}
			res := make([]bizlanding.AstronautFlight, len(r.LaunchCrew))
			for i, v := range r.LaunchCrew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		MissionEnd: r.MissionEnd,
		OnboardCrew: func() []bizlanding.AstronautFlight {
			if r.OnboardCrew == nil {
				return nil
			}
			res := make([]bizlanding.AstronautFlight, len(r.OnboardCrew))
			for i, v := range r.OnboardCrew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Spacecraft: mapProtoToBizSpacecraftDetailed(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftFlightNormal(r *landingv1.SpacecraftFlightNormal) *bizlanding.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftFlightNormal{
		Destination: r.Destination,
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapProtoToBizLandingRecord(r.Landing),
		Launch: mapProtoToBizLaunchNormal(r.Launch),
		MissionEnd: r.MissionEnd,
		ResponseMode: r.ResponseMode,
		Spacecraft: mapProtoToBizSpacecraftNormal(r.Spacecraft),
		TurnAroundTime: r.TurnAroundTime,
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftNormal(r *landingv1.SpacecraftNormal) *bizlanding.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftNormal{
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

func mapProtoToBizSpacecraftStatus(r *landingv1.SpacecraftStatus) *bizlanding.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &bizlanding.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizVidURL(r *landingv1.VidURL) *bizlanding.VidURL {
	if r == nil {
		return nil
	}
	return &bizlanding.VidURL{
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

func mapProtoToBizVidURLType(r *landingv1.VidURLType) *bizlanding.VidURLType {
	if r == nil {
		return nil
	}
	return &bizlanding.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

