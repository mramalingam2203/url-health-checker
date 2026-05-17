package main

import (
	"net/http"
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

	for _, url := range urls {
		result := checker.CheckURL(client, url)
		output.PrintResult(result)
	}
}
