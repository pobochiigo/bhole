package transport_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"

	"github.com/pobochiigo/bhole/client/transport"
	launchv1 "github.com/pobochiigo/bhole/proto/launch/v1"
	"github.com/pobochiigo/bhole/proto/launch/v1/launchv1connect"
)

func newLaunchClient(t *testing.T, handler http.HandlerFunc) launchv1connect.LaunchServiceClient {
	t.Helper()
	ts := httptest.NewServer(handler)
	t.Cleanup(ts.Close)
	return launchv1connect.NewLaunchServiceClient(transport.NewRESTClient(ts.URL, nil), ts.URL)
}

func TestRESTClient_MapsUpstreamStatusToConnectCode(t *testing.T) {
	cases := []struct {
		status int
		want   connect.Code
	}{
		{http.StatusBadRequest, connect.CodeInvalidArgument},
		{http.StatusUnauthorized, connect.CodeUnauthenticated},
		{http.StatusForbidden, connect.CodePermissionDenied},
		{http.StatusNotFound, connect.CodeNotFound},
		{http.StatusTooManyRequests, connect.CodeResourceExhausted},
		{http.StatusBadGateway, connect.CodeUnavailable},
		{http.StatusInternalServerError, connect.CodeInternal},
		{http.StatusTeapot, connect.CodeUnknown},
	}
	for _, tc := range cases {
		t.Run(http.StatusText(tc.status), func(t *testing.T) {
			client := newLaunchClient(t, func(w http.ResponseWriter, _ *http.Request) {
				http.Error(w, `{"detail":"upstream says no"}`, tc.status)
			})

			_, err := client.GetLaunch(context.Background(), connect.NewRequest(&launchv1.GetLaunchRequest{Id: "x"}))
			if err == nil {
				t.Fatal("expected an error")
			}
			if got := connect.CodeOf(err); got != tc.want {
				t.Fatalf("status %d: got code %v, want %v (err: %v)", tc.status, got, tc.want, err)
			}
		})
	}
}

func TestRESTClient_GetLaunchEscapesIDAndForwardsAuthorization(t *testing.T) {
	var (
		gotPath  string
		gotQuery string
		gotAuth  string
	)
	client := newLaunchClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.EscapedPath()
		gotQuery = r.URL.RawQuery
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"a b/c","name":"Escaped"}`))
	})

	req := connect.NewRequest(&launchv1.GetLaunchRequest{Id: "a b/c?d"})
	req.Header().Set("Authorization", "Token secret")

	resp, err := client.GetLaunch(context.Background(), req)
	if err != nil {
		t.Fatalf("GetLaunch failed: %v", err)
	}
	if resp.Msg.Launch == nil || resp.Msg.Launch.Name != "Escaped" {
		t.Fatalf("unexpected response: %+v", resp.Msg)
	}
	if want := "/2.3.0/launches/a%20b%2Fc%3Fd/"; gotPath != want {
		t.Errorf("path = %q, want %q", gotPath, want)
	}
	if gotQuery != "" {
		t.Errorf("expected no query string without mode, got %q", gotQuery)
	}
	if gotAuth != "Token secret" {
		t.Errorf("Authorization = %q, want forwarded token", gotAuth)
	}
}

func TestRESTClient_PropagatesCallerContext(t *testing.T) {
	upstreamDone := make(chan struct{})
	client := newLaunchClient(t, func(w http.ResponseWriter, r *http.Request) {
		defer close(upstreamDone)
		select {
		case <-r.Context().Done():
			// The caller's deadline reached the upstream request: this is what we want.
		case <-time.After(5 * time.Second):
			t.Error("upstream request was not cancelled with the caller's context")
		}
	})

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err := client.ListLaunches(ctx, connect.NewRequest(&launchv1.ListLaunchesRequest{}))
	if got := connect.CodeOf(err); got != connect.CodeDeadlineExceeded {
		t.Fatalf("got code %v (err: %v), want %v", got, err, connect.CodeDeadlineExceeded)
	}

	select {
	case <-upstreamDone:
	case <-time.After(5 * time.Second):
		t.Fatal("upstream handler never finished")
	}
}
