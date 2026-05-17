package reader

import (
	"bufio"
	"os"
)

func ReadURLs(filename string) ([]string, error) {
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
