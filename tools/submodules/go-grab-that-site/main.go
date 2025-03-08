package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DownloadFile saves the content from a URL to a local file
func DownloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	_, err = io.Copy(out, resp.Body)
	return err
}

// ExtractLinks extracts all URLs from an HTML document
func ExtractLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var links []string
	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			for _, attr := range t.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					link := attr.Val
					if strings.HasPrefix(link, "/") {
						link = url + link
					}
					links = append(links, link)
				}
			}
		default:
			// Handle any other token types
		}
	}
}

func main() {
	startURL := "https://example.com"
	outputDir := "downloaded_site"

	fmt.Println("Starting download from:", startURL)
	links, err := ExtractLinks(startURL)
	if err != nil {
		fmt.Println("Error extracting links:", err)
		return
	}

	for _, link := range links {
		fmt.Println("Downloading:", link)
		filePath := filepath.Join(outputDir, filepath.Base(link))
		if err := DownloadFile(link, filePath); err != nil {
			fmt.Println("Failed to download", link, ":", err)
		}
	}
	fmt.Println("Download complete!")
}
