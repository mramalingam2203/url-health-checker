package output

import (
	"encoding/json"
	"fmt"

	"healthchecker/internal/checker"
)

func PrintJSON(result checker.URLResult) {

	data, err := json.MarshalIndent(
		result,
		"",
		"  ",
	)

	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	fmt.Println(string(data))
}
