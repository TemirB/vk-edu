package finddiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// функция, которую мы тестируем
// func findDifference(a, b string) rune { ... }

func TestFindDifferenceBasic(t *testing.T) {
	tests := []struct {
		name string

		a string
		b string

		expected rune
	}{
		{"1", "abc", "cbad", 'd'},
		{"2", "aabbc", "babacd", 'd'},
		{"3", "", "z", 'z'},
		{"4", "hello", "ohlleo", 'o'},
		{"5", "qwerty", "qtrewyp", 'p'},
		{"6", "ab", "aba", 'a'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.a, tt.b)

			assert.Equal(t, got, tt.expected)
		})
	}
}

func TestFindDifferenceUnicode(t *testing.T) {
	a := "кот"
	b := "акот"
	expected := 'а'

	result := findDifference(a, b)
	if result != expected {
		t.Errorf("findDifference(%q, %q) = %q, expected %q",
			a, b, result, expected)
	}
}
