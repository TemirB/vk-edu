package feedtheanimals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_feedTheAnimal(t *testing.T) {
	tests := []struct {
		name string

		animals []int
		food    []int

		expected int
	}{
		{
			name: "1",

			animals: []int{3, 4, 7},
			food:    []int{8, 1, 2},

			expected: 1,
		},
		{
			name: "2",

			animals: []int{3, 8, 1, 4},
			food:    []int{1, 1, 2},

			expected: 1,
		},
		{
			name: "3",

			animals: []int{1, 2, 2},
			food:    []int{7, 1},

			expected: 2,
		},
		{
			name: "4",

			animals: []int{8, 2, 3, 2},
			food:    []int{1, 4, 3, 8},

			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := feedTheAnimals(test.food, test.animals)

			assert.Equal(t, actual, test.expected)
		})
	}
}
