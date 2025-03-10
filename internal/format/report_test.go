package format

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []pageData
	}{
		{
			name:  "Basic sorting",
			input: map[string]int{"a.com": 3, "b.com": 1, "c.com": 2},
			expected: []pageData{
				{"a.com", 3},
				{"c.com", 2},
				{"b.com", 1},
			},
		},
		{
			name: "Sorting with ties (alphabetically)",
			input: map[string]int{
				"b.com": 2,
				"a.com": 2,
				"c.com": 2,
			},
			expected: []pageData{
				{"a.com", 2},
				{"b.com", 2},
				{"c.com", 2},
			},
		},
		{
			name:     "Empty input",
			input:    map[string]int{},
			expected: []pageData{},
		},
		{
			name: "Single entry",
			input: map[string]int{
				"single.com": 5,
			},
			expected: []pageData{
				{"single.com", 5},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sortPages(tc.input)

			// Special handling for empty slices
			if len(got) == 0 && len(tc.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("sortPages(%v) = %v, expected %v", tc.input, got, tc.expected)
			}
		})
	}
}
