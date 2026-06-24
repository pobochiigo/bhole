package spacewalk

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizspacewalk "github.com/pobochiigo/bhole/internal/spacewalk"
	spacewalkv1 "github.com/pobochiigo/bhole/proto/spacewalk/v1"
	v1connect "github.com/pobochiigo/bhole/proto/spacewalk/v1/spacewalkv1connect"
	"connectrpc.com/connect"
)

func NewSpacewalkClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizspacewalk.Service {
	connectClient := v1connect.NewSpacewalkServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListSpacewalks: transport.NewConnectClient(
			connectClient.ListSpacewalks,
			encodeListSpacewalksRequest,
			decodeListSpacewalksResponse,
		),
		getSpacewalk: transport.NewConnectClient(
			connectClient.GetSpacewalk,
			encodeGetSpacewalkRequest,
			decodeGetSpacewalkResponse,
		),
	}
}

func encodeListSpacewalksRequest(_ context.Context, req *bizspacewalk.ListSpacewalksRequest) (*spacewalkv1.ListSpacewalksRequest, error) {
	return &spacewalkv1.ListSpacewalksRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetSpacewalkRequest(_ context.Context, req *bizspacewalk.GetSpacewalkRequest) (*spacewalkv1.GetSpacewalkRequest, error) {
	return &spacewalkv1.GetSpacewalkRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListSpacewalksResponse(ctx context.Context, resp *spacewalkv1.ListSpacewalksResponse) (*bizspacewalk.ListSpacewalksResponse, error) {
	results := make([]bizspacewalk.Spacewalk, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizSpacewalk(r)
	}
	return &bizspacewalk.ListSpacewalksResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetSpacewalkResponse(ctx context.Context, resp *spacewalkv1.GetSpacewalkResponse) (*bizspacewalk.Spacewalk, error) {
	if resp.Spacewalk == nil {
		return nil, nil
	}
	return mapProtoToBizSpacewalk(resp.Spacewalk), nil
}

func mapProtoToBizAgencyDetailed(r *spacewalkv1.AgencyDetailed) *bizspacewalk.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizspacewalk.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizspacewalk.Country, len(r.Country))
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
		SocialMediaLinks: func() []bizspacewalk.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizspacewalk.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAgencyMini(r *spacewalkv1.AgencyMini) *bizspacewalk.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *spacewalkv1.AgencyNormal) *bizspacewalk.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizspacewalk.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizspacewalk.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *spacewalkv1.AgencyType) *bizspacewalk.AgencyType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautDetailed(r *spacewalkv1.AstronautDetailed) *bizspacewalk.AstronautDetailed {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AstronautDetailed{
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
		Nationality: func() []bizspacewalk.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]bizspacewalk.Country, len(r.Nationality))
			for i, v := range r.Nationality {
				res[i] = *mapProtoToBizCountry(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []bizspacewalk.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizspacewalk.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAstronautFlight(r *spacewalkv1.AstronautFlight) *bizspacewalk.AstronautFlight {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AstronautFlight{
		Astronaut: mapProtoToBizAstronautDetailed(r.Astronaut),
		Id: r.Id,
		Role: mapProtoToBizAstronautRole(r.Role),
	}
}

func mapProtoToBizAstronautRole(r *spacewalkv1.AstronautRole) *bizspacewalk.AstronautRole {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
}

func mapProtoToBizAstronautStatus(r *spacewalkv1.AstronautStatus) *bizspacewalk.AstronautStatus {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautType(r *spacewalkv1.AstronautType) *bizspacewalk.AstronautType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCelestialBodyDetailed(r *spacewalkv1.CelestialBodyDetailed) *bizspacewalk.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &bizspacewalk.CelestialBodyDetailed{
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

func mapProtoToBizCelestialBodyMini(r *spacewalkv1.CelestialBodyMini) *bizspacewalk.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &bizspacewalk.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizCelestialBodyNormal(r *spacewalkv1.CelestialBodyNormal) *bizspacewalk.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.CelestialBodyNormal{
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

func mapProtoToBizCelestialBodyType(r *spacewalkv1.CelestialBodyType) *bizspacewalk.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *spacewalkv1.Country) *bizspacewalk.Country {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizDockingEventForChaserNormal(r *spacewalkv1.DockingEventForChaserNormal) *bizspacewalk.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.DockingEventForChaserNormal{
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

func mapProtoToBizDockingLocation(r *spacewalkv1.DockingLocation) *bizspacewalk.DockingLocation {
	if r == nil {
		return nil
	}
	return &bizspacewalk.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapProtoToBizPayloadMini(r.Payload),
		Spacecraft: mapProtoToBizSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapProtoToBizSpaceStationMini(r.Spacestation),
	}
}

func mapProtoToBizEventNormal(r *spacewalkv1.EventNormal) *bizspacewalk.EventNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.EventNormal{
		Date: r.Date,
		DatePrecision: mapProtoToBizNetPrecision(r.DatePrecision),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizspacewalk.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizspacewalk.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		Location: r.Location,
		Name: r.Name,
		Slug: r.Slug,
		TypeVal: mapProtoToBizEventType(r.Type),
		Url: r.Url,
		VidUrls: func() []bizspacewalk.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizspacewalk.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
		WebcastLive: r.WebcastLive,
	}
}

func mapProtoToBizEventType(r *spacewalkv1.EventType) *bizspacewalk.EventType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.EventType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizExpeditionNormalSerializerForSpacewalk(r *spacewalkv1.ExpeditionNormalSerializerForSpacewalk) *bizspacewalk.ExpeditionNormalSerializerForSpacewalk {
	if r == nil {
		return nil
	}
	return &bizspacewalk.ExpeditionNormalSerializerForSpacewalk{
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []bizspacewalk.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizspacewalk.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = *mapProtoToBizMissionPatch(v)
			}
			return res
		}(),
		Name: r.Name,
		Spacestation: mapProtoToBizSpaceStationNormal(r.Spacestation),
		Start: r.Start,
		Url: r.Url,
	}
}

func mapProtoToBizImage(r *spacewalkv1.Image) *bizspacewalk.Image {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizspacewalk.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizspacewalk.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *spacewalkv1.ImageLicense) *bizspacewalk.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizspacewalk.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *spacewalkv1.ImageVariant) *bizspacewalk.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizspacewalk.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *spacewalkv1.ImageVariantType) *bizspacewalk.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizInfoURL(r *spacewalkv1.InfoURL) *bizspacewalk.InfoURL {
	if r == nil {
		return nil
	}
	return &bizspacewalk.InfoURL{
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

func mapProtoToBizInfoURLType(r *spacewalkv1.InfoURLType) *bizspacewalk.InfoURLType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanding(r *spacewalkv1.Landing) *bizspacewalk.Landing {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Landing{
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

func mapProtoToBizLandingLocation(r *spacewalkv1.LandingLocation) *bizspacewalk.LandingLocation {
	if r == nil {
		return nil
	}
	return &bizspacewalk.LandingLocation{
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

func mapProtoToBizLandingType(r *spacewalkv1.LandingType) *bizspacewalk.LandingType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanguage(r *spacewalkv1.Language) *bizspacewalk.Language {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLaunchNormal(r *spacewalkv1.LaunchNormal) *bizspacewalk.LaunchNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.LaunchNormal{
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
		Program: func() []bizspacewalk.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizspacewalk.ProgramNormal, len(r.Program))
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

func mapProtoToBizLaunchStatus(r *spacewalkv1.LaunchStatus) *bizspacewalk.LaunchStatus {
	if r == nil {
		return nil
	}
	return &bizspacewalk.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfigFamilyMini(r *spacewalkv1.LauncherConfigFamilyMini) *bizspacewalk.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizspacewalk.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigList(r *spacewalkv1.LauncherConfigList) *bizspacewalk.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &bizspacewalk.LauncherConfigList{
		Families: func() []bizspacewalk.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]bizspacewalk.LauncherConfigFamilyMini, len(r.Families))
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

func mapProtoToBizLocation(r *spacewalkv1.Location) *bizspacewalk.Location {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Location{
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

func mapProtoToBizLocationSerializerNoCelestialBody(r *spacewalkv1.LocationSerializerNoCelestialBody) *bizspacewalk.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &bizspacewalk.LocationSerializerNoCelestialBody{
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

func mapProtoToBizMission(r *spacewalkv1.Mission) *bizspacewalk.Mission {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Mission{
		Agencies: func() []bizspacewalk.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspacewalk.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyDetailed(v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizspacewalk.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizspacewalk.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapProtoToBizOrbit(r.Orbit),
		TypeVal: r.Type,
		VidUrls: func() []bizspacewalk.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizspacewalk.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizMissionPatch(r *spacewalkv1.MissionPatch) *bizspacewalk.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizspacewalk.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizNetPrecision(r *spacewalkv1.NetPrecision) *bizspacewalk.NetPrecision {
	if r == nil {
		return nil
	}
	return &bizspacewalk.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizOrbit(r *spacewalkv1.Orbit) *bizspacewalk.Orbit {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapProtoToBizCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizPad(r *spacewalkv1.Pad) *bizspacewalk.Pad {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Pad{
		Active: r.Active,
		Agencies: func() []bizspacewalk.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspacewalk.AgencyNormal, len(r.Agencies))
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

func mapProtoToBizPayloadFlightNormal(r *spacewalkv1.PayloadFlightNormal) *bizspacewalk.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.PayloadFlightNormal{
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

func mapProtoToBizPayloadMini(r *spacewalkv1.PayloadMini) *bizspacewalk.PayloadMini {
	if r == nil {
		return nil
	}
	return &bizspacewalk.PayloadMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapProtoToBizAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizPayloadType(r.Type),
	}
}

func mapProtoToBizPayloadNormal(r *spacewalkv1.PayloadNormal) *bizspacewalk.PayloadNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapProtoToBizAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapProtoToBizAgencyNormal(r.Operator),
		Program: func() []bizspacewalk.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]bizspacewalk.ProgramMini, len(r.Program))
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

func mapProtoToBizPayloadType(r *spacewalkv1.PayloadType) *bizspacewalk.PayloadType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizProgramMini(r *spacewalkv1.ProgramMini) *bizspacewalk.ProgramMini {
	if r == nil {
		return nil
	}
	return &bizspacewalk.ProgramMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapProtoToBizProgramNormal(r *spacewalkv1.ProgramNormal) *bizspacewalk.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.ProgramNormal{
		Agencies: func() []bizspacewalk.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspacewalk.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizspacewalk.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizspacewalk.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *spacewalkv1.ProgramType) *bizspacewalk.ProgramType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizRocketNormal(r *spacewalkv1.RocketNormal) *bizspacewalk.RocketNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.RocketNormal{
		Configuration: mapProtoToBizLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapProtoToBizSocialMedia(r *spacewalkv1.SocialMedia) *bizspacewalk.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *spacewalkv1.SocialMediaLink) *bizspacewalk.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationMini(r *spacewalkv1.SpaceStationMini) *bizspacewalk.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpaceStationMini{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSpaceStationNormal(r *spacewalkv1.SpaceStationNormal) *bizspacewalk.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpaceStationNormal{
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

func mapProtoToBizSpaceStationStatus(r *spacewalkv1.SpaceStationStatus) *bizspacewalk.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpaceStationType(r *spacewalkv1.SpaceStationType) *bizspacewalk.SpaceStationType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftConfigDetailed(r *spacewalkv1.SpacecraftConfigDetailed) *bizspacewalk.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftConfigDetailed{
		Agency: mapProtoToBizAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []bizspacewalk.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]bizspacewalk.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapProtoToBizSpacecraftConfigFamilyDetailed(r *spacewalkv1.SpacecraftConfigFamilyDetailed) *bizspacewalk.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftConfigFamilyDetailed{
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

func mapProtoToBizSpacecraftConfigFamilyMini(r *spacewalkv1.SpacecraftConfigFamilyMini) *bizspacewalk.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigFamilyNormal(r *spacewalkv1.SpacecraftConfigFamilyNormal) *bizspacewalk.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapProtoToBizSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigNormal(r *spacewalkv1.SpacecraftConfigNormal) *bizspacewalk.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftConfigNormal{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Family: func() []bizspacewalk.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]bizspacewalk.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapProtoToBizSpacecraftConfigType(r *spacewalkv1.SpacecraftConfigType) *bizspacewalk.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraftDetailed(r *spacewalkv1.SpacecraftDetailed) *bizspacewalk.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftDetailed{
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

func mapProtoToBizSpacecraftFlightDetailed(r *spacewalkv1.SpacecraftFlightDetailed) *bizspacewalk.SpacecraftFlightDetailed {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftFlightDetailed{
		Destination: r.Destination,
		DockingEvents: func() []bizspacewalk.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]bizspacewalk.DockingEventForChaserNormal, len(r.DockingEvents))
			for i, v := range r.DockingEvents {
				res[i] = *mapProtoToBizDockingEventForChaserNormal(v)
			}
			return res
		}(),
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapProtoToBizLanding(r.Landing),
		LandingCrew: func() []bizspacewalk.AstronautFlight {
			if r.LandingCrew == nil {
				return nil
			}
			res := make([]bizspacewalk.AstronautFlight, len(r.LandingCrew))
			for i, v := range r.LandingCrew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		Launch: mapProtoToBizLaunchNormal(r.Launch),
		LaunchCrew: func() []bizspacewalk.AstronautFlight {
			if r.LaunchCrew == nil {
				return nil
			}
			res := make([]bizspacewalk.AstronautFlight, len(r.LaunchCrew))
			for i, v := range r.LaunchCrew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		MissionEnd: r.MissionEnd,
		OnboardCrew: func() []bizspacewalk.AstronautFlight {
			if r.OnboardCrew == nil {
				return nil
			}
			res := make([]bizspacewalk.AstronautFlight, len(r.OnboardCrew))
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

func mapProtoToBizSpacecraftFlightNormal(r *spacewalkv1.SpacecraftFlightNormal) *bizspacewalk.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftFlightNormal{
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

func mapProtoToBizSpacecraftNormal(r *spacewalkv1.SpacecraftNormal) *bizspacewalk.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftNormal{
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

func mapProtoToBizSpacecraftStatus(r *spacewalkv1.SpacecraftStatus) *bizspacewalk.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &bizspacewalk.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacewalk(r *spacewalkv1.Spacewalk) *bizspacewalk.Spacewalk {
	if r == nil {
		return nil
	}
	return &bizspacewalk.Spacewalk{
		Crew: func() []bizspacewalk.AstronautFlight {
			if r.Crew == nil {
				return nil
			}
			res := make([]bizspacewalk.AstronautFlight, len(r.Crew))
			for i, v := range r.Crew {
				res[i] = *mapProtoToBizAstronautFlight(v)
			}
			return res
		}(),
		Duration: r.Duration,
		End: r.End,
		Event: mapProtoToBizEventNormal(r.Event),
		Expedition: mapProtoToBizExpeditionNormalSerializerForSpacewalk(r.Expedition),
		Id: r.Id,
		Location: r.Location,
		Name: r.Name,
		Program: func() []bizspacewalk.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizspacewalk.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = *mapProtoToBizProgramNormal(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SpacecraftFlight: mapProtoToBizSpacecraftFlightDetailed(r.SpacecraftFlight),
		Spacestation: mapProtoToBizSpaceStationNormal(r.Spacestation),
		Start: r.Start,
		Url: r.Url,
	}
}

func mapProtoToBizVidURL(r *spacewalkv1.VidURL) *bizspacewalk.VidURL {
	if r == nil {
		return nil
	}
	return &bizspacewalk.VidURL{
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

func mapProtoToBizVidURLType(r *spacewalkv1.VidURLType) *bizspacewalk.VidURLType {
	if r == nil {
		return nil
	}
	return &bizspacewalk.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

