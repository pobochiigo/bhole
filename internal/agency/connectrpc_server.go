package agency

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	agencyv1 "com.gitlab/pobochiigo/bhole/proto/agency/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/agency/v1/agencyv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListAgencies transport.Handler[agencyv1.ListAgenciesRequest, agencyv1.ListAgenciesResponse]
	getAgency    transport.Handler[agencyv1.GetAgencyRequest, agencyv1.GetAgencyResponse]
}

func (s *server) ListAgencies(ctx context.Context, req *connect.Request[agencyv1.ListAgenciesRequest]) (*connect.Response[agencyv1.ListAgenciesResponse], error) {
	return s.listListAgencies(ctx, req)
}

func (s *server) GetAgency(ctx context.Context, req *connect.Request[agencyv1.GetAgencyRequest]) (*connect.Response[agencyv1.GetAgencyResponse], error) {
	return s.getAgency(ctx, req)
}

func NewAgencyHandler(svc Service) v1connect.AgencyServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListAgencies: transport.NewConnectServer(
			eps.listListAgencies,
			decodeListAgenciesRequest,
			encodeListAgenciesResponse,
		),
		getAgency: transport.NewConnectServer(
			eps.getAgency,
			decodeGetAgencyRequest,
			encodeGetAgencyResponse,
		),
	}
}

func decodeListAgenciesRequest(_ context.Context, req *agencyv1.ListAgenciesRequest) (*ListAgenciesRequest, error) {
	return &ListAgenciesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListAgenciesResponse(ctx context.Context, resp *ListAgenciesResponse) (*agencyv1.ListAgenciesResponse, error) {
	results := make([]*agencyv1.Agency, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoAgency(&resp.Results[i])
	}
	return &agencyv1.ListAgenciesResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetAgencyRequest(_ context.Context, req *agencyv1.GetAgencyRequest) (*GetAgencyRequest, error) {
	return &GetAgencyRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetAgencyResponse(ctx context.Context, resp *Agency) (*agencyv1.GetAgencyResponse, error) {
	return &agencyv1.GetAgencyResponse{
		Agency: mapBizToProtoAgency(resp),
	}, nil
}

func mapBizToProtoAgencyDetailed(r *AgencyDetailed) *agencyv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &agencyv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*agencyv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*agencyv1.Country, len(r.Country))
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
		SocialMediaLinks: func() []*agencyv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*agencyv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAgency(r *Agency) *agencyv1.Agency {
	if r == nil {
		return nil
	}
	return &agencyv1.Agency{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*agencyv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*agencyv1.Country, len(r.Country))
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
		LauncherList: func() []*agencyv1.LauncherConfigDetailedSerializerNoManufacturer {
			if r.LauncherList == nil {
				return nil
			}
			res := make([]*agencyv1.LauncherConfigDetailedSerializerNoManufacturer, len(r.LauncherList))
			for i := range r.LauncherList {
				res[i] = mapBizToProtoLauncherConfigDetailedSerializerNoManufacturer(&r.LauncherList[i])
			}
			return res
		}(),
		Launchers: r.Launchers,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Parent: r.Parent,
		PendingLaunches: r.PendingLaunches,
		ResponseMode: r.ResponseMode,
		SocialLogo: mapBizToProtoImage(r.SocialLogo),
		SocialMediaLinks: func() []*agencyv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*agencyv1.SocialMediaLink, len(r.SocialMediaLinks))
			for i := range r.SocialMediaLinks {
				res[i] = mapBizToProtoSocialMediaLink(&r.SocialMediaLinks[i])
			}
			return res
		}(),
		Spacecraft: r.Spacecraft,
		SpacecraftList: func() []*agencyv1.SpacecraftConfigDetailed {
			if r.SpacecraftList == nil {
				return nil
			}
			res := make([]*agencyv1.SpacecraftConfigDetailed, len(r.SpacecraftList))
			for i := range r.SpacecraftList {
				res[i] = mapBizToProtoSpacecraftConfigDetailed(&r.SpacecraftList[i])
			}
			return res
		}(),
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

func mapBizToProtoAgencyMini(r *AgencyMini) *agencyv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &agencyv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *agencyv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &agencyv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*agencyv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*agencyv1.Country, len(r.Country))
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

