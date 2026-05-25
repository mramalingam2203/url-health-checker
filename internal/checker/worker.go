package checker

import (
	"context"
	"net/http"
	"sync"
	"time"
)

func Worker(
	ctx context.Context,
	client *http.Client,
	jobs <-chan string,
	results chan<- URLResult,
	wg *sync.WaitGroup,
	maxRetries int,
	retryDelay time.Duration,
	limiter <-chan time.Time,
) {
	defer wg.Done()

	for url := range jobs {

		// Wait for rate-limit token
		<-limiter

		result := CheckURL(
			client,
			url,
			maxRetries,
			retryDelay,
		)

		results <- result
	}
}
