package checker

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestWorkerProcessesJobs(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)

		}),
	)

	defer server.Close()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	jobs := make(chan string, 1)
	results := make(chan URLResult, 1)

	var wg sync.WaitGroup

	ticker := time.NewTicker(1 * time.Millisecond)

	defer ticker.Stop()

	wg.Add(1)

	go Worker(
		context.Background(),
		client,
		jobs,
		results,
		&wg,
		1,
		1*time.Second,
		ticker.C,
	)

	jobs <- server.URL

	close(jobs)

	wg.Wait()

	result := <-results

	if !result.Success {
		t.Errorf("expected success")
	}
}
