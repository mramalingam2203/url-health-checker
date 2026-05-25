package reader

import (
	"os"
	"testing"
)

func TestReadURLs(t *testing.T) {

	filename := "test_urls.txt"

	content := `https://google.com
https://github.com`

	err := os.WriteFile(
		filename,
		[]byte(content),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(filename)

	urls, err := ReadURLs(filename)

	if err != nil {
		t.Fatal(err)
	}

	if len(urls) != 2 {
		t.Errorf(
			"expected 2 URLs, got %d",
			len(urls),
		)
	}
}
