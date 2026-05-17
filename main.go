package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func readURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := scanner.Text()

		if url != "" {
			urls = append(urls, url)
		}
	}

	return urls, scanner.Err()
}

func checkURL(client *http.Client, url string) {
	fmt.Println("Checking:", url)

	start := time.Now()

	resp, err := client.Get(url)

	latency := time.Since(start)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println()
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Latency:", latency)
	fmt.Println()
}

func main() {
	urls, err := readURLs("urls.txt")
	if err != nil {
		fmt.Println("Failed to read URL file:", err)
		return
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, url := range urls {
		checkURL(client, url)
	}
}
