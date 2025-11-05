package veryeasyproblem

/*
	Как быстро можно сделать N копий документа, используя два ксерокса,
	каждый копирует со своей скоростью (x и y минут)?

	Разрешается использовать как один ксерокс, так и оба одновременно.
	Можно копировать не только с оригинала, но и с копии.
*/

// Сначала делаем копию, чтобы работало 2 принтера,
// затем нужно скопировать N-1 документов,
// причем оба ксерокса работают одновременно.

func bSearch(N, x, y int) int {
	l := 0
	r := (N - 1) * max(x, y)

	for l+1 < r {
		middle := (r + l) / 2

		if middle/x+middle/y < N-1 {
			l = middle
		} else {
			r = middle
		}
	}

	return r + min(x, y)
}

func veryEasyProblem(N, x, y int) int {
	if N == 1 {
		return min(x, y)
	}

	first := min(x, y)
	need := N - 1

	L := lcm(x, y)   // длина циклов
	per := L/x + L/y // кол-во копий в одном цикле

	cycles := need / per // кол-во полных циклов
	times := cycles * L
	remains := need % per // сколько осталось сделать самому быстрому, чтобы закончить

	if remains == 0 {
		return first + times
	}

	low, high := 0, L
	for low+1 < high {
		mid := (low + high) / 2
		if mid/x+mid/y >= remains {
			high = mid
		} else {
			low = mid
		}
	}
	return first + times + high
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
