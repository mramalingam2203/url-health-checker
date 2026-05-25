package checker

import "time"

type URLResult struct {
	URL        string
	StatusCode int
	Status     string
	Latency    time.Duration
	Success    bool
	Error      string
	Attempts   int
}
