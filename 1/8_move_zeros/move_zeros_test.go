package movezeros

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_moveZeros(t *testing.T) {
	cases := []struct {
		name string

		array    []int
		expected []int
	}{
		{
			name: "1",

			array:    []int{0, 0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0, 0},
		},
		{
			name: "2",

			array:    []int{0, 33, 57, 88, 60, 0, 0, 80, 99},
			expected: []int{33, 57, 88, 60, 80, 99, 0, 0, 0},
		},
		{
			name: "3",

			array:    []int{0, 0, 0, 18, 16, 0, 0, 77, 99},
			expected: []int{18, 16, 77, 99, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			moveZeros(test.array)

			assert.Equal(t, test.array, test.expected)
		})
	}
}