func mapBizToProtoAgencyType(r *AgencyType) *agencyv1.AgencyType {
	if r == nil {
		return nil
	}
	return &agencyv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *agencyv1.Country {
	if r == nil {
		return nil
	}
	return &agencyv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoImage(r *Image) *agencyv1.Image {
	if r == nil {
		return nil
	}
	return &agencyv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*agencyv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*agencyv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *agencyv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &agencyv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *agencyv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &agencyv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *agencyv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &agencyv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLauncherConfigDetailedSerializerNoManufacturer(r *LauncherConfigDetailedSerializerNoManufacturer) *agencyv1.LauncherConfigDetailedSerializerNoManufacturer {
	if r == nil {
		return nil
	}
	return &agencyv1.LauncherConfigDetailedSerializerNoManufacturer{
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
		Families: func() []*agencyv1.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]*agencyv1.LauncherConfigFamilyDetailed, len(r.Families))
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
		MaxStage: r.MaxStage,
		MinStage: r.MinStage,
		Name: r.Name,
		PendingLaunches: r.PendingLaunches,
		Program: func() []*agencyv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*agencyv1.ProgramNormal, len(r.Program))
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

func mapBizToProtoLauncherConfigFamilyDetailed(r *LauncherConfigFamilyDetailed) *agencyv1.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &agencyv1.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []*agencyv1.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*agencyv1.AgencyDetailed, len(r.Manufacturer))
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

func mapBizToProtoLauncherConfigFamilyMini(r *LauncherConfigFamilyMini) *agencyv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &agencyv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigFamilyNormal(r *LauncherConfigFamilyNormal) *agencyv1.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &agencyv1.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []*agencyv1.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*agencyv1.AgencyNormal, len(r.Manufacturer))
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

func mapBizToProtoMissionPatch(r *MissionPatch) *agencyv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &agencyv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoProgramNormal(r *ProgramNormal) *agencyv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &agencyv1.ProgramNormal{
		Agencies: func() []*agencyv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*agencyv1.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []*agencyv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*agencyv1.MissionPatch, len(r.MissionPatches))
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

func mapBizToProtoProgramType(r *ProgramType) *agencyv1.ProgramType {
	if r == nil {
		return nil
	}
	return &agencyv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *agencyv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &agencyv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *agencyv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &agencyv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

func mapBizToProtoSpacecraftConfigDetailed(r *SpacecraftConfigDetailed) *agencyv1.SpacecraftConfigDetailed {
	if r == nil {
		return nil
	}
	return &agencyv1.SpacecraftConfigDetailed{
		Agency: mapBizToProtoAgencyNormal(r.Agency),
		AttemptedLandings: r.AttemptedLandings,
		Capability: r.Capability,
		CrewCapacity: r.CrewCapacity,
		Details: r.Details,
		Diameter: r.Diameter,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Family: func() []*agencyv1.SpacecraftConfigFamilyDetailed {
			if r.Family == nil {
				return nil
			}
			res := make([]*agencyv1.SpacecraftConfigFamilyDetailed, len(r.Family))
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

func mapBizToProtoSpacecraftConfigFamilyDetailed(r *SpacecraftConfigFamilyDetailed) *agencyv1.SpacecraftConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &agencyv1.SpacecraftConfigFamilyDetailed{
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

func mapBizToProtoSpacecraftConfigFamilyMini(r *SpacecraftConfigFamilyMini) *agencyv1.SpacecraftConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &agencyv1.SpacecraftConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigFamilyNormal(r *SpacecraftConfigFamilyNormal) *agencyv1.SpacecraftConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &agencyv1.SpacecraftConfigFamilyNormal{
		Description: r.Description,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: mapBizToProtoAgencyMini(r.Manufacturer),
		Name: r.Name,
		Parent: mapBizToProtoSpacecraftConfigFamilyMini(r.Parent),
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoSpacecraftConfigType(r *SpacecraftConfigType) *agencyv1.SpacecraftConfigType {
	if r == nil {
		return nil
	}
	return &agencyv1.SpacecraftConfigType{
		Id: r.Id,
		Name: r.Name,
	}
}

