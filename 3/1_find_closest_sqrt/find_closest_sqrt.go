package findclosestsqrt

/*
	Найти корень числа (ближайшее целое)
*/

func findClosestSqrt(target int) int {
	left, right := 0, target

	for left <= right {
		middle := (left + right) / 2
		targetOverMiddle := target / middle

		if middle > targetOverMiddle {
			right = middle - 1
			continue
		}

		if middle < targetOverMiddle {
			left = middle + 1
			continue
		}

		return middle
	}

	return right
}
