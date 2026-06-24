package event

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	eventv1 "com.gitlab/pobochiigo/bhole/proto/event/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/event/v1/eventv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListEvents transport.Handler[eventv1.ListEventsRequest, eventv1.ListEventsResponse]
	getEvent    transport.Handler[eventv1.GetEventRequest, eventv1.GetEventResponse]
}

func (s *server) ListEvents(ctx context.Context, req *connect.Request[eventv1.ListEventsRequest]) (*connect.Response[eventv1.ListEventsResponse], error) {
	return s.listListEvents(ctx, req)
}

func (s *server) GetEvent(ctx context.Context, req *connect.Request[eventv1.GetEventRequest]) (*connect.Response[eventv1.GetEventResponse], error) {
	return s.getEvent(ctx, req)
}

func NewEventHandler(svc Service) v1connect.EventServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListEvents: transport.NewConnectServer(
			eps.listListEvents,
			decodeListEventsRequest,
			encodeListEventsResponse,
		),
		getEvent: transport.NewConnectServer(
			eps.getEvent,
			decodeGetEventRequest,
			encodeGetEventResponse,
		),
	}
}

func decodeListEventsRequest(_ context.Context, req *eventv1.ListEventsRequest) (*ListEventsRequest, error) {
	return &ListEventsRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListEventsResponse(ctx context.Context, resp *ListEventsResponse) (*eventv1.ListEventsResponse, error) {
	results := make([]*eventv1.Event, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoEvent(&resp.Results[i])
	}
	return &eventv1.ListEventsResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetEventRequest(_ context.Context, req *eventv1.GetEventRequest) (*GetEventRequest, error) {
	return &GetEventRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetEventResponse(ctx context.Context, resp *Event) (*eventv1.GetEventResponse, error) {
	return &eventv1.GetEventResponse{
		Event: mapBizToProtoEvent(resp),
	}, nil
}

func mapBizToProtoAgencyMini(r *AgencyMini) *eventv1.AgencyMini {
	if r == nil {
		return nil
	}
	return &eventv1.AgencyMini{
		Abbrev: r.Abbrev,
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Type: mapBizToProtoAgencyType(r.TypeVal),
		Url: r.Url,
	}
}

func mapBizToProtoAgencyType(r *AgencyType) *eventv1.AgencyType {
	if r == nil {
		return nil
	}
	return &eventv1.AgencyType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoAstronautNormal(r *AstronautNormal) *eventv1.AstronautNormal {
	if r == nil {
		return nil
	}
	return &eventv1.AstronautNormal{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Name: r.Name,
		Status: mapBizToProtoAstronautStatus(r.Status),
		Url: r.Url,
	}
}

func mapBizToProtoAstronautStatus(r *AstronautStatus) *eventv1.AstronautStatus {
	if r == nil {
		return nil
	}
	return &eventv1.AstronautStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoEvent(r *Event) *eventv1.Event {
	if r == nil {
		return nil
	}
	return &eventv1.Event{
		Agencies: func() []*eventv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*eventv1.AgencyMini, len(r.Agencies))
			for i := range r.Agencies {
				res[i] = mapBizToProtoAgencyMini(&r.Agencies[i])
			}
			return res
		}(),
		Astronauts: func() []*eventv1.AstronautNormal {
			if r.Astronauts == nil {
				return nil
			}
			res := make([]*eventv1.AstronautNormal, len(r.Astronauts))
			for i := range r.Astronauts {
				res[i] = mapBizToProtoAstronautNormal(&r.Astronauts[i])
			}
			return res
		}(),
		Date: r.Date,
		DatePrecision: mapBizToProtoNetPrecision(r.DatePrecision),
		Description: r.Description,
		Duration: r.Duration,
		Expeditions: func() []*eventv1.ExpeditionNormal {
			if r.Expeditions == nil {
				return nil
			}
			res := make([]*eventv1.ExpeditionNormal, len(r.Expeditions))
			for i := range r.Expeditions {
				res[i] = mapBizToProtoExpeditionNormal(&r.Expeditions[i])
			}
			return res
		}(),
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		InfoUrls: func() []*eventv1.InfoURL {
			if r.InfoUrls == nil {
				return nil
			}
			res := make([]*eventv1.InfoURL, len(r.InfoUrls))
			for i := range r.InfoUrls {
				res[i] = mapBizToProtoInfoURL(&r.InfoUrls[i])
			}
			return res
		}(),
		LastUpdated: r.LastUpdated,
		Launches: func() []*eventv1.LaunchBasic {
			if r.Launches == nil {
				return nil
			}
			res := make([]*eventv1.LaunchBasic, len(r.Launches))
			for i := range r.Launches {
				res[i] = mapBizToProtoLaunchBasic(&r.Launches[i])
			}
			return res
		}(),
		Location: r.Location,
		Name: r.Name,
		Program: func() []*eventv1.ProgramNormal {
			if r.Program == nil {
				return nil
			}
			res := make([]*eventv1.ProgramNormal, len(r.Program))
			for i := range r.Program {
				res[i] = mapBizToProtoProgramNormal(&r.Program[i])
			}
			return res
		}(),
		ResponseMode: r.ResponseMode,
		Slug: r.Slug,
		Spacestations: func() []*eventv1.SpaceStationNormal {
			if r.Spacestations == nil {
				return nil
			}
			res := make([]*eventv1.SpaceStationNormal, len(r.Spacestations))
			for i := range r.Spacestations {
				res[i] = mapBizToProtoSpaceStationNormal(&r.Spacestations[i])
			}
			return res
		}(),
		Type: mapBizToProtoEventType(r.TypeVal),
		Updates: func() []*eventv1.Update {
			if r.Updates == nil {
				return nil
			}
			res := make([]*eventv1.Update, len(r.Updates))
			for i := range r.Updates {
				res[i] = mapBizToProtoUpdate(&r.Updates[i])
			}
			return res
		}(),
		Url: r.Url,
		VidUrls: func() []*eventv1.VidURL {
			if r.VidUrls == nil {
				return nil
			}
			res := make([]*eventv1.VidURL, len(r.VidUrls))
			for i := range r.VidUrls {
				res[i] = mapBizToProtoVidURL(&r.VidUrls[i])
			}
			return res
		}(),
		WebcastLive: r.WebcastLive,
	}
}

func mapBizToProtoEventType(r *EventType) *eventv1.EventType {
	if r == nil {
		return nil
	}
	return &eventv1.EventType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoExpeditionNormal(r *ExpeditionNormal) *eventv1.ExpeditionNormal {
	if r == nil {
		return nil
	}
	return &eventv1.ExpeditionNormal{
		End: r.End,
		Id: r.Id,
		MissionPatches: func() []*eventv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*eventv1.MissionPatch, len(r.MissionPatches))
			for i := range r.MissionPatches {
				res[i] = mapBizToProtoMissionPatch(&r.MissionPatches[i])
			}
			return res
		}(),
		Name: r.Name,
		ResponseMode: r.ResponseMode,
		Spacestation: mapBizToProtoSpaceStationNormal(r.Spacestation),
		Spacewalks: func() []*eventv1.SpacewalkList {
			if r.Spacewalks == nil {
				return nil
			}
			res := make([]*eventv1.SpacewalkList, len(r.Spacewalks))
			for i := range r.Spacewalks {
				res[i] = mapBizToProtoSpacewalkList(&r.Spacewalks[i])
			}
			return res
		}(),
		Start: r.Start,
		Url: r.Url,
	}
}

func mapBizToProtoImage(r *Image) *eventv1.Image {
	if r == nil {
		return nil
	}
	return &eventv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*eventv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*eventv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *eventv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &eventv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *eventv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &eventv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *eventv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &eventv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoInfoURL(r *InfoURL) *eventv1.InfoURL {
	if r == nil {
		return nil
	}
	return &eventv1.InfoURL{
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

func mapBizToProtoInfoURLType(r *InfoURLType) *eventv1.InfoURLType {
	if r == nil {
		return nil
	}
	return &eventv1.InfoURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLanguage(r *Language) *eventv1.Language {
	if r == nil {
		return nil
	}
	return &eventv1.Language{
		Code: r.Code,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLaunchBasic(r *LaunchBasic) *eventv1.LaunchBasic {
	if r == nil {
		return nil
	}
	return &eventv1.LaunchBasic{
		Id: r.Id,
		Image: mapBizToProtoImage(r.Image),
		Infographic: r.Infographic,
		LastUpdated: r.LastUpdated,
		LaunchDesignator: r.LaunchDesignator,
		Name: r.Name,
		Net: r.Net,
		NetPrecision: mapBizToProtoNetPrecision(r.NetPrecision),
		ResponseMode: r.ResponseMode,
		Slug: r.Slug,
		Status: mapBizToProtoLaunchStatus(r.Status),
		Url: r.Url,
		WindowEnd: r.WindowEnd,
		WindowStart: r.WindowStart,
	}
}

func mapBizToProtoLaunchStatus(r *LaunchStatus) *eventv1.LaunchStatus {
	if r == nil {
		return nil
	}
	return &eventv1.LaunchStatus{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoMissionPatch(r *MissionPatch) *eventv1.MissionPatch {
	if r == nil {
		return nil
	}
	return &eventv1.MissionPatch{
		Agency: mapBizToProtoAgencyMini(r.Agency),
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Name: r.Name,
		Priority: r.Priority,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoNetPrecision(r *NetPrecision) *eventv1.NetPrecision {
	if r == nil {
		return nil
	}
	return &eventv1.NetPrecision{
		Abbrev: r.Abbrev,
		Description: r.Description,
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoProgramNormal(r *ProgramNormal) *eventv1.ProgramNormal {
	if r == nil {
		return nil
	}
	return &eventv1.ProgramNormal{
		Agencies: func() []*eventv1.AgencyMini {
			if r.Agencies == nil {
				return nil
			}
			res := make([]*eventv1.AgencyMini, len(r.Agencies))
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
		MissionPatches: func() []*eventv1.MissionPatch {
			if r.MissionPatches == nil {
				return nil
			}
			res := make([]*eventv1.MissionPatch, len(r.MissionPatches))
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

func mapBizToProtoProgramType(r *ProgramType) *eventv1.ProgramType {
	if r == nil {
		return nil
	}
	return &eventv1.ProgramType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpaceStationNormal(r *SpaceStationNormal) *eventv1.SpaceStationNormal {
	if r == nil {
		return nil
	}
	return &eventv1.SpaceStationNormal{
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

func mapBizToProtoSpaceStationStatus(r *SpaceStationStatus) *eventv1.SpaceStationStatus {
	if r == nil {
		return nil
	}
	return &eventv1.SpaceStationStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpaceStationType(r *SpaceStationType) *eventv1.SpaceStationType {
	if r == nil {
		return nil
	}
	return &eventv1.SpaceStationType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoSpacewalkList(r *SpacewalkList) *eventv1.SpacewalkList {
	if r == nil {
		return nil
	}
	return &eventv1.SpacewalkList{
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

func mapBizToProtoUpdate(r *Update) *eventv1.Update {
	if r == nil {
		return nil
	}
	return &eventv1.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
}

func mapBizToProtoVidURL(r *VidURL) *eventv1.VidURL {
	if r == nil {
		return nil
	}
	return &eventv1.VidURL{
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

func mapBizToProtoVidURLType(r *VidURLType) *eventv1.VidURLType {
	if r == nil {
		return nil
	}
	return &eventv1.VidURLType{
		Id: r.Id,
		Name: r.Name,
	}
}

