package update

import (
	"context"

	"github.com/pobochiigo/bhole/internal/transport"
	updatev1 "github.com/pobochiigo/bhole/proto/update/v1"
	v1connect "github.com/pobochiigo/bhole/proto/update/v1/updatev1connect"
	"connectrpc.com/connect"
)

type server struct {
	listListUpdates transport.Handler[updatev1.ListUpdatesRequest, updatev1.ListUpdatesResponse]
	getUpdate    transport.Handler[updatev1.GetUpdateRequest, updatev1.GetUpdateResponse]
}

func (s *server) ListUpdates(ctx context.Context, req *connect.Request[updatev1.ListUpdatesRequest]) (*connect.Response[updatev1.ListUpdatesResponse], error) {
	return s.listListUpdates(ctx, req)
}

func (s *server) GetUpdate(ctx context.Context, req *connect.Request[updatev1.GetUpdateRequest]) (*connect.Response[updatev1.GetUpdateResponse], error) {
	return s.getUpdate(ctx, req)
}

func NewUpdateHandler(svc Service) v1connect.UpdateServiceHandler {
	eps := MakeEndpoints(svc)
	return &server{
		listListUpdates: transport.NewConnectServer(
			eps.listListUpdates,
			decodeListUpdatesRequest,
			encodeListUpdatesResponse,
		),
		getUpdate: transport.NewConnectServer(
			eps.getUpdate,
			decodeGetUpdateRequest,
			encodeGetUpdateResponse,
		),
	}
}

func decodeListUpdatesRequest(_ context.Context, req *updatev1.ListUpdatesRequest) (*ListUpdatesRequest, error) {
	return &ListUpdatesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
		Mode:   req.Mode,
	}, nil
}

func encodeListUpdatesResponse(ctx context.Context, resp *ListUpdatesResponse) (*updatev1.ListUpdatesResponse, error) {
	results := make([]*updatev1.Update, len(resp.Results))
	for i := range resp.Results {
		results[i] = mapBizToProtoUpdate(&resp.Results[i])
	}
	return &updatev1.ListUpdatesResponse{
		Count:    resp.Count,
		Next:     resp.Next,
		Previous: resp.Previous,
		Results:  results,
	}, nil
}

func decodeGetUpdateRequest(_ context.Context, req *updatev1.GetUpdateRequest) (*GetUpdateRequest, error) {
	return &GetUpdateRequest{
		ID:   req.Id,
		Mode: req.Mode,
	}, nil
}

func encodeGetUpdateResponse(ctx context.Context, resp *Update) (*updatev1.GetUpdateResponse, error) {
	return &updatev1.GetUpdateResponse{
		Update: mapBizToProtoUpdate(resp),
	}, nil
}

func mapBizToProtoUpdate(r *Update) *updatev1.Update {
	if r == nil {
		return nil
	}
	return &updatev1.Update{
		Comment: r.Comment,
		CreatedBy: r.CreatedBy,
		CreatedOn: r.CreatedOn,
		Id: r.Id,
		InfoUrl: r.InfoUrl,
		ProfileImage: r.ProfileImage,
	}
}

