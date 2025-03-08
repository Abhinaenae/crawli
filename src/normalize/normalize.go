package normalize

import (
	"net/url"
	"strings"
)

func NormalizeURL(rawURL string) (string, error) {
	urlData, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	normalized := urlData.Host + urlData.Path
	normalized = strings.TrimSuffix(normalized, "/")
	return normalized, nil
}
