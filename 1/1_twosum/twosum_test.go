package twosum

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_twosum(t *testing.T) {
	cases := []struct {
		name string

		array                []int
		target               int
		expected1, expected2 int
	}{
		{
			name: "1",

			array:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:    5,
			expected1: 0,
			expected2: 3,
		},
		{
			name: "2",

			array:     []int{3, 3},
			target:    6,
			expected1: 0,
			expected2: 1,
		},
		{
			name: "3",

			array:     []int{1, 2, 4},
			target:    6,
			expected1: 1,
			expected2: 2,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			idx1, idx2 := twosum(test.array, test.target)

			assert.Equal(t, idx1, test.expected1)
			assert.Equal(t, idx2, test.expected2)
		})
	}
}
