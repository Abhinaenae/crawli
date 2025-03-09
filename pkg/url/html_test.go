package url

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLSFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute URL",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html>
	<body>
		<a href="https://blog.example.dev">
			<span>Example.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.example.dev"},
		},
		{
			name:     "relative URL",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Example.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.example.dev/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>example.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>example.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.example.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>example.dev></span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "bad HTML",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html body>
	<a href="path/one">
		<span>example.dev></span>
	</a>
</html body>
`,
			expected: []string{"https://blog.example.dev/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>example.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "handle invalid base URL",
			inputURL: `:\\invalidBaseURL`,
			inputBody: `
<html>
	<body>
		<a href="/path">
			<span>example.dev</span>
		</a>
	</body>
</html>
`,
			expected:      nil,
			errorContains: "Couldn't parse base url",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetURLSFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) || err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
