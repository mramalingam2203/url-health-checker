package checker

import (
	"net/http"
	"sync"
)

func Worker(
	client *http.Client,
	jobs <-chan string,
	results chan<- URLResult,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for url := range jobs {
		result := CheckURL(client, url)

		results <- result
	}
}
