package url

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLSFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}
	var urls []string

	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, errors.New("Couldn't parse base url")
	}

	var extractLinks func(*html.Node)
	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// Look for the href attribute
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					parsedURL, err := url.Parse(attr.Val)
					if err == nil {
						// Resolve relative URL to absolute
						resolvedURL := base.ResolveReference(parsedURL)
						urls = append(urls, resolvedURL.String())
					}
				}
			}
		}
		// Recursively check child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	extractLinks(doc)
	return urls, nil
}

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return "", errors.New("HTTP request failed with status: " + res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", errors.New("invalid content type: " + contentType)
	}
	htmlBody, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(htmlBody), nil
}
