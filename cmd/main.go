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

	jobs := make(chan string)
	results := make(chan checker.URLResult)

	var wg sync.WaitGroup

	workerCount := 5

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go checker.Worker(
			client,
			jobs,
			results,
			&wg,
		)
	}

	// Send jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}

		close(jobs)
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		output.PrintResult(result)
	}
}
