package issubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseList(t *testing.T) {
	tests := []struct {
		name string

		a, b string

		expected bool
	}{
		{
			name: "1",

			a: "subque",
			b: "subsequence",

			expected: true,
		},
		{
			name: "2",

			a: "subqze",
			b: "subsequence",

			expected: false,
		},
		{
			name: "3",

			a: "abs",
			b: "abobus",

			expected: true,
		},
		{
			name: "4",

			a: "",
			b: "",

			expected: true,
		},
		{
			name: "5",

			a: "a",
			b: "",

			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsSubsequence(test.a, test.b)

			assert.Equal(t, test.expected, got)
		})
	}
}
