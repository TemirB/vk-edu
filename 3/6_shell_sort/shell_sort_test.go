package shellsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellSort_Basic(t *testing.T) {
	cases := []struct {
		name string

		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"one", []int{42}, []int{42}},
		{"sorted", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"dups", []int{3, 1, 2, 3, 2, 1}, []int{1, 1, 2, 2, 3, 3}},
		{"negs", []int{-5, -1, 0, -3, 2, 2}, []int{-5, -3, -1, 0, 2, 2}},
		{"zeros", []int{0, 0, 0}, []int{0, 0, 0}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.in))
			copy(got, tt.in)
			ShellSort(got)

			assert.Equal(t, got, tt.want)
		})
	}
}
