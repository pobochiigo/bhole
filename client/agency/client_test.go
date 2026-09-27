package agency_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"

	agencyclient "github.com/pobochiigo/bhole/client/agency"
	"github.com/pobochiigo/bhole/pkg/agency"
	agencyv1 "github.com/pobochiigo/bhole/proto/agency/v1"
	"github.com/pobochiigo/bhole/proto/agency/v1/agencyv1connect"
)

// emptyHandler answers GetAgency with a response that carries no agency, the
// shape a misbehaving upstream or an unimplemented stub would produce.
type emptyHandler struct {
	agencyv1connect.UnimplementedAgencyServiceHandler
}

func (emptyHandler) GetAgency(context.Context, *connect.Request[agencyv1.GetAgencyRequest]) (*connect.Response[agencyv1.GetAgencyResponse], error) {
	return connect.NewResponse(&agencyv1.GetAgencyResponse{}), nil
}

func (emptyHandler) ListAgencies(context.Context, *connect.Request[agencyv1.ListAgenciesRequest]) (*connect.Response[agencyv1.ListAgenciesResponse], error) {
	return connect.NewResponse(&agencyv1.ListAgenciesResponse{
		Count:   1,
		Results: []*agencyv1.Agency{{Id: 7, Name: "Seven"}},
	}), nil
}

func newClient(t *testing.T) agency.Service {
	t.Helper()
	mux := http.NewServeMux()
	mux.Handle(agencyv1connect.NewAgencyServiceHandler(emptyHandler{}))
	ts := httptest.NewServer(mux)
	t.Cleanup(ts.Close)
	return agencyclient.NewAgencyClient(http.DefaultClient, ts.URL)
}

func TestGetAgency_EmptyResponseIsNotFound(t *testing.T) {
	svc := newClient(t)

	got, err := svc.GetAgency(context.Background(), &agency.GetAgencyRequest{ID: 1})
	if got != nil {
		t.Fatalf("expected nil result, got %+v", got)
	}
	if code := connect.CodeOf(err); code != connect.CodeNotFound {
		t.Fatalf("got code %v (err: %v), want %v", code, err, connect.CodeNotFound)
	}
}

func TestListAgencies_MapsResults(t *testing.T) {
	svc := newClient(t)

	resp, err := svc.ListAgencies(context.Background(), &agency.ListAgenciesRequest{Limit: 1})
	if err != nil {
		t.Fatalf("ListAgencies failed: %v", err)
	}
	if resp.Count != 1 || len(resp.Results) != 1 || resp.Results[0].Id != 7 || resp.Results[0].Name != "Seven" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}
