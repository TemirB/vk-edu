package middlenode

import (
	"testing"

	"algorithms/2/pkg"

	"github.com/stretchr/testify/assert"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name string

		ll       *pkg.LinkedList
		expected int
	}{
		{
			name: "1",

			ll:       pkg.BuildList(1, 2, 3, 4, 5),
			expected: 3,
		},
		{
			name: "2",

			ll:       pkg.BuildList(1, 2, 3, 4),
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MiddleNode(test.ll)

			assert.Equal(t, got.Data(), test.expected)
		})
	}
}
