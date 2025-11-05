package twosum

import (
	"testing"
)

func sameValues(t *testing.T, ans [2]int, a, b int) bool {
	t.Helper()

	return (ans[0] == a && ans[1] == b) || (ans[0] == b && ans[1] == a)
}
func Test_twoSum(t *testing.T) {
	tests := []struct {
		name string

		data   []int
		target int

		wantA int
		wantB int
	}{
		{"1", []int{2, 7, 11, 15}, 9, 2, 7},
		{"2", []int{3, 2, 4}, 6, 2, 4},
		{"3", []int{3, 3}, 6, 3, 3},
		{"4", []int{-1, -2, -3, -4, -5}, -8, -3, -5},
		{"5", []int{1, 1, 1}, 2, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.data, tt.target)

			if !sameValues(t, got, tt.wantA, tt.wantB) {
				t.Fatalf("TwoSum(%v, %d) = (%v, %v), want (%d,%d), true",
					tt.data, tt.target, got[0], got[1], tt.wantA, tt.wantB)
			}
		})
	}
}

func TestTwoSum_NoSolution(t *testing.T) {
	if got := twoSum([]int{1, 2, 3}, 7); got != [2]int{0, 0} {
		t.Fatalf("expected (0, 0), got (%v, %v)", got[0], got[1])
	}
}

func TestTwoSum_NoReuseSameIndex(t *testing.T) {
	if got := twoSum([]int{3}, 6); got != [2]int{0, 0} {
		t.Fatalf("should not reuse same index; expected (nil, false), got (%v, %v)", got[0], got[1])
	}
}
