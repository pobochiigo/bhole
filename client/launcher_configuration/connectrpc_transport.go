package launcher_configuration

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizlauncher_configuration "com.gitlab/pobochiigo/bhole/internal/launcher_configuration"
	launcher_configurationv1 "com.gitlab/pobochiigo/bhole/proto/launcher_configuration/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/launcher_configuration/v1/launcher_configurationv1connect"
	"connectrpc.com/connect"
)

func NewLauncherConfigurationClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizlauncher_configuration.Service {
	connectClient := v1connect.NewLauncherConfigurationServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListLauncherConfigurations: transport.NewConnectClient(
			connectClient.ListLauncherConfigurations,
			encodeListLauncherConfigurationsRequest,
			decodeListLauncherConfigurationsResponse,
		),
		getLauncherConfiguration: transport.NewConnectClient(
			connectClient.GetLauncherConfiguration,
			encodeGetLauncherConfigurationRequest,
			decodeGetLauncherConfigurationResponse,
		),
	}
}

func encodeListLauncherConfigurationsRequest(_ context.Context, req *bizlauncher_configuration.ListLauncherConfigurationsRequest) (*launcher_configurationv1.ListLauncherConfigurationsRequest, error) {
	return &launcher_configurationv1.ListLauncherConfigurationsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetLauncherConfigurationRequest(_ context.Context, req *bizlauncher_configuration.GetLauncherConfigurationRequest) (*launcher_configurationv1.GetLauncherConfigurationRequest, error) {
	return &launcher_configurationv1.GetLauncherConfigurationRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListLauncherConfigurationsResponse(ctx context.Context, resp *launcher_configurationv1.ListLauncherConfigurationsResponse) (*bizlauncher_configuration.ListLauncherConfigurationsResponse, error) {
	results := make([]bizlauncher_configuration.LauncherConfiguration, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizLauncherConfiguration(r)
	}
	return &bizlauncher_configuration.ListLauncherConfigurationsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLauncherConfigurationResponse(ctx context.Context, resp *launcher_configurationv1.GetLauncherConfigurationResponse) (*bizlauncher_configuration.LauncherConfiguration, error) {
	if resp.LauncherConfiguration == nil {
		return nil, nil
	}
	return mapProtoToBizLauncherConfiguration(resp.LauncherConfiguration), nil
}

func mapProtoToBizAgencyDetailed(r *launcher_configurationv1.AgencyDetailed) *bizlauncher_configuration.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizlauncher_configuration.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.Country, len(r.Country))
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
		SocialMediaLinks: func() []bizlauncher_configuration.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAgencyMini(r *launcher_configurationv1.AgencyMini) *bizlauncher_configuration.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *launcher_configurationv1.AgencyNormal) *bizlauncher_configuration.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizlauncher_configuration.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *launcher_configurationv1.AgencyType) *bizlauncher_configuration.AgencyType {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *launcher_configurationv1.Country) *bizlauncher_configuration.Country {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizImage(r *launcher_configurationv1.Image) *bizlauncher_configuration.Image {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizlauncher_configuration.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *launcher_configurationv1.ImageLicense) *bizlauncher_configuration.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *launcher_configurationv1.ImageVariant) *bizlauncher_configuration.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *launcher_configurationv1.ImageVariantType) *bizlauncher_configuration.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfiguration(r *launcher_configurationv1.LauncherConfiguration) *bizlauncher_configuration.LauncherConfiguration {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.LauncherConfiguration{
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
		Families: func() []bizlauncher_configuration.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.LauncherConfigFamilyDetailed, len(r.Families))
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
		Program: func() []bizlauncher_configuration.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.ProgramNormal, len(r.Program))
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

func mapProtoToBizLauncherConfigFamilyDetailed(r *launcher_configurationv1.LauncherConfigFamilyDetailed) *bizlauncher_configuration.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []bizlauncher_configuration.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.AgencyDetailed, len(r.Manufacturer))
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

func mapProtoToBizLauncherConfigFamilyMini(r *launcher_configurationv1.LauncherConfigFamilyMini) *bizlauncher_configuration.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigFamilyNormal(r *launcher_configurationv1.LauncherConfigFamilyNormal) *bizlauncher_configuration.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []bizlauncher_configuration.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.AgencyNormal, len(r.Manufacturer))
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

func mapProtoToBizMissionPatch(r *launcher_configurationv1.MissionPatch) *bizlauncher_configuration.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizProgramNormal(r *launcher_configurationv1.ProgramNormal) *bizlauncher_configuration.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.ProgramNormal{
		Agencies: func() []bizlauncher_configuration.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizlauncher_configuration.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizlauncher_configuration.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *launcher_configurationv1.ProgramType) *bizlauncher_configuration.ProgramType {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSocialMedia(r *launcher_configurationv1.SocialMedia) *bizlauncher_configuration.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *launcher_configurationv1.SocialMediaLink) *bizlauncher_configuration.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizlauncher_configuration.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

