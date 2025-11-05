package anagramsarray

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func normalize2d(arr [][]string) [][]string {
	// сортируем слова внутри групп
	for i := range arr {
		sort.Strings(arr[i])
	}

	// сортируем группы как строки, чтобы порядок не влиял на равенство
	sort.Slice(arr, func(i, j int) bool {
		// сравним как строки (лексикографически)
		return arr[i][0] < arr[j][0]
	})

	return arr
}

func TestGroupAnagrams(t *testing.T) {

	tests := []struct {
		name string
		in   []string
		want [][]string
	}{
		{
			name: "example_1",
			in:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "example_2",
			in:   []string{"won", "now", "aaa", "ooo", "ooo"},
			want: [][]string{
				{"aaa"},
				{"ooo", "ooo"},
				{"won", "now"},
			},
		},
		{
			name: "unicode",
			in:   []string{"кот", "ток", "кто", "акт", "так"},
			want: [][]string{
				{"акт", "так"},
				{"кот", "кто", "ток"},
			},
		},
		{
			name: "duplicates",
			in:   []string{"ab", "ba", "ab"},
			want: [][]string{
				{"ab", "ab", "ba"},
			},
		},
		{
			name: "single",
			in:   []string{"xyz"},
			want: [][]string{
				{"xyz"},
			},
		},
		{
			name: "empty",
			in:   []string{},
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.in)

			assert.Equal(
				t,
				normalize2d(tt.want),
				normalize2d(got),
				"группы анаграмм должны совпадать",
			)
		})
	}
}
