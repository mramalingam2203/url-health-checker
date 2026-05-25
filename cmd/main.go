package main

import (
	"flag"
	"net/http"
	"sync"
	"time"

	"healthchecker/internal/checker"
	"healthchecker/internal/output"
	"healthchecker/internal/reader"
)

func main() {

	workers := flag.Int(
		"workers",
		5,
		"number of concurrent workers",
	)

	inputFile := flag.String(
		"input",
		"urls.txt",
		"path to URL input file",
	)

	timeout := flag.Int(
		"timeout",
		5,
		"HTTP timeout in seconds",
	)

	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	retries := flag.Int(
		"retries",
		3,
		"maximum retry attempts",
	)

	retryDelay := flag.Int(
		"retry-delay",
		1,
		"retry delay in seconds",
	)

	rateLimit := flag.Int(
		"rate",
		5,
		"maximum requests per second",
	)

	flag.Parse()

	interval := time.Second / time.Duration(*rateLimit)

	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	urls, err := reader.ReadURLs(*inputFile)
	if err != nil {
		panic(err)
	}

	jobs := make(chan string)
	results := make(chan checker.URLResult)

	var wg sync.WaitGroup

	workerCount := *workers
	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go checker.Worker(
			client,
			jobs,
			results,
			&wg,
			*retries,
			time.Duration(*retryDelay)*time.Second,
			ticker.C,
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
