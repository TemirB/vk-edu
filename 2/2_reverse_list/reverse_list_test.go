package reverselist

import (
	"testing"

	"algorithms/2/pkg"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseList(t *testing.T) {
	tests := []struct {
		name string

		ll       *pkg.LinkedList
		expected *pkg.LinkedList
	}{
		{
			name: "1",

			ll:       pkg.BuildList(1, 2, 3, 4, 5),
			expected: pkg.BuildList(5, 4, 3, 2, 1),
		},
		{
			name: "2",

			ll:       pkg.BuildList(1, 7, 3, 19, 5),
			expected: pkg.BuildList(5, 19, 3, 7, 1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ReverseList(test.ll)

			assert.True(t, test.ll.Compare(test.expected))
		})
	}
}
