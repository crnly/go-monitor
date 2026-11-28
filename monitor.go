package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const checkInterval = 5 * time.Second

type checkSample struct {
	success int
	elapsed time.Duration
}

func serveMonitor(client *http.Client, target string) error {
	// Publish both measurements together while Prometheus reads them.
	var latest atomic.Pointer[checkSample]
	go func() {
		ticker := time.NewTicker(checkInterval)
		defer ticker.Stop()

		for {
			code, elapsed, err := checkEndpoint(client, target)
			sample := checkSample{elapsed: elapsed}
			if isHealthy(code) {
				sample.success = 1
			}
			if err != nil {
				log.Printf("Endpoint check failed: %v", err)
			}
			latest.Store(&sample)
			<-ticker.C
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		if sample := latest.Load(); sample != nil {
			fmt.Fprintln(w, "# HELP endpoint_check_success Whether the last HTTP check returned a 2xx status.")
			fmt.Fprintln(w, "# TYPE endpoint_check_success gauge")
			fmt.Fprintf(w, "endpoint_check_success %d\n", sample.success)
			fmt.Fprintln(w, "# HELP endpoint_check_duration_seconds Elapsed time of the last HTTP check.")
			fmt.Fprintln(w, "# TYPE endpoint_check_duration_seconds gauge")
			fmt.Fprintf(w, "endpoint_check_duration_seconds %g\n", sample.elapsed.Seconds())
		}
	})

	server := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	return server.ListenAndServe()
}
