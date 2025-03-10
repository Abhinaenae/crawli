package url

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	Pages              map[string]int
	BaseURL            *url.URL
	Mu                 *sync.Mutex
	ConcurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
}

func Configure(rawBaseURL string, maxConcurrency int) (*Config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	return &Config{
		Pages:              make(map[string]int),
		BaseURL:            baseURL,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, maxConcurrency),
		Wg:                 &sync.WaitGroup{},
	}, nil
}
