package main

import (
	"context"
	"flag"
	"fmt"
	"healthchecker/internal/checker"
	"healthchecker/internal/output"
	"healthchecker/internal/reader"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		os.Interrupt,
		syscall.SIGTERM,
	)

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
	// FORMAT
	// =========================

	format := flag.String(
		"format",
		"console",
		"output format: console|json|csv",
	)

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
				ctx,
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

		var collectedResults []checker.URLResult

		for result := range results {

			collectedResults = append(
				collectedResults,
				result,
			)

			switch *format {

			case "json":
				output.PrintJSON(result)

			case "csv":
				output.PrintCSV(result)

			default:
				output.PrintResult(result)
			}
		}

		// =========================
		// METRICS SUMMARY
		// =========================

		metrics := checker.CalculateMetrics(
			collectedResults,
		)

		output.PrintMetrics(metrics)
	}

	go func() {

		<-signalChan

		fmt.Println()
		fmt.Println("Shutdown signal received...")
		fmt.Println("Stopping gracefully...")

		cancel()

	}()

	// =========================
	// RUN ONCE OR WATCH MODE
	// =========================

	if *watch == 0 {

		// Run once
		runCycle()

	} else {

		// Continuous monitoring
		for {

			select {

			case <-ctx.Done():

				fmt.Println("Monitoring stopped.")

				return

			default:

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
}
