package launch

import (
	"context"

	"github.com/pobochiigo/bhole/internal/transport"
	launchv1 "github.com/pobochiigo/bhole/proto/launch/v1"
	v1connect "github.com/pobochiigo/bhole/proto/launch/v1/launchv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListLaunches transport.Handler[launchv1.ListLaunchesRequest, launchv1.ListLaunchesResponse]
	getLaunch    transport.Handler[launchv1.GetLaunchRequest, launchv1.GetLaunchResponse]
}

func (s *server) ListLaunches(ctx context.Context, req *connect.Request[launchv1.ListLaunchesRequest]) (*connect.Response[launchv1.ListLaunchesResponse], error) {
	return s.listListLaunches(ctx, req)
}

func (s *server) GetLaunch(ctx context.Context, req *connect.Request[launchv1.GetLaunchRequest]) (*connect.Response[launchv1.GetLaunchResponse], error) {
	return s.getLaunch(ctx, req)
}

func NewLaunchHandler(svc Service) v1connect.LaunchServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListLaunches: transport.NewConnectServer(
			eps.listListLaunches,
			decodeListLaunchesRequest,
			encodeListLaunchesResponse,
		),
		getLaunch: transport.NewConnectServer(
			eps.getLaunch,
			decodeGetLaunchRequest,
			encodeGetLaunchResponse,
		),
	}
}

