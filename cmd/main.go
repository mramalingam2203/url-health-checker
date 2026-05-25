package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"healthchecker/internal/checker"
	"healthchecker/internal/output"
	"healthchecker/internal/reader"
)

func main() {

	// =========================
	// CLI FLAGS
	// =========================

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

	watch := flag.Int(
		"watch",
		0,
		"repeat checks every N seconds (0 = run once)",
	)

	flag.Parse()

	// =========================
	// READ URLS
	// =========================

	urls, err := reader.ReadURLs(*inputFile)
	if err != nil {
		panic(err)
	}

	// =========================
	// HTTP CLIENT
	// =========================

	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	// =========================
	// RATE LIMITER
	// =========================

	interval := time.Second / time.Duration(*rateLimit)

	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	// =========================
	// SINGLE MONITORING CYCLE
	// =========================

	runCycle := func() {

		fmt.Println("===================================")
		fmt.Println("Starting Health Check Cycle")
		fmt.Println("Time:", time.Now().Format(time.RFC3339))
		fmt.Println("===================================")

		results := make(chan checker.URLResult)

		// Run checks concurrently
		go func() {

			checker.RunChecks(
				client,
				urls,
				*workers,
				*retries,
				time.Duration(*retryDelay)*time.Second,
				ticker.C,
				results,
			)

			close(results)

		}()

		// Collect results
		for result := range results {
			output.PrintResult(result)
		}

		fmt.Println("Health Check Cycle Completed")
		fmt.Println()
	}

	// =========================
	// RUN ONCE OR WATCH MODE
	// =========================

	if *watch == 0 {

		// Run once
		runCycle()

	} else {

		// Continuous monitoring
		for {

			runCycle()

			fmt.Printf(
				"Sleeping for %d seconds...\n\n",
				*watch,
			)

			time.Sleep(
				time.Duration(*watch) * time.Second,
			)
		}
	}
}
