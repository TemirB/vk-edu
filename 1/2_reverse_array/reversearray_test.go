package reversearray

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_reverseArray(t *testing.T) {
	cases := []struct {
		name string

		array    []int
		expected []int
	}{
		{
			name: "1",

			array:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "2",

			array:    []int{3, 3},
			expected: []int{3, 3},
		},
		{
			name: "3",

			array:    []int{1, 2, 4},
			expected: []int{4, 2, 1},
		},
		{
			name: "4",

			array:    []int{},
			expected: []int{},
		},
		{
			name: "5",

			array:    []int{1},
			expected: []int{1},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			reverseArray(test.array)

			assert.Equal(t, test.array, test.expected)
		})
	}
}
