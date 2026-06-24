package transport_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"com.gitlab/pobochiigo/bhole/client/transport"
	launchv1 "com.gitlab/pobochiigo/bhole/proto/launch/v1"
	"com.gitlab/pobochiigo/bhole/proto/launch/v1/launchv1connect"
	"connectrpc.com/connect"
)

func TestRESTClient_ListLaunches(t *testing.T) {
	// Create mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/2.3.0/launches/" {
			t.Errorf("expected path /2.3.0/launches/, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("limit") != "2" {
			t.Errorf("expected limit=2, got %s", r.URL.Query().Get("limit"))
		}

		resp := transport.ListLaunchesResponseJSON{
			Count: 1,
			Results: []transport.LaunchDetailedJSON{
				{
					Id:   "test-uuid",
					Name: "Test Launch",
					Net:  "2026-06-17T12:00:00Z",
					Status: &transport.LaunchStatusJSON{
						Id:          1,
						Name:        "Success",
						Abbrev:      "Success",
						Description: "The launch was successful.",
					},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	// Initialize our RESTClient pointing to the mock server
	restClient := transport.NewRESTClient(ts.URL, nil)

	// Create ConnectRPC client
	launchesClient := launchv1connect.NewLaunchServiceClient(restClient, ts.URL)

	// Call the method
	resp, err := launchesClient.ListLaunches(context.Background(), connect.NewRequest(&launchv1.ListLaunchesRequest{
		Limit: 2,
	}))
	if err != nil {
		t.Fatalf("ListLaunches failed: %v", err)
	}

	if resp.Msg.Count != 1 {
		t.Errorf("expected count 1, got %d", resp.Msg.Count)
	}
	if len(resp.Msg.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(resp.Msg.Results))
	}

	l := resp.Msg.Results[0]
	if l.Id != "test-uuid" || l.Name != "Test Launch" {
		t.Errorf("unexpected launch values: %+v", l)
	}
	if l.Status == nil || l.Status.Name != "Success" {
		t.Errorf("unexpected status: %+v", l.Status)
	}
}
