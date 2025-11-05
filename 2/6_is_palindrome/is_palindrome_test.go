package ispalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseList(t *testing.T) {
	tests := []struct {
		name string

		s        string
		expected bool
	}{
		{
			name: "1",

			s: "abcdedcba",

			expected: true,
		},
		{
			name: "2",

			s: "avcdedcba",

			expected: false,
		},
		{
			name: "3",

			s: "a",

			expected: true,
		},
		{
			name: "4",

			s: "abccba",

			expected: true,
		},
		{
			name: "5",

			s: "bobba",

			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotFromStack := IsPalindromeOnStack(test.s)
			gotFromDeque := ispalindromeInDeque(test.s)

			assert.Equal(t, test.expected, gotFromDeque)
			assert.Equal(t, test.expected, gotFromStack)
		})
	}
}
