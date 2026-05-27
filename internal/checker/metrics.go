package checker

import "time"

type Metrics struct {
	TotalURLs   int
	Successful  int
	Failed      int
	SuccessRate float64

	AverageLatency time.Duration

	FastestURL     string
	FastestLatency time.Duration

	SlowestURL     string
	SlowestLatency time.Duration
}

func CalculateMetrics(
	results []URLResult,
) Metrics {

	metrics := Metrics{}

	if len(results) == 0 {
		return metrics
	}

	metrics.TotalURLs = len(results)

	var totalLatency time.Duration

	metrics.FastestLatency = results[0].Latency
	metrics.SlowestLatency = results[0].Latency

	for _, result := range results {

		if result.Success {

			metrics.Successful++

			totalLatency += result.Latency

			// Fastest
			if result.Latency < metrics.FastestLatency {

				metrics.FastestLatency = result.Latency
				metrics.FastestURL = result.URL
			}

			// Slowest
			if result.Latency > metrics.SlowestLatency {

				metrics.SlowestLatency = result.Latency
				metrics.SlowestURL = result.URL
			}

		} else {

			metrics.Failed++
		}
	}

	if metrics.Successful > 0 {

		metrics.AverageLatency =
			totalLatency / time.Duration(metrics.Successful)
	}

	metrics.SuccessRate =
		(float64(metrics.Successful) / float64(metrics.TotalURLs)) * 100

	return metrics
}
