package veryeasyproblem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_veryEasyProblem(t *testing.T) {
	tests := []struct {
		name string

		N, x, y  int
		expected int
	}{
		{
			name: "1",

			N: 5,
			x: 1,
			y: 2,

			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := veryEasyProblem(test.N, test.x, test.y)
			bSearchActual := bSearch(test.N, test.x, test.y)

			assert.Equal(t, bSearchActual, test.expected)
			assert.Equal(t, actual, test.expected)
		})
	}
}
