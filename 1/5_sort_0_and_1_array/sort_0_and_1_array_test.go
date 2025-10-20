package sort0and1arrays

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_sort(t *testing.T) {
	cases := []struct {
		name string

		array    []int
		expected []int
	}{
		{
			name: "1",

			array:    []int{0, 0, 1, 0, 1, 1, 0},
			expected: []int{0, 0, 0, 0, 1, 1, 1},
		},
		{
			name: "2",

			array:    []int{0, 1},
			expected: []int{0, 1},
		},
		{
			name: "3",

			array:    []int{1, 0, 0, 0},
			expected: []int{0, 0, 0, 1},
		},
		{
			name: "4",

			array:    []int{0, 1, 0, 1},
			expected: []int{0, 0, 1, 1},
		},
		{
			name: "5",

			array:    []int{1},
			expected: []int{1},
		},
		{
			name: "6",

			array:    []int{1, 0, 0, 1, 0, 0, 1},
			expected: []int{0, 0, 0, 0, 1, 1, 1},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			sort(test.array)

			assert.Equal(t, test.array, test.expected)
		})
	}
}
