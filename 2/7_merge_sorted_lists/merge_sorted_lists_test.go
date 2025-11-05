package mergesortedlists

import (
	"algorithms/2/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseList(t *testing.T) {
	tests := []struct {
		name string

		l1 *pkg.LinkedList
		l2 *pkg.LinkedList

		expected *pkg.LinkedList
	}{
		{
			name: "1",

			l1: pkg.BuildList(1, 3, 5),
			l2: pkg.BuildList(2, 4),

			expected: pkg.BuildList(1, 2, 3, 4, 5),
		},
		{
			name: "2",

			l1: pkg.BuildList(100),
			l2: pkg.BuildList(2, 4),

			expected: pkg.BuildList(2, 4, 100),
		},
		{
			name: "3",

			l1: pkg.BuildList(-1),
			l2: pkg.BuildList(2, 4),

			expected: pkg.BuildList(-1, 2, 4),
		},
		{
			name: "4",

			l1: pkg.BuildList(),
			l2: pkg.BuildList(2, 4),

			expected: pkg.BuildList(2, 4),
		},
		{
			name: "5",

			l1: pkg.BuildList(4, 6),
			l2: pkg.BuildList(),

			expected: pkg.BuildList(4, 6),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := MergeSortedList(test.l1, test.l2)

			assert.Equal(t, test.expected, l)
		})
	}
}
