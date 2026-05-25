package checker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheckURLSuccess(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)

		}),
	)

	defer server.Close()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	result := CheckURL(
		client,
		server.URL,
		3,
		1*time.Second,
	)

	if !result.Success {
		t.Errorf("expected success")
	}

	if result.StatusCode != 200 {
		t.Errorf(
			"expected 200, got %d",
			result.StatusCode,
		)
	}
}

func TestCheckURLFailure(t *testing.T) {

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	result := CheckURL(
		client,
		"http://invalid-url",
		2,
		1*time.Second,
	)

	if result.Success {
		t.Errorf("expected failure")
	}
}
