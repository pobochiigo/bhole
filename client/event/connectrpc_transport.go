package event

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizevent "com.gitlab/pobochiigo/bhole/internal/event"
	eventv1 "com.gitlab/pobochiigo/bhole/proto/event/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/event/v1/eventv1connect"
	"connectrpc.com/connect"
)

func NewEventClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizevent.Service {
	connectClient := v1connect.NewEventServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListEvents: transport.NewConnectClient(
			connectClient.ListEvents,
			encodeListEventsRequest,
			decodeListEventsResponse,
		),
		getEvent: transport.NewConnectClient(
			connectClient.GetEvent,
			encodeGetEventRequest,
			decodeGetEventResponse,
		),
	}
}

func encodeListEventsRequest(_ context.Context, req *bizevent.ListEventsRequest) (*eventv1.ListEventsRequest, error) {
	return &eventv1.ListEventsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetEventRequest(_ context.Context, req *bizevent.GetEventRequest) (*eventv1.GetEventRequest, error) {
	return &eventv1.GetEventRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListEventsResponse(ctx context.Context, resp *eventv1.ListEventsResponse) (*bizevent.ListEventsResponse, error) {
	results := make([]bizevent.Event, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizEvent(r)
	}
	return &bizevent.ListEventsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetEventResponse(ctx context.Context, resp *eventv1.GetEventResponse) (*bizevent.Event, error) {
	if resp.Event == nil {
		return nil, nil
	}
	return mapProtoToBizEvent(resp.Event), nil
}

func mapProtoToBizAgencyMini(r *eventv1.AgencyMini) *bizevent.AgencyMini {
	if r == nil {
		return nil
	}
	return &bizevent.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		TypeVal: mapProtoToBizAgencyType(r.Type),
		Url: r.Url,
	}
}

func mapProtoToBizAgencyType(r *eventv1.AgencyType) *bizevent.AgencyType {
	if r == nil {
		return nil
	}
	return &bizevent.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizAstronautNormal(r *eventv1.AstronautNormal) *bizevent.AstronautNormal {
	if r == nil {
		return nil
	}
	return &bizevent.AstronautNormal{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Name: r.Name,
		Status: mapProtoToBizAstronautStatus(r.Status),
		Url: r.Url,
	}
}

func mapProtoToBizAstronautStatus(r *eventv1.AstronautStatus) *bizevent.AstronautStatus {
	if r == nil {
		return nil
	}
	return &bizevent.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizEvent(r *eventv1.Event) *bizevent.Event {
	if r == nil {
		return nil
	}
	return &bizevent.Event{
		Agencies: func() []bizevent.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizevent.AgencyMini, len(r.Agencies))
			for i, v := range r.Agencies {
				res[i] = *mapProtoToBizAgencyMini(v)
			}
			return res
		}(),
		Astronauts: func() []bizevent.AstronautNormal {
			if r.Astronauts == nil {
				return nil
			}
			res := make([]bizevent.AstronautNormal, len(r.Astronauts))
			for i, v := range r.Astronauts {
				res[i] = *mapProtoToBizAstronautNormal(v)
			}
			return res
		}(),
		Date: r.Date,
		DatePrecision: mapProtoToBizNetPrecision(r.DatePrecision),
		Description: r.Description,
		Duration: r.Duration,
		Expeditions: func() []bizevent.ExpeditionNormal {
			if r.Expeditions == nil {
				return nil
			}
			res := make([]bizevent.ExpeditionNormal, len(r.Expeditions))
			for i, v := range r.Expeditions {
				res[i] = *mapProtoToBizExpeditionNormal(v)
			}
			return res
		}(),
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		InfoUrls: func() []bizevent.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]bizevent.InfoURL, len(r.InfoUrls))
			for i, v := range r.InfoUrls {
				res[i] = *mapProtoToBizInfoURL(v)
			}
			return res
		}(),
		LastUpdated: r.LastUpdated,
		Launches: func() []bizevent.LaunchBasic {
			if r.Launches == nil {
				return nil
			}
			res := make([]bizevent.LaunchBasic, len(r.Launches))
			for i, v := range r.Launches {
				res[i] = *mapProtoToBizLaunchBasic(v)
			}
			return res
		}(),
		Location: r.Location,
		Name: r.Name,
		Program: func() []bizevent.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]bizevent.ProgramNormal, len(r.Program))
			for i, v := range r.Program {
				res[i] = *mapProtoToBizProgramNormal(v)
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Slug: r.Slug,
		Spacestations: func() []bizevent.SpaceStationNormal {
			if r.Spacestations == nil {
				return nil
			}
			res := make([]bizevent.SpaceStationNormal, len(r.Spacestations))
			for i, v := range r.Spacestations {
				res[i] = *mapProtoToBizSpaceStationNormal(v)
			}
			return res
		}(),
		TypeVal: mapProtoToBizEventType(r.Type),
		Updates: func() []bizevent.Update {
			if r.Updates == nil {
				return nil
			}
			res := make([]bizevent.Update, len(r.Updates))
			for i, v := range r.Updates {
				res[i] = *mapProtoToBizUpdate(v)
			}
			return res
		}(),
		Url: r.Url,
		VidUrls: func() []bizevent.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]bizevent.VidURL, len(r.VidUrls))
			for i, v := range r.VidUrls {
				res[i] = *mapProtoToBizVidURL(v)
			}
			return res
		}(),
		WebcastLive: r.WebcastLive,
	}
}

