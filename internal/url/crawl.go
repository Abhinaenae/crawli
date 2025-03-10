package url

import (
	"fmt"
	"net/url"
)

func CrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Make sure the rawCurrentURL is on the same domain as the rawBaseURL. If it's not, return
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	//Get a normalized version of the rawCurrentURL.
	normalizedCurrentURL, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		return
	}
	if normalizedCurrentURL == "" {
		return
	}
	//increment if visited
	_, ok := pages[normalizedCurrentURL]
	if ok {
		pages[normalizedCurrentURL]++
		return
	}

	//mark as visited
	pages[normalizedCurrentURL] = 1

	// skip other websites
	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	//Get the HTML from the current URL and find urls
	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error - GetHTML: %v", err)
		return
	}
	nextURLs, err := GetURLSFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("error - GetURLSFromHTML: %v", err)
		return
	}

	//Recursively crawl each URL on the page
	for _, nextURL := range nextURLs {
		CrawlPage(rawBaseURL, nextURL, pages)
	}

}
