package output

import (
	"fmt"

	"healthchecker/internal/checker"
)

func PrintMetrics(metrics checker.Metrics) {

	fmt.Println("===================================")
	fmt.Println("Monitoring Summary")
	fmt.Println("===================================")

	fmt.Printf(
		"Total URLs:        %d\n",
		metrics.TotalURLs,
	)

	fmt.Printf(
		"Successful:        %d\n",
		metrics.Successful,
	)

	fmt.Printf(
		"Failed:            %d\n",
		metrics.Failed,
	)

	fmt.Printf(
		"Success Rate:      %.2f%%\n",
		metrics.SuccessRate,
	)

	fmt.Println()

	fmt.Printf(
		"Average Latency:   %v\n",
		metrics.AverageLatency,
	)

	fmt.Printf(
		"Fastest URL:       %s\n",
		metrics.FastestURL,
	)

	fmt.Printf(
		"Fastest Latency:   %v\n",
		metrics.FastestLatency,
	)

	fmt.Printf(
		"Slowest URL:       %s\n",
		metrics.SlowestURL,
	)

	fmt.Printf(
		"Slowest Latency:   %v\n",
		metrics.SlowestLatency,
	)

	fmt.Println()
}
