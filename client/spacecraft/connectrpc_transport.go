package spacecraft

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizspacecraft "com.gitlab/pobochiigo/bhole/internal/spacecraft"
	spacecraftv1 "com.gitlab/pobochiigo/bhole/proto/spacecraft/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/spacecraft/v1/spacecraftv1connect"
	"connectrpc.com/connect"
)

func NewSpacecraftClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizspacecraft.Service {
	connectClient := v1connect.NewSpacecraftServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListSpacecrafts: transport.NewConnectClient(
			connectClient.ListSpacecrafts,
			encodeListSpacecraftsRequest,
			decodeListSpacecraftsResponse,
		),
		getSpacecraft: transport.NewConnectClient(
			connectClient.GetSpacecraft,
			encodeGetSpacecraftRequest,
			decodeGetSpacecraftResponse,
		),
	}
}

func encodeListSpacecraftsRequest(_ context.Context, req *bizspacecraft.ListSpacecraftsRequest) (*spacecraftv1.ListSpacecraftsRequest, error) {
	return &spacecraftv1.ListSpacecraftsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetSpacecraftRequest(_ context.Context, req *bizspacecraft.GetSpacecraftRequest) (*spacecraftv1.GetSpacecraftRequest, error) {
	return &spacecraftv1.GetSpacecraftRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListSpacecraftsResponse(ctx context.Context, resp *spacecraftv1.ListSpacecraftsResponse) (*bizspacecraft.ListSpacecraftsResponse, error) {
	results := make([]bizspacecraft.Spacecraft, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizSpacecraft(r)
	}
	return &bizspacecraft.ListSpacecraftsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetSpacecraftResponse(ctx context.Context, resp *spacecraftv1.GetSpacecraftResponse) (*bizspacecraft.Spacecraft, error) {
	if resp.Spacecraft == nil {
		return nil, nil
	}
	return mapProtoToBizSpacecraft(resp.Spacecraft), nil
}

func mapProtoToBizAgencyDetailed(r *spacecraftv1.AgencyDetailed) *bizspacecraft.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizspacecraft.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizspacecraft.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizspacecraft.Country, len(r.Country))
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
		SocialMediaLinks: func() []bizspacecraft.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizspacecraft.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAgencyMini(r *spacecraftv1.AgencyMini) *bizspacecraft.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizspacecraft.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *spacecraftv1.AgencyNormal) *bizspacecraft.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizspacecraft.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizspacecraft.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *spacecraftv1.AgencyType) *bizspacecraft.AgencyType {
	if r == nil {
		return nil
	}
	return &bizspacecraft.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCelestialBodyDetailed(r *spacecraftv1.CelestialBodyDetailed) *bizspacecraft.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &bizspacecraft.CelestialBodyDetailed{
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

func mapProtoToBizCelestialBodyMini(r *spacecraftv1.CelestialBodyMini) *bizspacecraft.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &bizspacecraft.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizCelestialBodyNormal(r *spacecraftv1.CelestialBodyNormal) *bizspacecraft.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.CelestialBodyNormal{
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

func mapProtoToBizCelestialBodyType(r *spacecraftv1.CelestialBodyType) *bizspacecraft.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &bizspacecraft.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *spacecraftv1.Country) *bizspacecraft.Country {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizImage(r *spacecraftv1.Image) *bizspacecraft.Image {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizspacecraft.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizspacecraft.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *spacecraftv1.ImageLicense) *bizspacecraft.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizspacecraft.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *spacecraftv1.ImageVariant) *bizspacecraft.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizspacecraft.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *spacecraftv1.ImageVariantType) *bizspacecraft.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizspacecraft.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizInfoURL(r *spacecraftv1.InfoURL) *bizspacecraft.InfoURL {
	if r == nil {
		return nil
	}
	return &bizspacecraft.InfoURL{
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

func mapProtoToBizInfoURLType(r *spacecraftv1.InfoURLType) *bizspacecraft.InfoURLType {
	if r == nil {
		return nil
	}
	return &bizspacecraft.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanding(r *spacecraftv1.Landing) *bizspacecraft.Landing {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Landing{
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

func mapProtoToBizLandingLocation(r *spacecraftv1.LandingLocation) *bizspacecraft.LandingLocation {
	if r == nil {
		return nil
	}
	return &bizspacecraft.LandingLocation{
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

func mapProtoToBizLandingType(r *spacecraftv1.LandingType) *bizspacecraft.LandingType {
	if r == nil {
		return nil
	}
	return &bizspacecraft.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanguage(r *spacecraftv1.Language) *bizspacecraft.Language {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLaunchNormal(r *spacecraftv1.LaunchNormal) *bizspacecraft.LaunchNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.LaunchNormal{
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
		Program: func() []bizspacecraft.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizspacecraft.ProgramNormal, len(r.Program))
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

func mapProtoToBizLaunchStatus(r *spacecraftv1.LaunchStatus) *bizspacecraft.LaunchStatus {
	if r == nil {
		return nil
	}
	return &bizspacecraft.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfigFamilyMini(r *spacecraftv1.LauncherConfigFamilyMini) *bizspacecraft.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizspacecraft.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigList(r *spacecraftv1.LauncherConfigList) *bizspacecraft.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &bizspacecraft.LauncherConfigList{
		Families: func() []bizspacecraft.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]bizspacecraft.LauncherConfigFamilyMini, len(r.Families))
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

func mapProtoToBizLocation(r *spacecraftv1.Location) *bizspacecraft.Location {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Location{
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

func mapProtoToBizLocationSerializerNoCelestialBody(r *spacecraftv1.LocationSerializerNoCelestialBody) *bizspacecraft.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &bizspacecraft.LocationSerializerNoCelestialBody{
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

func mapProtoToBizMission(r *spacecraftv1.Mission) *bizspacecraft.Mission {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Mission{
		Agencies: func() []bizspacecraft.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspacecraft.AgencyDetailed, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyDetailed(v)
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizspacecraft.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizspacecraft.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapProtoToBizOrbit(r.Orbit),
		TypeVal: r.Type,
		VidUrls: func() []bizspacecraft.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizspacecraft.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizMissionPatch(r *spacecraftv1.MissionPatch) *bizspacecraft.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizspacecraft.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizNetPrecision(r *spacecraftv1.NetPrecision) *bizspacecraft.NetPrecision {
	if r == nil {
		return nil
	}
	return &bizspacecraft.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizOrbit(r *spacecraftv1.Orbit) *bizspacecraft.Orbit {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapProtoToBizCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizPad(r *spacecraftv1.Pad) *bizspacecraft.Pad {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Pad{
		Active: r.Active,
		Agencies: func() []bizspacecraft.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspacecraft.AgencyNormal, len(r.Agencies))
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

func mapProtoToBizProgramNormal(r *spacecraftv1.ProgramNormal) *bizspacecraft.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.ProgramNormal{
		Agencies: func() []bizspacecraft.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizspacecraft.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizspacecraft.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizspacecraft.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *spacecraftv1.ProgramType) *bizspacecraft.ProgramType {
	if r == nil {
		return nil
	}
	return &bizspacecraft.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizRocketNormal(r *spacecraftv1.RocketNormal) *bizspacecraft.RocketNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.RocketNormal{
		Configuration: mapProtoToBizLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapProtoToBizSocialMedia(r *spacecraftv1.SocialMedia) *bizspacecraft.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *spacecraftv1.SocialMediaLink) *bizspacecraft.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftConfigDetailed(r *spacecraftv1.SpacecraftConfigDetailed) *bizspacecraft.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftConfigDetailed{
		Agency: mapProtoToBizAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []bizspacecraft.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]bizspacecraft.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapProtoToBizSpacecraftConfigFamilyDetailed(r *spacecraftv1.SpacecraftConfigFamilyDetailed) *bizspacecraft.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftConfigFamilyDetailed{
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

func mapProtoToBizSpacecraftConfigFamilyMini(r *spacecraftv1.SpacecraftConfigFamilyMini) *bizspacecraft.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigFamilyNormal(r *spacecraftv1.SpacecraftConfigFamilyNormal) *bizspacecraft.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapProtoToBizSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigNormal(r *spacecraftv1.SpacecraftConfigNormal) *bizspacecraft.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftConfigNormal{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Family: func() []bizspacecraft.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]bizspacecraft.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapProtoToBizSpacecraftConfigType(r *spacecraftv1.SpacecraftConfigType) *bizspacecraft.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacecraft(r *spacecraftv1.Spacecraft) *bizspacecraft.Spacecraft {
	if r == nil {
		return nil
	}
	return &bizspacecraft.Spacecraft{
		Description: r.Description,
		FastestTurnaround: r.FastestTurnaround,
		Flights: func() []bizspacecraft.SpacecraftFlightNormal {
			if r.Flights == nil {
				return nil
			}
			res := make([]bizspacecraft.SpacecraftFlightNormal, len(r.Flights))
			for i, v := range r.Flights {
				res[i] = *mapProtoToBizSpacecraftFlightNormal(v)
			}
			return res
		}(),
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

func mapProtoToBizSpacecraftFlightNormal(r *spacecraftv1.SpacecraftFlightNormal) *bizspacecraft.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftFlightNormal{
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

func mapProtoToBizSpacecraftNormal(r *spacecraftv1.SpacecraftNormal) *bizspacecraft.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftNormal{
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

func mapProtoToBizSpacecraftStatus(r *spacecraftv1.SpacecraftStatus) *bizspacecraft.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &bizspacecraft.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizVidURL(r *spacecraftv1.VidURL) *bizspacecraft.VidURL {
	if r == nil {
		return nil
	}
	return &bizspacecraft.VidURL{
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

func mapProtoToBizVidURLType(r *spacecraftv1.VidURLType) *bizspacecraft.VidURLType {
	if r == nil {
		return nil
	}
	return &bizspacecraft.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

