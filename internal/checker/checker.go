package checker

import (
	"net/http"
	"time"
)

func CheckURL(client *http.Client, url string) URLResult {
	start := time.Now()

	resp, err := client.Get(url)

	latency := time.Since(start)

	if err != nil {
		return URLResult{
			URL:     url,
			Latency: latency,
			Success: false,
			Error:   err.Error(),
		}
	}

	defer resp.Body.Close()

	return URLResult{
		URL:        url,
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Latency:    latency,
		Success:    true,
	}
}