func mapProtoToBizEventType(r *eventv1.EventType) *bizevent.EventType {
	if r == nil {
		return nil
	}
	return &bizevent.EventType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizExpeditionNormal(r *eventv1.ExpeditionNormal) *bizevent.ExpeditionNormal {
	if r == nil {
		return nil
	}
	return &bizevent.ExpeditionNormal{
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []bizevent.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizevent.MissionPatch, len(r.MissionPatches))
			for i, v := range r.MissionPatches {
				res[i] = *mapProtoToBizMissionPatch(v)
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Spacestation: mapProtoToBizSpaceStationNormal(r.Spacestation),
		Spacewalks: func() []bizevent.SpacewalkList {
			if r.Spacewalks == nil {
				return nil
			}
			res := make([]bizevent.SpacewalkList, len(r.Spacewalks))
			for i, v := range r.Spacewalks {
				res[i] = *mapProtoToBizSpacewalkList(v)
			}
			return res
		}(),
		Start: r.Start,
		Url: r.Url,
	}
}

func mapProtoToBizImage(r *eventv1.Image) *bizevent.Image {
	if r == nil {
		return nil
	}
	return &bizevent.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizevent.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizevent.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *eventv1.ImageLicense) *bizevent.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizevent.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *eventv1.ImageVariant) *bizevent.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizevent.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *eventv1.ImageVariantType) *bizevent.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizevent.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizInfoURL(r *eventv1.InfoURL) *bizevent.InfoURL {
	if r == nil {
		return nil
	}
	return &bizevent.InfoURL{
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

func mapProtoToBizInfoURLType(r *eventv1.InfoURLType) *bizevent.InfoURLType {
	if r == nil {
		return nil
	}
	return &bizevent.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLanguage(r *eventv1.Language) *bizevent.Language {
	if r == nil {
		return nil
	}
	return &bizevent.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLaunchBasic(r *eventv1.LaunchBasic) *bizevent.LaunchBasic {
	if r == nil {
		return nil
	}
	return &bizevent.LaunchBasic{
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapProtoToBizNetPrecision(r.NetPrecision),
		ResponseMode: r.ResponseMode,
		Slug: r.Slug,
		Status: mapProtoToBizLaunchStatus(r.Status),
		Url: r.Url,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
}

func mapProtoToBizLaunchStatus(r *eventv1.LaunchStatus) *bizevent.LaunchStatus {
	if r == nil {
		return nil
	}
	return &bizevent.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizMissionPatch(r *eventv1.MissionPatch) *bizevent.MissionPatch {
	if r == nil {
		return nil
	}
	return &bizevent.MissionPatch{
		Agency: mapProtoToBizAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizNetPrecision(r *eventv1.NetPrecision) *bizevent.NetPrecision {
	if r == nil {
		return nil
	}
	return &bizevent.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizProgramNormal(r *eventv1.ProgramNormal) *bizevent.ProgramNormal {
	if r == nil {
		return nil
	}
	return &bizevent.ProgramNormal{
		Agencies: func() []bizevent.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]bizevent.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []bizevent.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]bizevent.MissionPatch, len(r.MissionPatches))
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

func mapProtoToBizProgramType(r *eventv1.ProgramType) *bizevent.ProgramType {
	if r == nil {
		return nil
	}
	return &bizevent.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpaceStationNormal(r *eventv1.SpaceStationNormal) *bizevent.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &bizevent.SpaceStationNormal{
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

func mapProtoToBizSpaceStationStatus(r *eventv1.SpaceStationStatus) *bizevent.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &bizevent.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpaceStationType(r *eventv1.SpaceStationType) *bizevent.SpaceStationType {
	if r == nil {
		return nil
	}
	return &bizevent.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizSpacewalkList(r *eventv1.SpacewalkList) *bizevent.SpacewalkList {
	if r == nil {
		return nil
	}
	return &bizevent.SpacewalkList{
		Duration: r.Duration,
		End: r.End,
		Id: r.Id,
		Location: r.Location,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Start: r.Start,
		Url: r.Url,
	}
}

func mapProtoToBizUpdate(r *eventv1.Update) *bizevent.Update {
	if r == nil {
		return nil
	}
	return &bizevent.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
}

func mapProtoToBizVidURL(r *eventv1.VidURL) *bizevent.VidURL {
	if r == nil {
		return nil
	}
	return &bizevent.VidURL{
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

func mapProtoToBizVidURLType(r *eventv1.VidURLType) *bizevent.VidURLType {
	if r == nil {
		return nil
	}
	return &bizevent.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

