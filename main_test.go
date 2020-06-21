package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthcheck(t *testing.T) {
	userAgent := "Test"

	t.Run("exitcode 0", func(t *testing.T) {
		testserver := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/healthcheck" {
				t.Error("request path should be /healthcheck")
			}
			if req.Header.Get("User-Agent") != userAgent {
				t.Error("request User-Agent should be", userAgent)
			}
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`OK`))
		}))
		defer testserver.Close()

		exitcode := healthcheck(testserver.URL+"/healthcheck", userAgent, 100*time.Millisecond)
		if exitcode != 0 {
			t.Error("exitcode should be 0")
		}
	})

	t.Run("exitcode 1", func(t *testing.T) {
		exitcode := healthcheck("http://localhost:9999999", userAgent, 100*time.Millisecond)
		if exitcode != 1 {
			t.Error("exitcode should be 1")
		}
	})
}
