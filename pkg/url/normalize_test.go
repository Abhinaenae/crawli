package url

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.example.org/path",
			expected: "blog.example.org/path",
		},
		{
			name:     "remove HTTP scheme",
			inputURL: "http://blog.example.org/path",
			expected: "blog.example.org/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.example.org/path/",
			expected: "blog.example.org/path",
		},
		{
			name:     "handle empty path",
			inputURL: "https://blog.example.org/",
			expected: "blog.example.org",
		},
		{
			name:     "handle empty path without slash",
			inputURL: "https://blog.example.org",
			expected: "blog.example.org",
		},
		{
			name:     "handle complex path",
			inputURL: "http://example.net/foo/bar/",
			expected: "example.net/foo/bar",
		},
		{
			name:     "handle subdomains",
			inputURL: "https://sub.testsite.com/",
			expected: "sub.testsite.com",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
