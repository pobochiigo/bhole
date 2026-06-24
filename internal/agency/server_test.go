package agency_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"com.gitlab/pobochiigo/bhole/internal/agency"
	agencyv1 "com.gitlab/pobochiigo/bhole/proto/agency/v1"
	"com.gitlab/pobochiigo/bhole/proto/agency/v1/agencyv1connect"
	"connectrpc.com/connect"
)

type mockService struct {
	agency.Service
	listAgenciesFn func(ctx context.Context, req *agency.ListAgenciesRequest) (*agency.ListAgenciesResponse, error)
	getAgencyFn    func(ctx context.Context, req *agency.GetAgencyRequest) (*agency.Agency, error)
}

func (m *mockService) ListAgencies(ctx context.Context, req *agency.ListAgenciesRequest) (*agency.ListAgenciesResponse, error) {
	return m.listAgenciesFn(ctx, req)
}

func (m *mockService) GetAgency(ctx context.Context, req *agency.GetAgencyRequest) (*agency.Agency, error) {
	return m.getAgencyFn(ctx, req)
}

func TestServerHandler(t *testing.T) {
	mockSvc := &mockService{
		listAgenciesFn: func(ctx context.Context, req *agency.ListAgenciesRequest) (*agency.ListAgenciesResponse, error) {
			if req.Limit != 5 {
				t.Errorf("expected limit 5, got %d", req.Limit)
			}
			return &agency.ListAgenciesResponse{
				Count: 1,
				Results: []agency.Agency{
					{
						Id:   42,
						Name: "Test Agency",
					},
				},
			}, nil
		},
		getAgencyFn: func(ctx context.Context, req *agency.GetAgencyRequest) (*agency.Agency, error) {
			if req.ID != 99 {
				t.Errorf("expected ID 99, got %d", req.ID)
			}
			return &agency.Agency{
				Id:   99,
				Name: "Single Agency",
			}, nil
		},
	}

	_, handler := agencyv1connect.NewAgencyServiceHandler(agency.NewAgencyHandler(mockSvc))

	ts := httptest.NewServer(handler)
	defer ts.Close()

	client := agencyv1connect.NewAgencyServiceClient(http.DefaultClient, ts.URL)

	ctx := context.Background()

	// 1. Test ListAgencies
	resp, err := client.ListAgencies(ctx, connect.NewRequest(&agencyv1.ListAgenciesRequest{
		Limit: 5,
	}))
	if err != nil {
		t.Fatalf("ListAgencies failed: %v", err)
	}

	if resp.Msg.Count != 1 {
		t.Errorf("expected count 1, got %d", resp.Msg.Count)
	}
	if len(resp.Msg.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(resp.Msg.Results))
	}
	if resp.Msg.Results[0].Name != "Test Agency" || resp.Msg.Results[0].Id != 42 {
		t.Errorf("unexpected agency: %+v", resp.Msg.Results[0])
	}

	// 2. Test GetAgency
	getResp, err := client.GetAgency(ctx, connect.NewRequest(&agencyv1.GetAgencyRequest{
		Id: 99,
	}))
	if err != nil {
		t.Fatalf("GetAgency failed: %v", err)
	}

	if getResp.Msg.Agency.Name != "Single Agency" || getResp.Msg.Agency.Id != 99 {
		t.Errorf("unexpected agency: %+v", getResp.Msg.Agency)
	}
}
