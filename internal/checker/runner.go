package checker

import (
	"context"
	"net/http"
	"sync"
	"time"
)

func RunChecks(
	ctx context.Context,
	client *http.Client,
	urls []string,
	workerCount int,
	maxRetries int,
	retryDelay time.Duration,
	limiter <-chan time.Time,
	results chan<- URLResult,
) {

	jobs := make(chan string)

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < workerCount; i++ {

		wg.Add(1)

		go Worker(
			ctx,
			client,
			jobs,
			results,
			&wg,
			maxRetries,
			retryDelay,
			limiter,
		)
	}

	// Send jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}

		close(jobs)
	}()

	// Wait for workers
	wg.Wait()
}
