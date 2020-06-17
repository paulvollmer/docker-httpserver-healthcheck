package main

import (
	"net/http"
	"os"
	"time"
)

var (
	healthcheckURL   string
	healthcheckAgent = "Docker_Healthcheck_Agent"
)

func main() {
	os.Exit(healthcheck(healthcheckURL, healthcheckAgent))
}

// healthcheck return an integer that will be used as the exit code.
//
// See https://docs.docker.com/engine/reference/builder/#healthcheck
// The command’s exit status indicates the health status of the container. The possible values are:
//   0: success   - the container is healthy and ready for use
//   1: unhealthy - the container is not working correctly
//   2: reserved  - do not use this exit code
func healthcheck(url, useragent string) int {
	client := &http.Client{
		Timeout: 100 * time.Millisecond,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 1
	}
	req.Header.Set("User-Agent", useragent)
	resp, err := client.Do(req)
	if err != nil {
		return 1
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 1
	}

	return 0
}
