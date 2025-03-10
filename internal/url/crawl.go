package url

import (
	"fmt"
	"net/url"
)

func (cfg *Config) CrawlPage(rawCurrentURL string) {
	cfg.ConcurrencyControl <- struct{}{}
	defer func() {
		<-cfg.ConcurrencyControl
		cfg.Wg.Done()
	}()

	// Make sure the rawCurrentURL is on the same domain as the rawBaseURL. If it's not, return
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	//Get a normalized version of the rawCurrentURL.
	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		return
	}
	if normalizedCurrentURL == "" {
		return
	}
	//increment if visited
	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}

	// skip other websites
	if currentURL.Hostname() != cfg.BaseURL.Hostname() {
		return
	}

	//Get the HTML from the current URL and find urls
	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error - GetHTML: %v", err)
		return
	}
	nextURLs, err := getURLSFromHTML(htmlBody, cfg.BaseURL)
	if err != nil {
		fmt.Printf("error - GetURLSFromHTML: %v", err)
		return
	}

	//Recursively crawl each URL on the page
	for _, nextURL := range nextURLs {
		cfg.Wg.Add(1)
		go cfg.CrawlPage(nextURL)
	}

}

func (cfg *Config) addPageVisit(normalizedURL string) bool {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()
	_, visited := cfg.Pages[normalizedURL]
	if visited {
		cfg.Pages[normalizedURL]++
		return false
	}
	//mark as visited
	cfg.Pages[normalizedURL] = 1
	return true
}
