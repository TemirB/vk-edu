package finddiff

/*
	Найти разницу между двух строк

	На вход функции подается две строки: a и b.
	Строка b образована из строки a путем перемешивания и добавления одной буквы.
	Необходимо вернуть эту букву
*/

func findDifference(a, b string) rune {
	primer := make(map[rune]int)

	for _, aChar := range a {
		primer[aChar]++
	}

	for _, bChar := range b {
		primer[bChar]--
		if primer[bChar] < 0 {
			return bChar
		}
	}

	return ' '
}