func decodeListLaunchesRequest(_ context.Context, req *launchv1.ListLaunchesRequest) (*ListLaunchesRequest, error) {
	return &ListLaunchesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListLaunchesResponse(ctx context.Context, resp *ListLaunchesResponse) (*launchv1.ListLaunchesResponse, error) {
	results := make([]*launchv1.Launch, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoLaunch(&resp.Results[i])
	}
	return &launchv1.ListLaunchesResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLaunchRequest(_ context.Context, req *launchv1.GetLaunchRequest) (*GetLaunchRequest, error) {
	return &GetLaunchRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetLaunchResponse(ctx context.Context, resp *Launch) (*launchv1.GetLaunchResponse, error) {
	return &launchv1.GetLaunchResponse{
		Launch: mapBizToProtoLaunch(resp),
	}, nil
}

func mapBizToProtoAgencyDetailed(r *AgencyDetailed) *launchv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*launchv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*launchv1.Country, len(r.Country))
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
		SocialMediaLinks: func() []*launchv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*launchv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAgencyMini(r *AgencyMini) *launchv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &launchv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *launchv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &launchv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*launchv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*launchv1.Country, len(r.Country))
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

func mapBizToProtoAgencyType(r *AgencyType) *launchv1.AgencyType {
	if r == nil {
		return nil
	}
	return &launchv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronautDetailed(r *AstronautDetailed) *launchv1.AstronautDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.AstronautDetailed{
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
		Nationality: func() []*launchv1.Country {
			if r.Nationality == nil {
				return nil
			}
			res := make([]*launchv1.Country, len(r.Nationality))
			for i := range r.Nationality {
				res[i] = mapBizToProtoCountry(&r.Nationality[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		SocialMediaLinks: func() []*launchv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*launchv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAstronautFlight(r *AstronautFlight) *launchv1.AstronautFlight {
	if r == nil {
		return nil
	}
	return &launchv1.AstronautFlight{
		Astronaut: mapBizToProtoAstronautDetailed(r.Astronaut),
		Id: r.Id,
		Role: mapBizToProtoAstronautRole(r.Role),
	}
}

func mapBizToProtoAstronautRole(r *AstronautRole) *launchv1.AstronautRole {
	if r == nil {
		return nil
	}
	return &launchv1.AstronautRole{
		Id: r.Id,
		Priority: r.Priority,
		Role: r.Role,
	}
}

func mapBizToProtoAstronautStatus(r *AstronautStatus) *launchv1.AstronautStatus {
	if r == nil {
		return nil
	}
	return &launchv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronautType(r *AstronautType) *launchv1.AstronautType {
	if r == nil {
		return nil
	}
	return &launchv1.AstronautType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCelestialBodyDetailed(r *CelestialBodyDetailed) *launchv1.CelestialBodyDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.CelestialBodyDetailed{
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

func mapBizToProtoCelestialBodyMini(r *CelestialBodyMini) *launchv1.CelestialBodyMini {
	if r == nil {
		return nil
	}
	return &launchv1.CelestialBodyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoCelestialBodyNormal(r *CelestialBodyNormal) *launchv1.CelestialBodyNormal {
	if r == nil {
		return nil
	}
	return &launchv1.CelestialBodyNormal{
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

func mapBizToProtoCelestialBodyType(r *CelestialBodyType) *launchv1.CelestialBodyType {
	if r == nil {
		return nil
	}
	return &launchv1.CelestialBodyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *launchv1.Country {
	if r == nil {
		return nil
	}
	return &launchv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoDockingEventForChaserNormal(r *DockingEventForChaserNormal) *launchv1.DockingEventForChaserNormal {
	if r == nil {
		return nil
	}
	return &launchv1.DockingEventForChaserNormal{
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

func mapBizToProtoDockingLocation(r *DockingLocation) *launchv1.DockingLocation {
	if r == nil {
		return nil
	}
	return &launchv1.DockingLocation{
		Id: r.Id,
		Name: r.Name,
		Payload: mapBizToProtoPayloadMini(r.Payload),
		Spacecraft: mapBizToProtoSpacecraftConfigNormal(r.Spacecraft),
		Spacestation: mapBizToProtoSpaceStationMini(r.Spacestation),
	}
}

func mapBizToProtoFirstStageNormal(r *FirstStageNormal) *launchv1.FirstStageNormal {
	if r == nil {
		return nil
	}
	return &launchv1.FirstStageNormal{
		Id: r.Id,
		Landing: mapBizToProtoLanding(r.Landing),
		Launcher: mapBizToProtoLauncherNormal(r.Launcher),
		LauncherFlightNumber: r.LauncherFlightNumber,
		PreviousFlight: mapBizToProtoLaunchMini(r.PreviousFlight),
		PreviousFlightDate: r.PreviousFlightDate,
		Reused: r.Reused,
		TurnAroundTime: r.TurnAroundTime,
		Type: r.TypeVal,
	}
}

func mapBizToProtoImage(r *Image) *launchv1.Image {
	if r == nil {
		return nil
	}
	return &launchv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*launchv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*launchv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *launchv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &launchv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *launchv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &launchv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *launchv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &launchv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoInfoURL(r *InfoURL) *launchv1.InfoURL {
	if r == nil {
		return nil
	}
	return &launchv1.InfoURL{
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

func mapBizToProtoInfoURLType(r *InfoURLType) *launchv1.InfoURLType {
	if r == nil {
		return nil
	}
	return &launchv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanding(r *Landing) *launchv1.Landing {
	if r == nil {
		return nil
	}
	return &launchv1.Landing{
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

func mapBizToProtoLandingLocation(r *LandingLocation) *launchv1.LandingLocation {
	if r == nil {
		return nil
	}
	return &launchv1.LandingLocation{
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

func mapBizToProtoLandingType(r *LandingType) *launchv1.LandingType {
	if r == nil {
		return nil
	}
	return &launchv1.LandingType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanguage(r *Language) *launchv1.Language {
	if r == nil {
		return nil
	}
	return &launchv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLaunch(r *Launch) *launchv1.Launch {
	if r == nil {
		return nil
	}
	return &launchv1.Launch{
		AgencyLaunchAttemptCount: r.AgencyLaunchAttemptCount,
		AgencyLaunchAttemptCountYear: r.AgencyLaunchAttemptCountYear,
		Failreason: r.Failreason,
		FlightclubUrl: r.FlightclubUrl,
		Hashtag: r.Hashtag,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrls: func() []*launchv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*launchv1.InfoURL, len(r.InfoUrls))
			for i := range r.InfoUrls {
				res[i] = mapBizToProtoInfoURL(&r.InfoUrls[i])
			}
			return res
		}(),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		LaunchServiceProvider: mapBizToProtoAgencyDetailed(r.LaunchServiceProvider),
		LocationLaunchAttemptCount: r.LocationLaunchAttemptCount,
		LocationLaunchAttemptCountYear: r.LocationLaunchAttemptCountYear,
		Mission: mapBizToProtoMission(r.Mission),
		MissionPatches: func() []*launchv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*launchv1.MissionPatch, len(r.MissionPatches))
			for i := range r.MissionPatches {
				res[i] = mapBizToProtoMissionPatch(&r.MissionPatches[i])
			}
			return res
		}(),
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapBizToProtoNetPrecision(r.NetPrecision),
		OrbitalLaunchAttemptCount: r.OrbitalLaunchAttemptCount,
		OrbitalLaunchAttemptCountYear: r.OrbitalLaunchAttemptCountYear,
		Pad: mapBizToProtoPad(r.Pad),
		PadLaunchAttemptCount: r.PadLaunchAttemptCount,
		PadLaunchAttemptCountYear: r.PadLaunchAttemptCountYear,
		PadTurnaround: r.PadTurnaround,
		Probability: r.Probability,
		Program: func() []*launchv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramNormal, len(r.Program))
			for i := range r.Program {
				res[i] = mapBizToProtoProgramNormal(&r.Program[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Rocket: mapBizToProtoRocketDetailed(r.Rocket),
		Slug: r.Slug,
		Status: mapBizToProtoLaunchStatus(r.Status),
		Timeline: func() []*launchv1.TimelineEvent {
			if r.Timeline == nil {
				return nil
			}
			res := make([]*launchv1.TimelineEvent, len(r.Timeline))
			for i := range r.Timeline {
				res[i] = mapBizToProtoTimelineEvent(&r.Timeline[i])
			}
			return res
		}(),
		Updates: func() []*launchv1.Update {
			if r.Updates == nil {
				return nil
			}
			res := make([]*launchv1.Update, len(r.Updates))
			for i := range r.Updates {
				res[i] = mapBizToProtoUpdate(&r.Updates[i])
			}
			return res
		}(),
		Url: r.Url,
		VidUrls: func() []*launchv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*launchv1.VidURL, len(r.VidUrls))
			for i := range r.VidUrls {
				res[i] = mapBizToProtoVidURL(&r.VidUrls[i])
			}
			return res
		}(),
		WeatherConcerns: r.WeatherConcerns,
		WebcastLive: r.WebcastLive,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
}

func mapBizToProtoLaunchMini(r *LaunchMini) *launchv1.LaunchMini {
	if r == nil {
		return nil
	}
	return &launchv1.LaunchMini{
		Id: r.Id,
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoLaunchNormal(r *LaunchNormal) *launchv1.LaunchNormal {
	if r == nil {
		return nil
	}
	return &launchv1.LaunchNormal{
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
		Program: func() []*launchv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramNormal, len(r.Program))
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

func mapBizToProtoLaunchStatus(r *LaunchStatus) *launchv1.LaunchStatus {
	if r == nil {
		return nil
	}
	return &launchv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLauncherConfigDetailed(r *LauncherConfigDetailed) *launchv1.LauncherConfigDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.LauncherConfigDetailed{
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
		Families: func() []*launchv1.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]*launchv1.LauncherConfigFamilyDetailed, len(r.Families))
			for i := range r.Families {
				res[i] = mapBizToProtoLauncherConfigFamilyDetailed(&r.Families[i])
			}
			return res
		}(),
		FastestTurnaround: r.FastestTurnaround,
		FullName: r.FullName,
		GeoCapacity: r.GeoCapacity,
		GtoCapacity: r.GtoCapacity,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		IsPlaceholder: r.IsPlaceholder,
		LaunchCost: r.LaunchCost,
		LaunchMass: r.LaunchMass,
		Length: r.Length,
		LeoCapacity: r.LeoCapacity,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapBizToProtoAgencyDetailed(r.Manufacturer),
		MaxStage: r.MaxStage,
		MinStage: r.MinStage,
		Name: r.Name,
		PendingLaunches: r.PendingLaunches,
		Program: func() []*launchv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramNormal, len(r.Program))
			for i := range r.Program {
				res[i] = mapBizToProtoProgramNormal(&r.Program[i])
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

func mapBizToProtoLauncherConfigFamilyDetailed(r *LauncherConfigFamilyDetailed) *launchv1.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []*launchv1.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*launchv1.AgencyDetailed, len(r.Manufacturer))
			for i := range r.Manufacturer {
				res[i] = mapBizToProtoAgencyDetailed(&r.Manufacturer[i])
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapBizToProtoLauncherConfigFamilyNormal(r.Parent),
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SuccessfulLandings: r.SuccessfulLandings,
		SuccessfulLaunches: r.SuccessfulLaunches,
		TotalLaunchCount: r.TotalLaunchCount,
	}
}

func mapBizToProtoLauncherConfigFamilyMini(r *LauncherConfigFamilyMini) *launchv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &launchv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigFamilyNormal(r *LauncherConfigFamilyNormal) *launchv1.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &launchv1.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []*launchv1.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*launchv1.AgencyNormal, len(r.Manufacturer))
			for i := range r.Manufacturer {
				res[i] = mapBizToProtoAgencyNormal(&r.Manufacturer[i])
			}
			return res
		}(),
		Name: r.Name,
		Parent: mapBizToProtoLauncherConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigList(r *LauncherConfigList) *launchv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &launchv1.LauncherConfigList{
		Families: func() []*launchv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*launchv1.LauncherConfigFamilyMini, len(r.Families))
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

func mapBizToProtoLauncherNormal(r *LauncherNormal) *launchv1.LauncherNormal {
	if r == nil {
		return nil
	}
	return &launchv1.LauncherNormal{
		AttemptedLandings: r.AttemptedLandings,
		Details: r.Details,
		FastestTurnaround: r.FastestTurnaround,
		FirstLaunchDate: r.FirstLaunchDate,
		FlightProven: r.FlightProven,
		Flights: r.Flights,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		IsPlaceholder: r.IsPlaceholder,
		LastLaunchDate: r.LastLaunchDate,
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		Status: mapBizToProtoLauncherStatus(r.Status),
		SuccessfulLandings: r.SuccessfulLandings,
		Url: r.Url,
	}
}

func mapBizToProtoLauncherStatus(r *LauncherStatus) *launchv1.LauncherStatus {
	if r == nil {
		return nil
	}
	return &launchv1.LauncherStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLocation(r *Location) *launchv1.Location {
	if r == nil {
		return nil
	}
	return &launchv1.Location{
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

func mapBizToProtoLocationSerializerNoCelestialBody(r *LocationSerializerNoCelestialBody) *launchv1.LocationSerializerNoCelestialBody {
	if r == nil {
		return nil
	}
	return &launchv1.LocationSerializerNoCelestialBody{
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

func mapBizToProtoMission(r *Mission) *launchv1.Mission {
	if r == nil {
		return nil
	}
	return &launchv1.Mission{
		Agencies: func() []*launchv1.AgencyDetailed {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*launchv1.AgencyDetailed, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyDetailed(&r.Agencies[i])
			}
			return res
		}(),
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrls: func() []*launchv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*launchv1.InfoURL, len(r.InfoUrls))
			for i := range r.InfoUrls {
				res[i] = mapBizToProtoInfoURL(&r.InfoUrls[i])
			}
			return res
		}(),
		Name: r.Name,
		Orbit: mapBizToProtoOrbit(r.Orbit),
		Type: r.TypeVal,
		VidUrls: func() []*launchv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*launchv1.VidURL, len(r.VidUrls))
			for i := range r.VidUrls {
				res[i] = mapBizToProtoVidURL(&r.VidUrls[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *launchv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &launchv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoNetPrecision(r *NetPrecision) *launchv1.NetPrecision {
	if r == nil {
		return nil
	}
	return &launchv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoOrbit(r *Orbit) *launchv1.Orbit {
	if r == nil {
		return nil
	}
	return &launchv1.Orbit{
		Abbrev: r.Abbrev,
		CelestialBody: mapBizToProtoCelestialBodyMini(r.CelestialBody),
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoPad(r *Pad) *launchv1.Pad {
	if r == nil {
		return nil
	}
	return &launchv1.Pad{
		Active: r.Active,
		Agencies: func() []*launchv1.AgencyNormal {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*launchv1.AgencyNormal, len(r.Agencies))
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

func mapBizToProtoPayloadDetailed(r *PayloadDetailed) *launchv1.PayloadDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.PayloadDetailed{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapBizToProtoAgencyDetailed(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapBizToProtoAgencyDetailed(r.Operator),
		Program: func() []*launchv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramNormal, len(r.Program))
			for i := range r.Program {
				res[i] = mapBizToProtoProgramNormal(&r.Program[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoPayloadType(r.TypeVal),
		WikiLink: r.WikiLink,
	}
}

func mapBizToProtoPayloadFlightNormal(r *PayloadFlightNormal) *launchv1.PayloadFlightNormal {
	if r == nil {
		return nil
	}
	return &launchv1.PayloadFlightNormal{
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

func mapBizToProtoPayloadFlightSerializerNoLaunch(r *PayloadFlightSerializerNoLaunch) *launchv1.PayloadFlightSerializerNoLaunch {
	if r == nil {
		return nil
	}
	return &launchv1.PayloadFlightSerializerNoLaunch{
		Amount: r.Amount,
		Destination: r.Destination,
		DockingEvents: func() []*launchv1.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]*launchv1.DockingEventForChaserNormal, len(r.DockingEvents))
			for i := range r.DockingEvents {
				res[i] = mapBizToProtoDockingEventForChaserNormal(&r.DockingEvents[i])
			}
			return res
		}(),
		Id: r.Id,
		Landing: mapBizToProtoLanding(r.Landing),
		Payload: mapBizToProtoPayloadDetailed(r.Payload),
		ResponseMode: r.ResponseMode,
		Url: r.Url,
	}
}

func mapBizToProtoPayloadMini(r *PayloadMini) *launchv1.PayloadMini {
	if r == nil {
		return nil
	}
	return &launchv1.PayloadMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Operator: mapBizToProtoAgencyMini(r.Operator),
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoPayloadType(r.TypeVal),
	}
}

func mapBizToProtoPayloadNormal(r *PayloadNormal) *launchv1.PayloadNormal {
	if r == nil {
		return nil
	}
	return &launchv1.PayloadNormal{
		Cost: r.Cost,
		Description: r.Description,
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoLink: r.InfoLink,
		Manufacturer: mapBizToProtoAgencyNormal(r.Manufacturer),
		Mass: r.Mass,
		Name: r.Name,
		Operator: mapBizToProtoAgencyNormal(r.Operator),
		Program: func() []*launchv1.ProgramMini {
			if r.Program == nil {
				return nil
			}
			res := make([]*launchv1.ProgramMini, len(r.Program))
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

func mapBizToProtoPayloadType(r *PayloadType) *launchv1.PayloadType {
	if r == nil {
		return nil
	}
	return &launchv1.PayloadType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoProgramMini(r *ProgramMini) *launchv1.ProgramMini {
	if r == nil {
		return nil
	}
	return &launchv1.ProgramMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrl: r.InfoUrl,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Url: r.Url,
		WikiUrl: r.WikiUrl,
	}
}

func mapBizToProtoProgramNormal(r *ProgramNormal) *launchv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &launchv1.ProgramNormal{
		Agencies: func() []*launchv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*launchv1.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []*launchv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*launchv1.MissionPatch, len(r.MissionPatches))
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

func mapBizToProtoProgramType(r *ProgramType) *launchv1.ProgramType {
	if r == nil {
		return nil
	}
	return &launchv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoRocketDetailed(r *RocketDetailed) *launchv1.RocketDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.RocketDetailed{
		Configuration: mapBizToProtoLauncherConfigDetailed(r.Configuration),
		Id: r.Id,
		LauncherStage: func() []*launchv1.FirstStageNormal {
			if r.LauncherStage == nil {
				return nil
			}
			res := make([]*launchv1.FirstStageNormal, len(r.LauncherStage))
			for i := range r.LauncherStage {
				res[i] = mapBizToProtoFirstStageNormal(&r.LauncherStage[i])
			}
			return res
		}(),
		Payloads: func() []*launchv1.PayloadFlightSerializerNoLaunch {
			if r.Payloads == nil {
				return nil
			}
			res := make([]*launchv1.PayloadFlightSerializerNoLaunch, len(r.Payloads))
			for i := range r.Payloads {
				res[i] = mapBizToProtoPayloadFlightSerializerNoLaunch(&r.Payloads[i])
			}
			return res
		}(),
		SpacecraftStage: func() []*launchv1.SpacecraftFlightDetailedSerializerNoLaunch {
			if r.SpacecraftStage == nil {
				return nil
			}
			res := make([]*launchv1.SpacecraftFlightDetailedSerializerNoLaunch, len(r.SpacecraftStage))
			for i := range r.SpacecraftStage {
				res[i] = mapBizToProtoSpacecraftFlightDetailedSerializerNoLaunch(&r.SpacecraftStage[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoRocketNormal(r *RocketNormal) *launchv1.RocketNormal {
	if r == nil {
		return nil
	}
	return &launchv1.RocketNormal{
		Configuration: mapBizToProtoLauncherConfigList(r.Configuration),
		Id: r.Id,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *launchv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &launchv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *launchv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &launchv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationMini(r *SpaceStationMini) *launchv1.SpaceStationMini {
	if r == nil {
		return nil
	}
	return &launchv1.SpaceStationMini{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSpaceStationNormal(r *SpaceStationNormal) *launchv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &launchv1.SpaceStationNormal{
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

func mapBizToProtoSpaceStationStatus(r *SpaceStationStatus) *launchv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &launchv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpaceStationType(r *SpaceStationType) *launchv1.SpaceStationType {
	if r == nil {
		return nil
	}
	return &launchv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraftConfigDetailed(r *SpacecraftConfigDetailed) *launchv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftConfigDetailed{
		Agency: mapBizToProtoAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*launchv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*launchv1.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapBizToProtoSpacecraftConfigFamilyDetailed(r *SpacecraftConfigFamilyDetailed) *launchv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftConfigFamilyDetailed{
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

func mapBizToProtoSpacecraftConfigFamilyMini(r *SpacecraftConfigFamilyMini) *launchv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigFamilyNormal(r *SpacecraftConfigFamilyNormal) *launchv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapBizToProtoSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigNormal(r *SpacecraftConfigNormal) *launchv1.SpacecraftConfigNormal {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftConfigNormal{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Family: func() []*launchv1.SpacecraftConfigFamilyNormal {
			if r.Family == nil {
				return nil
			}
			res := make([]*launchv1.SpacecraftConfigFamilyNormal, len(r.Family))
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

func mapBizToProtoSpacecraftConfigType(r *SpacecraftConfigType) *launchv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacecraftDetailed(r *SpacecraftDetailed) *launchv1.SpacecraftDetailed {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftDetailed{
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

func mapBizToProtoSpacecraftFlightDetailedSerializerNoLaunch(r *SpacecraftFlightDetailedSerializerNoLaunch) *launchv1.SpacecraftFlightDetailedSerializerNoLaunch {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftFlightDetailedSerializerNoLaunch{
		Destination: r.Destination,
		DockingEvents: func() []*launchv1.DockingEventForChaserNormal {
			if r.DockingEvents == nil {
				return nil
			}
			res := make([]*launchv1.DockingEventForChaserNormal, len(r.DockingEvents))
			for i := range r.DockingEvents {
				res[i] = mapBizToProtoDockingEventForChaserNormal(&r.DockingEvents[i])
			}
			return res
		}(),
		Duration: r.Duration,
		Id: r.Id,
		Landing: mapBizToProtoLanding(r.Landing),
		LandingCrew: func() []*launchv1.AstronautFlight {
			if r.LandingCrew == nil {
				return nil
			}
			res := make([]*launchv1.AstronautFlight, len(r.LandingCrew))
			for i := range r.LandingCrew {
				res[i] = mapBizToProtoAstronautFlight(&r.LandingCrew[i])
			}
			return res
		}(),
		LaunchCrew: func() []*launchv1.AstronautFlight {
			if r.LaunchCrew == nil {
				return nil
			}
			res := make([]*launchv1.AstronautFlight, len(r.LaunchCrew))
			for i := range r.LaunchCrew {
				res[i] = mapBizToProtoAstronautFlight(&r.LaunchCrew[i])
			}
			return res
		}(),
		MissionEnd: r.MissionEnd,
		OnboardCrew: func() []*launchv1.AstronautFlight {
			if r.OnboardCrew == nil {
				return nil
			}
			res := make([]*launchv1.AstronautFlight, len(r.OnboardCrew))
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

func mapBizToProtoSpacecraftFlightNormal(r *SpacecraftFlightNormal) *launchv1.SpacecraftFlightNormal {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftFlightNormal{
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

func mapBizToProtoSpacecraftNormal(r *SpacecraftNormal) *launchv1.SpacecraftNormal {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftNormal{
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

func mapBizToProtoSpacecraftStatus(r *SpacecraftStatus) *launchv1.SpacecraftStatus {
	if r == nil {
		return nil
	}
	return &launchv1.SpacecraftStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoTimelineEvent(r *TimelineEvent) *launchv1.TimelineEvent {
	if r == nil {
		return nil
	}
	return &launchv1.TimelineEvent{
		RelativeTime: r.RelativeTime,
		Type: mapBizToProtoTimelineEventType(r.TypeVal),
	}
}

func mapBizToProtoTimelineEventType(r *TimelineEventType) *launchv1.TimelineEventType {
	if r == nil {
		return nil
	}
	return &launchv1.TimelineEventType{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
	}
}

func mapBizToProtoUpdate(r *Update) *launchv1.Update {
	if r == nil {
		return nil
	}
	return &launchv1.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
}

func mapBizToProtoVidURL(r *VidURL) *launchv1.VidURL {
	if r == nil {
		return nil
	}
	return &launchv1.VidURL{
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

func mapBizToProtoVidURLType(r *VidURLType) *launchv1.VidURLType {
	if r == nil {
		return nil
	}
	return &launchv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

