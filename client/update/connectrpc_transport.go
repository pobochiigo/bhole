package update

import (
	"context"

	"github.com/pobochiigo/bhole/client/transport"
	bizupdate "github.com/pobochiigo/bhole/internal/update"
	updatev1 "github.com/pobochiigo/bhole/proto/update/v1"
	v1connect "github.com/pobochiigo/bhole/proto/update/v1/updatev1connect"
	"connectrpc.com/connect"
)

func NewUpdateClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) bizupdate.Service {
	connectClient := v1connect.NewUpdateServiceClient(httpClient, baseURL, opts...)

	return &endpoints{
		listListUpdates: transport.NewConnectClient(
			connectClient.ListUpdates,
			encodeListUpdatesRequest,
			decodeListUpdatesResponse,
		),
		getUpdate: transport.NewConnectClient(
			connectClient.GetUpdate,
			encodeGetUpdateRequest,
			decodeGetUpdateResponse,
		),
	}
}

func encodeListUpdatesRequest(_ context.Context, req *bizupdate.ListUpdatesRequest) (*updatev1.ListUpdatesRequest, error) {
	return &updatev1.ListUpdatesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeGetUpdateRequest(_ context.Context, req *bizupdate.GetUpdateRequest) (*updatev1.GetUpdateRequest, error) {
	return &updatev1.GetUpdateRequest{
		Id:   req.ID,
		Mode: req.Mode,
	}, nil
}

func decodeListUpdatesResponse(ctx context.Context, resp *updatev1.ListUpdatesResponse) (*bizupdate.ListUpdatesResponse, error) {
	results := make([]bizupdate.Update, len(resp.Results))
	for i, r := range resp.Results {
		results[i] = *mapProtoToBizUpdate(r)
	}
	return &bizupdate.ListUpdatesResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetUpdateResponse(ctx context.Context, resp *updatev1.GetUpdateResponse) (*bizupdate.Update, error) {
	if resp.Update == nil {
		return nil, nil
	}
	return mapProtoToBizUpdate(resp.Update), nil
}

func mapProtoToBizUpdate(r *updatev1.Update) *bizupdate.Update {
	if r == nil {
		return nil
	}
	return &bizupdate.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
}

