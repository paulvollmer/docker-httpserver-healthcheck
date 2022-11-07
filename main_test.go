package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthcheck(t *testing.T) {
	t.Parallel()

	userAgent := "Test"
	timeout := 100 * time.Millisecond

	t.Run("exitcode 0", func(t *testing.T) {
		t.Parallel()

		testserver := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/healthcheck" {
				t.Error("request path should be /healthcheck")
			}
			if req.Header.Get("User-Agent") != userAgent {
				t.Error("request User-Agent should be", userAgent)
			}
			res.WriteHeader(http.StatusOK)
			_, _ = res.Write([]byte(`OK`))
		}))
		defer testserver.Close()

		exitcode := healthcheck(testserver.URL+"/healthcheck", userAgent, timeout)
		if exitcode != 0 {
			t.Error("exitcode should be 0")
		}
	})

	t.Run("exitcode 1 statuscode not 200", func(t *testing.T) {
		t.Parallel()

		testserver := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(`Server Error`))
		}))
		defer testserver.Close()

		exitcode := healthcheck(testserver.URL, userAgent, timeout)
		if exitcode != 1 {
			t.Error("exitcode should be 1")
		}
	})

	t.Run("exitcode 1", func(t *testing.T) {
		t.Parallel()

		exitcode := healthcheck("http://localhost:9999999", userAgent, timeout)
		if exitcode != 1 {
			t.Error("exitcode should be 1")
		}
	})
}
