package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"time"
)

const (
	defaulTimeout = 100
)

var (
	healthcheckURL   string
	healthcheckAgent = "Docker_Healthcheck_Agent"
)

func main() {
	timeout := flag.Int("t", defaulTimeout, "the request timeout in milliseconds")
	flag.Parse()
	os.Exit(healthcheck(healthcheckURL, healthcheckAgent, time.Duration(*timeout)*time.Millisecond))
}

// healthcheck return an integer that will be used as the exit code.
//
// See https://docs.docker.com/engine/reference/builder/#healthcheck
// The command’s exit status indicates the health status of the container. The possible values are:
//
//	0: success   - the container is healthy and ready for use
//	1: unhealthy - the container is not working correctly
//	2: reserved  - do not use this exit code
func healthcheck(url, useragent string, timeout time.Duration) int {
	client := &http.Client{ //nolint: exhaustruct
		Timeout: timeout,
	}

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return 1
	}

	req.Header.Set("User-Agent", useragent)
	resp, err := client.Do(req)

	if err != nil {
		return 1
	}

	err = resp.Body.Close()
	if err != nil {
		return 1
	}

	if resp.StatusCode != http.StatusOK {
		return 1
	}

	return 0
}
