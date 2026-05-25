package checker

import (
	"net/http"
	"time"
)

func CheckURL(
	client *http.Client,
	url string,
	maxRetries int,
	retryDelay time.Duration,
) URLResult {

	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {

		start := time.Now()

		resp, err := client.Get(url)

		latency := time.Since(start)

		if err == nil {
			defer resp.Body.Close()

			return URLResult{
				URL:        url,
				StatusCode: resp.StatusCode,
				Status:     resp.Status,
				Latency:    latency,
				Success:    true,
				Attempts:   attempt,
			}
		}

		lastErr = err

		time.Sleep(retryDelay)
	}

	return URLResult{
		URL:      url,
		Success:  false,
		Error:    lastErr.Error(),
		Attempts: maxRetries,
	}
}
