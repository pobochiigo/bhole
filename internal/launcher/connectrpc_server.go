package launcher

import (
	"context"

	"com.gitlab/pobochiigo/bhole/internal/transport"
	launcherv1 "com.gitlab/pobochiigo/bhole/proto/launcher/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/launcher/v1/launcherv1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListLaunchers transport.Handler[launcherv1.ListLaunchersRequest, launcherv1.ListLaunchersResponse]
	getLauncher    transport.Handler[launcherv1.GetLauncherRequest, launcherv1.GetLauncherResponse]
}

func (s *server) ListLaunchers(ctx context.Context, req *connect.Request[launcherv1.ListLaunchersRequest]) (*connect.Response[launcherv1.ListLaunchersResponse], error) {
	return s.listListLaunchers(ctx, req)
}

func (s *server) GetLauncher(ctx context.Context, req *connect.Request[launcherv1.GetLauncherRequest]) (*connect.Response[launcherv1.GetLauncherResponse], error) {
	return s.getLauncher(ctx, req)
}

func NewLauncherHandler(svc Service) v1connect.LauncherServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListLaunchers: transport.NewConnectServer(
			eps.listListLaunchers,
			decodeListLaunchersRequest,
			encodeListLaunchersResponse,
		),
		getLauncher: transport.NewConnectServer(
			eps.getLauncher,
			decodeGetLauncherRequest,
			encodeGetLauncherResponse,
		),
	}
}

func decodeListLaunchersRequest(_ context.Context, req *launcherv1.ListLaunchersRequest) (*ListLaunchersRequest, error) {
	return &ListLaunchersRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListLaunchersResponse(ctx context.Context, resp *ListLaunchersResponse) (*launcherv1.ListLaunchersResponse, error) {
	results := make([]*launcherv1.Launcher, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoLauncher(&resp.Results[i])
	}
	return &launcherv1.ListLaunchersResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLauncherRequest(_ context.Context, req *launcherv1.GetLauncherRequest) (*GetLauncherRequest, error) {
	return &GetLauncherRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetLauncherResponse(ctx context.Context, resp *Launcher) (*launcherv1.GetLauncherResponse, error) {
	return &launcherv1.GetLauncherResponse{
		Launcher: mapBizToProtoLauncher(resp),
	}, nil
}

func mapBizToProtoImage(r *Image) *launcherv1.Image {
	if r == nil {
		return nil
	}
	return &launcherv1.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapBizToProtoImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []*launcherv1.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]*launcherv1.ImageVariant, len(r.Variants))
			for i := range r.Variants {
				res[i] = mapBizToProtoImageVariant(&r.Variants[i])
			}
			return res
		}(),
	}
}

func mapBizToProtoImageLicense(r *ImageLicense) *launcherv1.ImageLicense {
	if r == nil {
		return nil
	}
	return &launcherv1.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapBizToProtoImageVariant(r *ImageVariant) *launcherv1.ImageVariant {
	if r == nil {
		return nil
	}
	return &launcherv1.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		Type: mapBizToProtoImageVariantType(r.TypeVal),
	}
}

func mapBizToProtoImageVariantType(r *ImageVariantType) *launcherv1.ImageVariantType {
	if r == nil {
		return nil
	}
	return &launcherv1.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapBizToProtoLauncherConfigFamilyMini(r *LauncherConfigFamilyMini) *launcherv1.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &launcherv1.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapBizToProtoLauncherConfigList(r *LauncherConfigList) *launcherv1.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &launcherv1.LauncherConfigList{
		Families: func() []*launcherv1.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]*launcherv1.LauncherConfigFamilyMini, len(r.Families))
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

func mapBizToProtoLauncher(r *Launcher) *launcherv1.Launcher {
	if r == nil {
		return nil
	}
	return &launcherv1.Launcher{
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
		LauncherConfig: mapBizToProtoLauncherConfigList(r.LauncherConfig),
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		Status: mapBizToProtoLauncherStatus(r.Status),
		SuccessfulLandings: r.SuccessfulLandings,
		Url: r.Url,
	}
}

func mapBizToProtoLauncherStatus(r *LauncherStatus) *launcherv1.LauncherStatus {
	if r == nil {
		return nil
	}
	return &launcherv1.LauncherStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

