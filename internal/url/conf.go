package url

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	Pages              map[string]int
	MaxPages           int
	BaseURL            *url.URL
	Mu                 *sync.Mutex
	ConcurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
}

func Configure(rawBaseURL string, maxPages, maxConcurrency int) (*Config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	return &Config{
		Pages:              make(map[string]int),
		MaxPages:           maxPages,
		BaseURL:            baseURL,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, maxConcurrency),
		Wg:                 &sync.WaitGroup{},
	}, nil
}

func (cfg *Config) pagesLen() int {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()
	return len(cfg.Pages)
}
