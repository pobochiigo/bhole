package agency

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizagency "com.gitlab/pobochiigo/bhole/internal/agency"
	agencyv1 "com.gitlab/pobochiigo/bhole/proto/agency/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/agency/v1/agencyv1connect"
	"connectrpc.com/connect"
)

func NewAgencyClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizagency.Service {
	connectClient := v1connect.NewAgencyServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListAgencies: transport.NewConnectClient(
			connectClient.ListAgencies,
			encodeListAgenciesRequest,
			decodeListAgenciesResponse,
		),
		getAgency: transport.NewConnectClient(
			connectClient.GetAgency,
			encodeGetAgencyRequest,
			decodeGetAgencyResponse,
		),
	}
}

func encodeListAgenciesRequest(_ context.Context, req *bizagency.ListAgenciesRequest) (*agencyv1.ListAgenciesRequest, error) {
	return &agencyv1.ListAgenciesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetAgencyRequest(_ context.Context, req *bizagency.GetAgencyRequest) (*agencyv1.GetAgencyRequest, error) {
	return &agencyv1.GetAgencyRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListAgenciesResponse(ctx context.Context, resp *agencyv1.ListAgenciesResponse) (*bizagency.ListAgenciesResponse, error) {
	results := make([]bizagency.Agency, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizAgency(r)
	}
	return &bizagency.ListAgenciesResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetAgencyResponse(ctx context.Context, resp *agencyv1.GetAgencyResponse) (*bizagency.Agency, error) {
	if resp.Agency == nil {
		return nil, nil
	}
	return mapProtoToBizAgency(resp.Agency), nil
}

func mapProtoToBizAgencyDetailed(r *agencyv1.AgencyDetailed) *bizagency.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &bizagency.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizagency.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizagency.Country, len(r.Country))
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
		SocialMediaLinks: func() []bizagency.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizagency.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapProtoToBizAgency(r *agencyv1.Agency) *bizagency.Agency {
	if r == nil {
		return nil
	}
	return &bizagency.Agency{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []bizagency.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizagency.Country, len(r.Country))
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
		LauncherList: func() []bizagency.LauncherConfigDetailedSerializerNoManufacturer {
			if r.LauncherList == nil {
				return nil
			}
			res := make([]bizagency.LauncherConfigDetailedSerializerNoManufacturer, len(r.LauncherList))
			for i, v := range r.LauncherList {
				res[i] = *mapProtoToBizLauncherConfigDetailedSerializerNoManufacturer(v)
			}
			return res
		}(),
		Launchers: r.Launchers,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapProtoToBizImage(r.SocialLogo),
		SocialMediaLinks: func() []bizagency.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]bizagency.SocialMediaLink, len(r.SocialMediaLinks))
			for i, v := range r.SocialMediaLinks {
				res[i] = *mapProtoToBizSocialMediaLink(v)
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SpacecraftList: func() []bizagency.SpacecraftConfigDetailed {
			if r.SpacecraftList == nil {
				return nil
			}
			res := make([]bizagency.SpacecraftConfigDetailed, len(r.SpacecraftList))
			for i, v := range r.SpacecraftList {
				res[i] = *mapProtoToBizSpacecraftConfigDetailed(v)
			}
			return res
		}(),
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

func mapProtoToBizAgencyMini(r *agencyv1.AgencyMini) *bizagency.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizagency.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyNormal(r *agencyv1.AgencyNormal) *bizagency.AgencyNormal {
	if r == nil {
		return nil
	}
	return &bizagency.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []bizagency.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]bizagency.Country, len(r.Country))
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

func mapProtoToBizAgencyType(r *agencyv1.AgencyType) *bizagency.AgencyType {
	if r == nil {
		return nil
	}
	return &bizagency.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizCountry(r *agencyv1.Country) *bizagency.Country {
	if r == nil {
		return nil
	}
	return &bizagency.Country{
		Alpha2Code: r.Alpha_2Code,
		Alpha3Code: r.Alpha_3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapProtoToBizImage(r *agencyv1.Image) *bizagency.Image {
	if r == nil {
		return nil
	}
	return &bizagency.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizagency.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizagency.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *agencyv1.ImageLicense) *bizagency.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizagency.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *agencyv1.ImageVariant) *bizagency.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizagency.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *agencyv1.ImageVariantType) *bizagency.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizagency.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfigDetailedSerializerNoManufacturer(r *agencyv1.LauncherConfigDetailedSerializerNoManufacturer) *bizagency.LauncherConfigDetailedSerializerNoManufacturer {
	if r == nil {
		return nil
	}
	return &bizagency.LauncherConfigDetailedSerializerNoManufacturer{
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
		Families: func() []bizagency.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]bizagency.LauncherConfigFamilyDetailed, len(r.Families))
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
		MaxStage: r.MaxStage,
		MinStage: r.MinStage,
		Name: r.Name,
		PendingLaunches: r.PendingLaunches,
		Program: func() []bizagency.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizagency.ProgramNormal, len(r.Program))
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

func mapProtoToBizLauncherConfigFamilyDetailed(r *agencyv1.LauncherConfigFamilyDetailed) *bizagency.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizagency.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []bizagency.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]bizagency.AgencyDetailed, len(r.Manufacturer))
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

func mapProtoToBizLauncherConfigFamilyMini(r *agencyv1.LauncherConfigFamilyMini) *bizagency.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizagency.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigFamilyNormal(r *agencyv1.LauncherConfigFamilyNormal) *bizagency.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizagency.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []bizagency.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]bizagency.AgencyNormal, len(r.Manufacturer))
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

func mapProtoToBizMissionPatch(r *agencyv1.MissionPatch) *bizagency.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizagency.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizProgramNormal(r *agencyv1.ProgramNormal) *bizagency.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizagency.ProgramNormal{
		Agencies: func() []bizagency.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizagency.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizagency.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizagency.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *agencyv1.ProgramType) *bizagency.ProgramType {
	if r == nil {
		return nil
	}
	return &bizagency.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSocialMedia(r *agencyv1.SocialMedia) *bizagency.SocialMedia {
	if r == nil {
		return nil
	}
	return &bizagency.SocialMedia{
		Id: r.Id,
		Logo: mapProtoToBizImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapProtoToBizSocialMediaLink(r *agencyv1.SocialMediaLink) *bizagency.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &bizagency.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapProtoToBizSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapProtoToBizSpacecraftConfigDetailed(r *agencyv1.SpacecraftConfigDetailed) *bizagency.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &bizagency.SpacecraftConfigDetailed{
		Agency: mapProtoToBizAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []bizagency.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]bizagency.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapProtoToBizSpacecraftConfigFamilyDetailed(r *agencyv1.SpacecraftConfigFamilyDetailed) *bizagency.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &bizagency.SpacecraftConfigFamilyDetailed{
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

func mapProtoToBizSpacecraftConfigFamilyMini(r *agencyv1.SpacecraftConfigFamilyMini) *bizagency.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizagency.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigFamilyNormal(r *agencyv1.SpacecraftConfigFamilyNormal) *bizagency.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &bizagency.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapProtoToBizAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapProtoToBizSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizSpacecraftConfigType(r *agencyv1.SpacecraftConfigType) *bizagency.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &bizagency.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

