package anagramsarray

import "sort"

/*
	Массив анаграмм

	Дан массив строк strs. Сгруппируйте анаграммы вместе. Вы можете вернуть ответ в любом порядке.

	Анаграмма — это слово или фраза,
	полученная путем перестановки букв другого слова или фразы,
	обычно с использованием всех исходных букв ровно один раз.

	Входные данные: ["eat","tea","tan","ate","nat","bat"]
	Выходные: [["bat"],["nat","tan"],["ate","eat","tea"]]
	Входные данные: ["won","now","aaa","ooo","ooo"]
	Выходные: [["aaa"],["ooo", "ooo"],["won","now"]]
*/

func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	buckets := make(map[string][]string)
	for _, s := range strs {
		key := sortString(s)
		buckets[key] = append(buckets[key], s)
	}

	res := make([][]string, 0, len(buckets))
	for _, group := range buckets {
		res = append(res, group)
	}
	return res
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
