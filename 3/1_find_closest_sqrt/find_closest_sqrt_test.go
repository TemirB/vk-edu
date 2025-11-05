package findclosestsqrt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findClosestSqrt(t *testing.T) {
	tests := []struct {
		name string

		target   int
		expected int
	}{
		{
			name: "1",

			target:   5,
			expected: 2,
		},
		{
			name: "2",

			target:   4,
			expected: 2,
		},
		{
			name: "3",

			target:   8,
			expected: 2,
		},
		{
			name: "4",

			target:   17,
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findClosestSqrt(test.target)

			assert.Equal(t, actual, test.expected)
		})
	}
}
