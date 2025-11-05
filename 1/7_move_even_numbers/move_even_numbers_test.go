package moveevennumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveEvenNumbers(t *testing.T) {
	cases := []struct {
		name string

		array    []int
		expected []int
	}{
		{
			name: "1",

			array:    []int{3, 2, 4, 1, 11, 8, 9},
			expected: []int{2, 4, 8, 1, 11, 3, 9},
		},
		{
			name: "2",

			array:    []int{8, 4, 1, 8},
			expected: []int{8, 4, 8, 1},
		},
		{
			name: "3",

			array:    []int{1, 2, 2, 0},
			expected: []int{2, 2, 0, 1},
		},
		{
			name: "4",

			array:    []int{2, 1, 0, 1},
			expected: []int{2, 0, 1, 1},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			moveEvenNumbers(test.array)

			assert.Equal(t, test.array, test.expected)
		})
	}
}
