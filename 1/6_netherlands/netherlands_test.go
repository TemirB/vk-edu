package netherlands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_netherlands(t *testing.T) {
	cases := []struct {
		name string

		array    []int
		expected []int
	}{
		{
			name: "1",

			array:    []int{0, 0, 1, 2, 1, 1, 2},
			expected: []int{0, 0, 1, 1, 1, 2, 2},
		},
		{
			name: "2",

			array:    []int{0, 1, 2},
			expected: []int{0, 1, 2},
		},
		{
			name: "3",

			array:    []int{1, 2, 2, 0},
			expected: []int{0, 1, 2, 2},
		},
		{
			name: "4",

			array:    []int{2, 1, 0, 1},
			expected: []int{0, 1, 1, 2},
		},
		{
			name: "5",

			array:    []int{1, 2, 0, 1, 0, 0, 1},
			expected: []int{0, 0, 0, 1, 1, 1, 2},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			sortNetherlands(test.array)

			assert.Equal(t, test.array, test.expected)
		})
	}
}
