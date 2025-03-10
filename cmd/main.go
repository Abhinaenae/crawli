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
	const maxConcurrency = 5
	cfg, err := interurl.Configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}
	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	cfg.Wg.Add(1)
	go cfg.CrawlPage(rawBaseURL)
	cfg.Wg.Wait()

	for normalizedURL, count := range cfg.Pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}

}
