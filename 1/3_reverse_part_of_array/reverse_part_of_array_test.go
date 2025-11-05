package reversepartofarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reversePartOfArray(t *testing.T) {
	cases := []struct {
		name string

		array    []int
		k        int
		expected []int
	}{
		{
			name: "1",

			array:    []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "2",

			array:    []int{3, 3},
			k:        0,
			expected: []int{3, 3},
		},
		{
			name: "3",

			array:    []int{1, 2, 4},
			k:        0,
			expected: []int{2, 4, 1},
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
			reversePartOfArray(test.array, test.k)

			assert.Equal(t, test.array, test.expected)
		})
	}
}
