package launcher_configuration

import (
	"context"

	"github.com/pobochiigo/bhole/internal/transport"
	launcher_configurationv1 "github.com/pobochiigo/bhole/proto/launcher_configuration/v1"
	v1connect "github.com/pobochiigo/bhole/proto/launcher_configuration/v1/launcher_configurationv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListLauncherConfigurations transport.Handler[launcher_configurationv1.ListLauncherConfigurationsRequest, launcher_configurationv1.ListLauncherConfigurationsResponse]
	getLauncherConfiguration    transport.Handler[launcher_configurationv1.GetLauncherConfigurationRequest, launcher_configurationv1.GetLauncherConfigurationResponse]
}

func (s *server) ListLauncherConfigurations(ctx context.Context, req *connect.Request[launcher_configurationv1.ListLauncherConfigurationsRequest]) (*connect.Response[launcher_configurationv1.ListLauncherConfigurationsResponse], error) {
	return s.listListLauncherConfigurations(ctx, req)
}

func (s *server) GetLauncherConfiguration(ctx context.Context, req *connect.Request[launcher_configurationv1.GetLauncherConfigurationRequest]) (*connect.Response[launcher_configurationv1.GetLauncherConfigurationResponse], error) {
	return s.getLauncherConfiguration(ctx, req)
}

func NewLauncherConfigurationHandler(svc Service) v1connect.LauncherConfigurationServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListLauncherConfigurations: transport.NewConnectServer(
			eps.listListLauncherConfigurations,
			decodeListLauncherConfigurationsRequest,
			encodeListLauncherConfigurationsResponse,
		),
		getLauncherConfiguration: transport.NewConnectServer(
			eps.getLauncherConfiguration,
			decodeGetLauncherConfigurationRequest,
			encodeGetLauncherConfigurationResponse,
		),
	}
}

func decodeListLauncherConfigurationsRequest(_ context.Context, req *launcher_configurationv1.ListLauncherConfigurationsRequest) (*ListLauncherConfigurationsRequest, error) {
	return &ListLauncherConfigurationsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListLauncherConfigurationsResponse(ctx context.Context, resp *ListLauncherConfigurationsResponse) (*launcher_configurationv1.ListLauncherConfigurationsResponse, error) {
	results := make([]*launcher_configurationv1.LauncherConfiguration, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoLauncherConfiguration(&resp.Results[i])
	}
	return &launcher_configurationv1.ListLauncherConfigurationsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLauncherConfigurationRequest(_ context.Context, req *launcher_configurationv1.GetLauncherConfigurationRequest) (*GetLauncherConfigurationRequest, error) {
	return &GetLauncherConfigurationRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetLauncherConfigurationResponse(ctx context.Context, resp *LauncherConfiguration) (*launcher_configurationv1.GetLauncherConfigurationResponse, error) {
	return &launcher_configurationv1.GetLauncherConfigurationResponse{
		LauncherConfiguration: mapBizToProtoLauncherConfiguration(resp),
	}, nil
}

func mapBizToProtoAgencyDetailed(r *AgencyDetailed) *launcher_configurationv1.AgencyDetailed {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.AgencyDetailed{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		AttemptedLandings: r.AttemptedLandings,
		AttemptedLandingsPayload: r.AttemptedLandingsPayload,
		AttemptedLandingsSpacecraft: r.AttemptedLandingsSpacecraft,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Country: func() []*launcher_configurationv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.Country, len(r.Country))
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
		SocialMediaLinks: func() []*launcher_configurationv1.SocialMediaLink {
			if r.SocialMediaLinks == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.SocialMediaLink, len(r.SocialMediaLinks))
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

func mapBizToProtoAgencyMini(r *AgencyMini) *launcher_configurationv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyNormal(r *AgencyNormal) *launcher_configurationv1.AgencyNormal {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.AgencyNormal{
		Abbrev: r.Abbrev,
		Administrator: r.Administrator,
		Country: func() []*launcher_configurationv1.Country {
			if r.Country == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.Country, len(r.Country))
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

func mapBizToProtoAgencyType(r *AgencyType) *launcher_configurationv1.AgencyType {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoCountry(r *Country) *launcher_configurationv1.Country {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.Country{
		Alpha_2Code: r.Alpha2Code,
		Alpha_3Code: r.Alpha3Code,
		Id: r.Id,
		Name: r.Name,
		NationalityName: r.NationalityName,
		NationalityNameComposed: r.NationalityNameComposed,
	}
}

func mapBizToProtoImage(r *Image) *launcher_configurationv1.Image {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*launcher_configurationv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *launcher_configurationv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *launcher_configurationv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *launcher_configurationv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLauncherConfiguration(r *LauncherConfiguration) *launcher_configurationv1.LauncherConfiguration {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.LauncherConfiguration{
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
		Families: func() []*launcher_configurationv1.LauncherConfigFamilyDetailed {
			if r.Families == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.LauncherConfigFamilyDetailed, len(r.Families))
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
		Program: func() []*launcher_configurationv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.ProgramNormal, len(r.Program))
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

func mapBizToProtoLauncherConfigFamilyDetailed(r *LauncherConfigFamilyDetailed) *launcher_configurationv1.LauncherConfigFamilyDetailed {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.LauncherConfigFamilyDetailed{
		Active: r.Active,
		AttemptedLandings: r.AttemptedLandings,
		ConsecutiveSuccessfulLandings: r.ConsecutiveSuccessfulLandings,
		ConsecutiveSuccessfulLaunches: r.ConsecutiveSuccessfulLaunches,
		Description: r.Description,
		FailedLandings: r.FailedLandings,
		FailedLaunches: r.FailedLaunches,
		Id: r.Id,
		MaidenFlight: r.MaidenFlight,
		Manufacturer: func() []*launcher_configurationv1.AgencyDetailed {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.AgencyDetailed, len(r.Manufacturer))
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

func mapBizToProtoLauncherConfigFamilyMini(r *LauncherConfigFamilyMini) *launcher_configurationv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigFamilyNormal(r *LauncherConfigFamilyNormal) *launcher_configurationv1.LauncherConfigFamilyNormal {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.LauncherConfigFamilyNormal{
		Id: r.Id,
		Manufacturer: func() []*launcher_configurationv1.AgencyNormal {
			if r.Manufacturer == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.AgencyNormal, len(r.Manufacturer))
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

func mapBizToProtoMissionPatch(r *MissionPatch) *launcher_configurationv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoProgramNormal(r *ProgramNormal) *launcher_configurationv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.ProgramNormal{
		Agencies: func() []*launcher_configurationv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []*launcher_configurationv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*launcher_configurationv1.MissionPatch, len(r.MissionPatches))
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

func mapBizToProtoProgramType(r *ProgramType) *launcher_configurationv1.ProgramType {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSocialMedia(r *SocialMedia) *launcher_configurationv1.SocialMedia {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.SocialMedia{
		Id: r.Id,
		Logo: mapBizToProtoImage(r.Logo),
		Name: r.Name,
		Url: r.Url,
	}
}

func mapBizToProtoSocialMediaLink(r *SocialMediaLink) *launcher_configurationv1.SocialMediaLink {
	if r == nil {
		return nil
	}
	return &launcher_configurationv1.SocialMediaLink{
		Id: r.Id,
		SocialMedia: mapBizToProtoSocialMedia(r.SocialMedia),
		Url: r.Url,
	}
}

