package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheckEndpointHTTPStatus(t *testing.T) {
	for _, test := range []struct {
		status  int
		healthy bool
	}{
		{http.StatusOK, true},
		{http.StatusServiceUnavailable, false},
	} {
		t.Run(http.StatusText(test.status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
			}))
			defer server.Close()

			code, _, err := checkEndpoint(server.Client(), server.URL)
			if err != nil {
				t.Fatal(err)
			}
			if code != test.status || isHealthy(code) != test.healthy {
				t.Errorf("got status %d and healthy=%t, want %d and healthy=%t",
					code, isHealthy(code), test.status, test.healthy)
			}
		})
	}
}

func TestCheckEndpointTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	defer server.Close()

	client := server.Client()
	client.Timeout = 100 * time.Millisecond
	code, _, err := checkEndpoint(client, server.URL)
	if code != 0 || !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("got status %d and error %v, want status 0 and a timeout", code, err)
	}
}
