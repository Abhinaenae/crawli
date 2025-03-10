package main

import (
	"fmt"
	"os"

	interurl "github.com/abhinaenae/crawli/internal/url"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args[1:]) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	pages := make(map[string]int)

	interurl.CrawlPage(rawBaseURL, rawBaseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
