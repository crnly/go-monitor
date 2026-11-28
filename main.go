package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	switch {
	case len(os.Args) == 2:
		runCheck(client, os.Args[1])
	case len(os.Args) == 3 && os.Args[1] == "-serve":
		log.Fatal(serveMonitor(client, os.Args[2]))
	default:
		fmt.Println("Usage: go run . [-serve] <url>")
	}
}

func runCheck(client *http.Client, url string) {
	statusCode, elapsed, err := checkEndpoint(client, url)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	fmt.Printf("%s: %d %s (%v), healthy=%t\n",
		url, statusCode, http.StatusText(statusCode),
		elapsed.Round(time.Millisecond), isHealthy(statusCode))
}

func checkEndpoint(client *http.Client, url string) (int, time.Duration, error) {
	start := time.Now()
	response, err := client.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		return 0, elapsed, err
	}
	defer response.Body.Close()

	return response.StatusCode, elapsed, nil
}

func isHealthy(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
