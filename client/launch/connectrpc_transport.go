package launch

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizlaunch "github.com/pobochiigo/bhole/internal/launch"
	launchv1 "github.com/pobochiigo/bhole/proto/launch/v1"
	v1connect "github.com/pobochiigo/bhole/proto/launch/v1/launchv1connect"
	"connectrpc.com/connect"
)

func NewLaunchClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizlaunch.Service {
	connectClient := v1connect.NewLaunchServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListLaunches: transport.NewConnectClient(
			connectClient.ListLaunches,
			encodeListLaunchesRequest,
			decodeListLaunchesResponse,
		),
		getLaunch: transport.NewConnectClient(
			connectClient.GetLaunch,
			encodeGetLaunchRequest,
			decodeGetLaunchResponse,
		),
	}
}

func encodeListLaunchesRequest(_ context.Context, req *bizlaunch.ListLaunchesRequest) (*launchv1.ListLaunchesRequest, error) {
	return &launchv1.ListLaunchesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetLaunchRequest(_ context.Context, req *bizlaunch.GetLaunchRequest) (*launchv1.GetLaunchRequest, error) {
	return &launchv1.GetLaunchRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListLaunchesResponse(ctx context.Context, resp *launchv1.ListLaunchesResponse) (*bizlaunch.ListLaunchesResponse, error) {
	results := make([]bizlaunch.Launch, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizLaunch(r)
	}
	return &bizlaunch.ListLaunchesResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLaunchResponse(ctx context.Context, resp *launchv1.GetLaunchResponse) (*bizlaunch.Launch, error) {
	if resp.Launch == nil {
		return nil, nil
	}
	return mapProtoToBizLaunch(resp.Launch), nil
}

func mapProtoToBizAgencyDetailed(r *launchv1.AgencyDetailed) *bizlaunch.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizlaunch.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizlaunch.Country, len(r.Country))
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
		SocialMediaLinks: func() []bizlaunch.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizlaunch.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAgencyMini(r *launchv1.AgencyMini) *bizlaunch.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizlaunch.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *launchv1.AgencyNormal) *bizlaunch.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizlaunch.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizlaunch.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *launchv1.AgencyType) *bizlaunch.AgencyType {
	if r == nil {
		return nil
	}
	return &bizlaunch.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautDetailed(r *launchv1.AstronautDetailed) *bizlaunch.AstronautDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.AstronautDetailed{
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
		Nationality: func() []bizlaunch.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]bizlaunch.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = *mapProtoToBizCountry(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []bizlaunch.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizlaunch.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAstronautFlight(r *launchv1.AstronautFlight) *bizlaunch.AstronautFlight {
	if r == nil {
		return nil
	}
	return &bizlaunch.AstronautFlight{
		Astronaut: mapProtoToBizAstronautDetailed(r.Astronaut),
		Id: r.Id,
		Role: mapProtoToBizAstronautRole(r.Role),
	}
}

func mapProtoToBizAstronautRole(r *launchv1.AstronautRole) *bizlaunch.AstronautRole {
	if r == nil {
		return nil
	}
	return &bizlaunch.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
}

func mapProtoToBizAstronautStatus(r *launchv1.AstronautStatus) *bizlaunch.AstronautStatus {
	if r == nil {
		return nil
	}
	return &bizlaunch.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautType(r *launchv1.AstronautType) *bizlaunch.AstronautType {
	if r == nil {
		return nil
	}
	return &bizlaunch.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCelestialBodyDetailed(r *launchv1.CelestialBodyDetailed) *bizlaunch.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.CelestialBodyDetailed{
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

func mapProtoToBizCelestialBodyMini(r *launchv1.CelestialBodyMini) *bizlaunch.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &bizlaunch.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizCelestialBodyNormal(r *launchv1.CelestialBodyNormal) *bizlaunch.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.CelestialBodyNormal{
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

func mapProtoToBizCelestialBodyType(r *launchv1.CelestialBodyType) *bizlaunch.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizlaunch.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *launchv1.Country) *bizlaunch.Country {
	if r == nil {
		return nil
	}
	return &bizlaunch.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizDockingEventForChaserNormal(r *launchv1.DockingEventForChaserNormal) *bizlaunch.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.DockingEventForChaserNormal{
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

func mapProtoToBizDockingLocation(r *launchv1.DockingLocation) *bizlaunch.DockingLocation {
	if r == nil {
		return nil
	}
	return &bizlaunch.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapProtoToBizPayloadMini(r.Payload),
		Spacecraft: mapProtoToBizSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapProtoToBizSpaceStationMini(r.Spacestation),
	}
}

func mapProtoToBizFirstStageNormal(r *launchv1.FirstStageNormal) *bizlaunch.FirstStageNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.FirstStageNormal{
		Id: r.Id,
		Landing: mapProtoToBizLanding(r.Landing),
		Launcher: mapProtoToBizLauncherNormal(r.Launcher),
		LauncherFlightNumber: r.LauncherFlightNumber,
		PreviousFlight: mapProtoToBizLaunchMini(r.PreviousFlight),
		PreviousFlightDate: r.PreviousFlightDate,
		Reused: r.Reused,
		TurnAroundTime: r.TurnAroundTime,
		TypeVal: r.Type,
	}
}

func mapProtoToBizImage(r *launchv1.Image) *bizlaunch.Image {
	if r == nil {
		return nil
	}
	return &bizlaunch.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizlaunch.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizlaunch.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *launchv1.ImageLicense) *bizlaunch.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizlaunch.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *launchv1.ImageVariant) *bizlaunch.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizlaunch.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *launchv1.ImageVariantType) *bizlaunch.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizlaunch.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizInfoURL(r *launchv1.InfoURL) *bizlaunch.InfoURL {
	if r == nil {
		return nil
	}
	return &bizlaunch.InfoURL{
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

func mapProtoToBizInfoURLType(r *launchv1.InfoURLType) *bizlaunch.InfoURLType {
	if r == nil {
		return nil
	}
	return &bizlaunch.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanding(r *launchv1.Landing) *bizlaunch.Landing {
	if r == nil {
		return nil
	}
	return &bizlaunch.Landing{
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

func mapProtoToBizLandingLocation(r *launchv1.LandingLocation) *bizlaunch.LandingLocation {
	if r == nil {
		return nil
	}
	return &bizlaunch.LandingLocation{
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

func mapProtoToBizLandingType(r *launchv1.LandingType) *bizlaunch.LandingType {
	if r == nil {
		return nil
	}
	return &bizlaunch.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanguage(r *launchv1.Language) *bizlaunch.Language {
	if r == nil {
		return nil
	}
	return &bizlaunch.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLaunch(r *launchv1.Launch) *bizlaunch.Launch {
	if r == nil {
		return nil
	}
	return &bizlaunch.Launch{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		FlightclubUrl: r.FlightclubUrl,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizlaunch.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizlaunch.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapProtoToBizAgencyDetailed(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapProtoToBizMission(r.Mission),
		MissionPatches: func() []bizlaunch.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizlaunch.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = *mapProtoToBizMissionPatch(v)
			}
			return res
		}(),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapProtoToBizNetPrecision(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapProtoToBizPad(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		PadTurnaround: r.PadTurnaround,
		Probability: r.Probability,
		Program: func() []bizlaunch.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlaunch.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = *mapProtoToBizProgramNormal(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapProtoToBizRocketDetailed(r.Rocket),
		Slug: r.Slug,
		Status: mapProtoToBizLaunchStatus(r.Status),
		Timeline: func() []bizlaunch.TimelineEvent {
			if r.Timeline == nil {
				return nil
			}
			res := make([]bizlaunch.TimelineEvent, len(r.Timeline))
			for i, v := range r.Timeline {
				res[i] = *mapProtoToBizTimelineEvent(v)
			}
			return res
		}(),
		Updates: func() []bizlaunch.Update {
			if r.Updates == nil {
				return nil
			}
			res := make([]bizlaunch.Update, len(r.Updates))
			for i, v := range r.Updates {
				res[i] = *mapProtoToBizUpdate(v)
			}
			return res
		}(),
		Url: r.Url,
		VidUrls: func() []bizlaunch.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizlaunch.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
}

func mapProtoToBizLaunchMini(r *launchv1.LaunchMini) *bizlaunch.LaunchMini {
	if r == nil {
		return nil
	}
	return &bizlaunch.LaunchMini{
		Id: r.Id,
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizLaunchNormal(r *launchv1.LaunchNormal) *bizlaunch.LaunchNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.LaunchNormal{
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
		Program: func() []bizlaunch.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlaunch.ProgramNormal, len(r.Program))
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

func mapProtoToBizLaunchStatus(r *launchv1.LaunchStatus) *bizlaunch.LaunchStatus {
	if r == nil {
		return nil
	}
	return &bizlaunch.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfigDetailed(r *launchv1.LauncherConfigDetailed) *bizlaunch.LauncherConfigDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.LauncherConfigDetailed{
		Active: r.Active,
		Alias: r.Alias,
		Apogee: r.Apogee,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Families: func() []bizlaunch.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]bizlaunch.LauncherConfigFamilyDetailed, len(r.Families))
			for i, v := range r.Families {
				res[i] = *mapProtoToBizLauncherConfigFamilyDetailed(v)
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FullName: r.FullName,
		GeoCapacity: r.GeoCapacity,
		GtoCapacity: r.GtoCapacity,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		IsPlaceholder: r.IsPlaceholder,
		LaunchCost: r.LaunchCost,
		LaunchMass: r.LaunchMass,
		Length: r.Length,
		LeoCapacity: r.LeoCapacity,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyDetailed(r.Manufacturer),
		MaxStage: r.MaxStage,
		MinStage: r.MinStage,
		Name: r.Name,
		PendingLaunches: r.PendingLaunches,
		Program: func() []bizlaunch.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlaunch.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = *mapProtoToBizProgramNormal(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Reusable: r.Reusable,
		SsoCapacity: r.SsoCapacity,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		ToThrust: r.ToThrust,
		TotalLaunchCount: r.TotalLaunchCount,
		Url: r.Url,
		Variant: r.Variant,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizLauncherConfigFamilyDetailed(r *launchv1.LauncherConfigFamilyDetailed) *bizlaunch.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []bizlaunch.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]bizlaunch.AgencyDetailed, len(r.Manufacturer))
			for i, v := range r.Manufacturer {
				res[i] = *mapProtoToBizAgencyDetailed(v)
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapProtoToBizLauncherConfigFamilyNormal(r.Parent),
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
}

func mapProtoToBizLauncherConfigFamilyMini(r *launchv1.LauncherConfigFamilyMini) *bizlaunch.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizlaunch.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigFamilyNormal(r *launchv1.LauncherConfigFamilyNormal) *bizlaunch.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []bizlaunch.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]bizlaunch.AgencyNormal, len(r.Manufacturer))
			for i, v := range r.Manufacturer {
				res[i] = *mapProtoToBizAgencyNormal(v)
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapProtoToBizLauncherConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigList(r *launchv1.LauncherConfigList) *bizlaunch.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &bizlaunch.LauncherConfigList{
		Families: func() []bizlaunch.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]bizlaunch.LauncherConfigFamilyMini, len(r.Families))
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

func mapProtoToBizLauncherNormal(r *launchv1.LauncherNormal) *bizlaunch.LauncherNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.LauncherNormal{
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

func mapProtoToBizLauncherStatus(r *launchv1.LauncherStatus) *bizlaunch.LauncherStatus {
	if r == nil {
		return nil
	}
	return &bizlaunch.LauncherStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLocation(r *launchv1.Location) *bizlaunch.Location {
	if r == nil {
		return nil
	}
	return &bizlaunch.Location{
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

func mapProtoToBizLocationSerializerNoCelestialBody(r *launchv1.LocationSerializerNoCelestialBody) *bizlaunch.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &bizlaunch.LocationSerializerNoCelestialBody{
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

func mapProtoToBizMission(r *launchv1.Mission) *bizlaunch.Mission {
	if r == nil {
		return nil
	}
	return &bizlaunch.Mission{
		Agencies: func() []bizlaunch.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizlaunch.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyDetailed(v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizlaunch.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizlaunch.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapProtoToBizOrbit(r.Orbit),
		TypeVal: r.Type,
		VidUrls: func() []bizlaunch.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizlaunch.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizMissionPatch(r *launchv1.MissionPatch) *bizlaunch.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizlaunch.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizNetPrecision(r *launchv1.NetPrecision) *bizlaunch.NetPrecision {
	if r == nil {
		return nil
	}
	return &bizlaunch.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizOrbit(r *launchv1.Orbit) *bizlaunch.Orbit {
	if r == nil {
		return nil
	}
	return &bizlaunch.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapProtoToBizCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizPad(r *launchv1.Pad) *bizlaunch.Pad {
	if r == nil {
		return nil
	}
	return &bizlaunch.Pad{
		Active: r.Active,
		Agencies: func() []bizlaunch.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizlaunch.AgencyNormal, len(r.Agencies))
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

func mapProtoToBizPayloadDetailed(r *launchv1.PayloadDetailed) *bizlaunch.PayloadDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.PayloadDetailed{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapProtoToBizAgencyDetailed(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapProtoToBizAgencyDetailed(r.Operator),
		Program: func() []bizlaunch.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlaunch.ProgramNormal, len(r.Program))
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

func mapProtoToBizPayloadFlightNormal(r *launchv1.PayloadFlightNormal) *bizlaunch.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.PayloadFlightNormal{
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

func mapProtoToBizPayloadFlightSerializerNoLaunch(r *launchv1.PayloadFlightSerializerNoLaunch) *bizlaunch.PayloadFlightSerializerNoLaunch {
	if r == nil {
		return nil
	}
	return &bizlaunch.PayloadFlightSerializerNoLaunch{
		Amount: r.Amount,
		Destination: r.Destination,
		DockingEvents: func() []bizlaunch.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]bizlaunch.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = *mapProtoToBizDockingEventForChaserNormal(v)
			}
			return res
		}(),
		Id: r.Id,
		Landing: mapProtoToBizLanding(r.Landing),
		Payload: mapProtoToBizPayloadDetailed(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
}

func mapProtoToBizPayloadMini(r *launchv1.PayloadMini) *bizlaunch.PayloadMini {
	if r == nil {
		return nil
	}
	return &bizlaunch.PayloadMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapProtoToBizAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizPayloadType(r.Type),
	}
}

func mapProtoToBizPayloadNormal(r *launchv1.PayloadNormal) *bizlaunch.PayloadNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapProtoToBizAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapProtoToBizAgencyNormal(r.Operator),
		Program: func() []bizlaunch.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlaunch.ProgramMini, len(r.Program))
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

func mapProtoToBizPayloadType(r *launchv1.PayloadType) *bizlaunch.PayloadType {
	if r == nil {
		return nil
	}
	return &bizlaunch.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizProgramMini(r *launchv1.ProgramMini) *bizlaunch.ProgramMini {
	if r == nil {
		return nil
	}
	return &bizlaunch.ProgramMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizProgramNormal(r *launchv1.ProgramNormal) *bizlaunch.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.ProgramNormal{
		Agencies: func() []bizlaunch.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizlaunch.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizlaunch.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizlaunch.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *launchv1.ProgramType) *bizlaunch.ProgramType {
	if r == nil {
		return nil
	}
	return &bizlaunch.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizRocketDetailed(r *launchv1.RocketDetailed) *bizlaunch.RocketDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.RocketDetailed{
		Configuration: mapProtoToBizLauncherConfigDetailed(r.Configuration),
		Id: r.Id,
		LauncherStage: func() []bizlaunch.FirstStageNormal {
			if r.LauncherStage == nil {
				return nil
			}
			res := make([]bizlaunch.FirstStageNormal, len(r.LauncherStage))
			for i, v := range r.LauncherStage {
				res[i] = *mapProtoToBizFirstStageNormal(v)
			}
			return res
		}(),
		Payloads: func() []bizlaunch.PayloadFlightSerializerNoLaunch {
			if r.Payloads == nil {
				return nil
			}
			res := make([]bizlaunch.PayloadFlightSerializerNoLaunch, len(r.Payloads))
			for i, v := range r.Payloads {
				res[i] = *mapProtoToBizPayloadFlightSerializerNoLaunch(v)
			}
			return res
		}(),
		SpacecraftStage: func() []bizlaunch.SpacecraftFlightDetailedSerializerNoLaunch {
			if r.SpacecraftStage == nil {
				return nil
			}
			res := make([]bizlaunch.SpacecraftFlightDetailedSerializerNoLaunch, len(r.SpacecraftStage))
			for i, v := range r.SpacecraftStage {
				res[i] = *mapProtoToBizSpacecraftFlightDetailedSerializerNoLaunch(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizRocketNormal(r *launchv1.RocketNormal) *bizlaunch.RocketNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.RocketNormal{
		Configuration: mapProtoToBizLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapProtoToBizSocialMedia(r *launchv1.SocialMedia) *bizlaunch.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizlaunch.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *launchv1.SocialMediaLink) *bizlaunch.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizlaunch.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationMini(r *launchv1.SpaceStationMini) *bizlaunch.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpaceStationMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationNormal(r *launchv1.SpaceStationNormal) *bizlaunch.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpaceStationNormal{
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

func mapProtoToBizSpaceStationStatus(r *launchv1.SpaceStationStatus) *bizlaunch.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpaceStationType(r *launchv1.SpaceStationType) *bizlaunch.SpaceStationType {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftConfigDetailed(r *launchv1.SpacecraftConfigDetailed) *bizlaunch.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftConfigDetailed{
		Agency: mapProtoToBizAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []bizlaunch.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]bizlaunch.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapProtoToBizSpacecraftConfigFamilyDetailed(r *launchv1.SpacecraftConfigFamilyDetailed) *bizlaunch.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftConfigFamilyDetailed{
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

func mapProtoToBizSpacecraftConfigFamilyMini(r *launchv1.SpacecraftConfigFamilyMini) *bizlaunch.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigFamilyNormal(r *launchv1.SpacecraftConfigFamilyNormal) *bizlaunch.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapProtoToBizSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigNormal(r *launchv1.SpacecraftConfigNormal) *bizlaunch.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftConfigNormal{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Family: func() []bizlaunch.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]bizlaunch.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapProtoToBizSpacecraftConfigType(r *launchv1.SpacecraftConfigType) *bizlaunch.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftDetailed(r *launchv1.SpacecraftDetailed) *bizlaunch.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftDetailed{
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

func mapProtoToBizSpacecraftFlightDetailedSerializerNoLaunch(r *launchv1.SpacecraftFlightDetailedSerializerNoLaunch) *bizlaunch.SpacecraftFlightDetailedSerializerNoLaunch {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftFlightDetailedSerializerNoLaunch{
		Destination: r.Destination,
		DockingEvents: func() []bizlaunch.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]bizlaunch.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = *mapProtoToBizDockingEventForChaserNormal(v)
			}
			return res
		}(),
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapProtoToBizLanding(r.Landing),
		LandingCrew: func() []bizlaunch.AstronautFlight {
			if r.LandingCrew == nil {
				return nil
			}
			res := make([]bizlaunch.AstronautFlight, len(r.LandingCrew))
			for i, v := range r.LandingCrew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		LaunchCrew: func() []bizlaunch.AstronautFlight {
			if r.LaunchCrew == nil {
				return nil
			}
			res := make([]bizlaunch.AstronautFlight, len(r.LaunchCrew))
			for i, v := range r.LaunchCrew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		MissionEnd: r.MissionEnd,
		OnboardCrew: func() []bizlaunch.AstronautFlight {
			if r.OnboardCrew == nil {
				return nil
			}
			res := make([]bizlaunch.AstronautFlight, len(r.OnboardCrew))
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

func mapProtoToBizSpacecraftFlightNormal(r *launchv1.SpacecraftFlightNormal) *bizlaunch.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftFlightNormal{
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

func mapProtoToBizSpacecraftNormal(r *launchv1.SpacecraftNormal) *bizlaunch.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftNormal{
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

func mapProtoToBizSpacecraftStatus(r *launchv1.SpacecraftStatus) *bizlaunch.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &bizlaunch.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizTimelineEvent(r *launchv1.TimelineEvent) *bizlaunch.TimelineEvent {
	if r == nil {
		return nil
	}
	return &bizlaunch.TimelineEvent{
		RelativeTime: r.RelativeTime,
		TypeVal: mapProtoToBizTimelineEventType(r.Type),
	}
}

func mapProtoToBizTimelineEventType(r *launchv1.TimelineEventType) *bizlaunch.TimelineEventType {
	if r == nil {
		return nil
	}
	return &bizlaunch.TimelineEventType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
	}
}

func mapProtoToBizUpdate(r *launchv1.Update) *bizlaunch.Update {
	if r == nil {
		return nil
	}
	return &bizlaunch.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
}

func mapProtoToBizVidURL(r *launchv1.VidURL) *bizlaunch.VidURL {
	if r == nil {
		return nil
	}
	return &bizlaunch.VidURL{
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

func mapProtoToBizVidURLType(r *launchv1.VidURLType) *bizlaunch.VidURLType {
	if r == nil {
		return nil
	}
	return &bizlaunch.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

