package output

import (
	"fmt"

	"healthchecker/internal/checker"
)

func PrintCSV(result checker.URLResult) {

	fmt.Printf(
		"%s,%d,%s,%v,%t,%s,%d\n",
		result.URL,
		result.StatusCode,
		result.Status,
		result.Latency,
		result.Success,
		result.Error,
		result.Attempts,
	)
}
