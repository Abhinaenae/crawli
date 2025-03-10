package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abhinaenae/crawli/internal/format"
	interurl "github.com/abhinaenae/crawli/internal/url"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawli <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}

	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	maxConcurrencyStr := os.Args[2]
	maxPagesStr := os.Args[3]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyStr)
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v", err)
		return
	}

	maxPages, err := strconv.Atoi(maxPagesStr)
	if err != nil {
		fmt.Printf("Error - maxPages: %v", err)
		return
	}
	cfg, err := interurl.Configure(rawBaseURL, maxPages, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}
	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	cfg.Wg.Add(1)
	go cfg.CrawlPage(rawBaseURL)
	cfg.Wg.Wait()

	format.PrintReport(cfg.Pages, rawBaseURL)

}
