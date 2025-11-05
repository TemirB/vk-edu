package detectcycle

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"algorithms/2/pkg"
)

func Test_HasCycle(t *testing.T) {
	cycledNodes := make([]*pkg.Node, 8)
	for i := 0; i < 8; i++ {
		cycledNodes[i] = pkg.NewNode(i, nil)
	}

	for i := range 7 {
		cycledNodes[i].SetNext(cycledNodes[i+1])
	}

	cycledNodes[7].SetNext(cycledNodes[2]) // замыкания

	cycleQueue := pkg.NewList(cycledNodes[0])

	cases := []struct {
		name string

		ll       *pkg.LinkedList
		expected bool
	}{
		{
			name: "1",

			ll:       pkg.NewList(pkg.NewNode(1, nil)),
			expected: false,
		},
		{
			name: "2",

			ll:       cycleQueue,
			expected: true,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			actual := HasCycle(test.ll)

			assert.Equal(t, actual, test.expected)
		})
	}
}
