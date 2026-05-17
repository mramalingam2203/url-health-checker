package output

import (
	"fmt"

	"healthchecker/internal/checker"
)

func PrintResult(result checker.URLResult) {
	fmt.Println("===================================")
	fmt.Println("URL:", result.URL)

	if result.Success {
		fmt.Println("Status:", result.Status)
		fmt.Println("Status Code:", result.StatusCode)
		fmt.Println("Latency:", result.Latency)
	} else {
		fmt.Println("Error:", result.Error)
	}

	fmt.Println()
}
