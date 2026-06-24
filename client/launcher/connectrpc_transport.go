package launcher

import (
	"context"

	"com.gitlab/pobochiigo/bhole/client/transport"
	bizlauncher "com.gitlab/pobochiigo/bhole/internal/launcher"
	launcherv1 "com.gitlab/pobochiigo/bhole/proto/launcher/v1"
	v1connect "com.gitlab/pobochiigo/bhole/proto/launcher/v1/launcherv1connect"
	"connectrpc.com/connect"
)

func NewLauncherClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizlauncher.Service {
	connectClient := v1connect.NewLauncherServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListLaunchers: transport.NewConnectClient(
			connectClient.ListLaunchers,
			encodeListLaunchersRequest,
			decodeListLaunchersResponse,
		),
		getLauncher: transport.NewConnectClient(
			connectClient.GetLauncher,
			encodeGetLauncherRequest,
			decodeGetLauncherResponse,
		),
	}
}

func encodeListLaunchersRequest(_ context.Context, req *bizlauncher.ListLaunchersRequest) (*launcherv1.ListLaunchersRequest, error) {
	return &launcherv1.ListLaunchersRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetLauncherRequest(_ context.Context, req *bizlauncher.GetLauncherRequest) (*launcherv1.GetLauncherRequest, error) {
	return &launcherv1.GetLauncherRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListLaunchersResponse(ctx context.Context, resp *launcherv1.ListLaunchersResponse) (*bizlauncher.ListLaunchersResponse, error) {
	results := make([]bizlauncher.Launcher, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizLauncher(r)
	}
	return &bizlauncher.ListLaunchersResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetLauncherResponse(ctx context.Context, resp *launcherv1.GetLauncherResponse) (*bizlauncher.Launcher, error) {
	if resp.Launcher == nil {
		return nil, nil
	}
	return mapProtoToBizLauncher(resp.Launcher), nil
}

func mapProtoToBizImage(r *launcherv1.Image) *bizlauncher.Image {
	if r == nil {
		return nil
	}
	return &bizlauncher.Image{
		Credit: r.Credit,
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		License: mapProtoToBizImageLicense(r.License),
		Name: r.Name,
		SingleUse: r.SingleUse,
		ThumbnailUrl: r.ThumbnailUrl,
		Variants: func() []bizlauncher.ImageVariant {
			if r.Variants == nil {
				return nil
			}
			res := make([]bizlauncher.ImageVariant, len(r.Variants))
			for i, v := range r.Variants {
				res[i] = *mapProtoToBizImageVariant(v)
			}
			return res
		}(),
	}
}

func mapProtoToBizImageLicense(r *launcherv1.ImageLicense) *bizlauncher.ImageLicense {
	if r == nil {
		return nil
	}
	return &bizlauncher.ImageLicense{
		Id: r.Id,
		Link: r.Link,
		Name: r.Name,
		Priority: r.Priority,
	}
}

func mapProtoToBizImageVariant(r *launcherv1.ImageVariant) *bizlauncher.ImageVariant {
	if r == nil {
		return nil
	}
	return &bizlauncher.ImageVariant{
		Id: r.Id,
		ImageUrl: r.ImageUrl,
		TypeVal: mapProtoToBizImageVariantType(r.Type),
	}
}

func mapProtoToBizImageVariantType(r *launcherv1.ImageVariantType) *bizlauncher.ImageVariantType {
	if r == nil {
		return nil
	}
	return &bizlauncher.ImageVariantType{
		Id: r.Id,
		Name: r.Name,
	}
}

func mapProtoToBizLauncherConfigFamilyMini(r *launcherv1.LauncherConfigFamilyMini) *bizlauncher.LauncherConfigFamilyMini {
	if r == nil {
		return nil
	}
	return &bizlauncher.LauncherConfigFamilyMini{
		Id: r.Id,
		Name: r.Name,
		ResponseMode: r.ResponseMode,
	}
}

func mapProtoToBizLauncherConfigList(r *launcherv1.LauncherConfigList) *bizlauncher.LauncherConfigList {
	if r == nil {
		return nil
	}
	return &bizlauncher.LauncherConfigList{
		Families: func() []bizlauncher.LauncherConfigFamilyMini {
			if r.Families == nil {
				return nil
			}
			res := make([]bizlauncher.LauncherConfigFamilyMini, len(r.Families))
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

func mapProtoToBizLauncher(r *launcherv1.Launcher) *bizlauncher.Launcher {
	if r == nil {
		return nil
	}
	return &bizlauncher.Launcher{
		AttemptedLandings: r.AttemptedLandings,
		Details: r.Details,
		FastestTurnaround: r.FastestTurnaround,
		FirstLaunchDate: r.FirstLaunchDate,
		FlightProven: r.FlightProven,
		Flights: r.Flights,
		Id: r.Id,
		Image: mapProtoToBizImage(r.Image),
		IsPlaceholder: r.IsPlaceholder,
		LastLaunchDate: r.LastLaunchDate,
		LauncherConfig: mapProtoToBizLauncherConfigList(r.LauncherConfig),
		ResponseMode: r.ResponseMode,
		SerialNumber: r.SerialNumber,
		Status: mapProtoToBizLauncherStatus(r.Status),
		SuccessfulLandings: r.SuccessfulLandings,
		Url: r.Url,
	}
}

func mapProtoToBizLauncherStatus(r *launcherv1.LauncherStatus) *bizlauncher.LauncherStatus {
	if r == nil {
		return nil
	}
	return &bizlauncher.LauncherStatus{
		Id: r.Id,
		Name: r.Name,
	}
}

