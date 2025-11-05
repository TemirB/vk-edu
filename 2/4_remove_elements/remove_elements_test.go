package removeelements

import (
	"algorithms/2/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseList(t *testing.T) {
	tests := []struct {
		name string

		ll  *pkg.LinkedList
		val int

		expected *pkg.LinkedList
	}{
		{
			name: "1",

			ll:  pkg.BuildList(1, 2, 3, 4, 5),
			val: 3,

			expected: pkg.BuildList(1, 2, 4, 5),
		},
		{
			name: "2",

			ll:  pkg.BuildList(1, 7, 3, 19, 5),
			val: 1,

			expected: pkg.BuildList(7, 3, 19, 5),
		},
		{
			name: "3",

			ll:  pkg.BuildList(1, 7, 3, 19, 5),
			val: 5,

			expected: pkg.BuildList(1, 7, 3, 19),
		},
		{
			name: "4",

			ll:  pkg.BuildList(1, 7, 3, 19, 5),
			val: 77,

			expected: pkg.BuildList(1, 7, 3, 19, 5),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			RemoveElement(test.ll, test.val)

			assert.True(t, test.ll.Compare(test.expected))
		})
	}
}
