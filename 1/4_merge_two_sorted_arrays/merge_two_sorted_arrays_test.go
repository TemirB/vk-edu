package merge2sortarray

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_merge(t *testing.T) {
	cases := []struct {
		name string

		array1   []int
		array2   []int
		expected []int
	}{
		{
			name: "1",

			array1:   []int{1, 3, 5, 7, 0, 0, 0, 0},
			array2:   []int{2, 4, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "2",

			array1:   []int{1, 3, 5, 7, 0, 0, 0, 0},
			array2:   []int{-4, -3, -2, -1},
			expected: []int{-4, -3, -2, -1, 1, 3, 5, 7},
		},
		{
			name: "3",

			array1:   []int{1, 2, 3, 4, 0, 0, 0},
			array2:   []int{-1, 5, 6},
			expected: []int{-1, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "4",

			array1:   []int{1, 2, 3, 4},
			array2:   []int{},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "5",

			array1:   []int{1, 0},
			array2:   []int{1},
			expected: []int{1, 1},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			merge(test.array1, test.array2)

			assert.Equal(t, test.array1, test.expected)
		})
	}
}

func Test_mergeWithNewSlice(t *testing.T) {
	cases := []struct {
		name string

		array1   []int
		array2   []int
		expected []int
	}{
		{
			name: "1",

			array1:   []int{1, 3, 5, 7},
			array2:   []int{2, 4, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "2",

			array1:   []int{1, 3, 5, 7},
			array2:   []int{-4, -3, -2, -1},
			expected: []int{-4, -3, -2, -1, 1, 3, 5, 7},
		},
		{
			name: "3",

			array1:   []int{1, 2, 3, 4},
			array2:   []int{-1, 5, 6},
			expected: []int{-1, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "4",

			array1:   []int{1, 2, 3, 4},
			array2:   []int{},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "5",

			array1:   []int{1},
			array2:   []int{0, 1},
			expected: []int{0, 1, 1},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			actual := mergeWithNewSlice(test.array1, test.array2)

			assert.Equal(t, actual, test.expected)
		})
	}
}
