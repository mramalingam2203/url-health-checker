package main

import (
	"net/http"
	"sync"
	"time"

	"healthchecker/internal/checker"
	"healthchecker/internal/output"
	"healthchecker/internal/reader"
)

func main() {
	urls, err := reader.ReadURLs("urls.txt")
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	results := make(chan checker.URLResult)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			result := checker.CheckURL(client, u)

			results <- result

		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		output.PrintResult(result)
	}
}
