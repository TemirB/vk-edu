package reversepartofarray

/*
	Дан массив целых чисел.
	Необходимо повернуть (сдвинуть) справа налево часть массива, которая указана вторым параметром.
	Сделать это надо за линейное время без дополнительных аллокаций

	Исходный массив: 1, 2, 3, 4, 5, 6, 7
	k = 3
	Результат: 5, 6, 7, 1, 2, 3, 4
*/

func reverseArray(arr []int, left, right int) {
	for right > left {
		arr[left], arr[right] = arr[right], arr[left]

		left++
		right--
	}
}

func reversePartOfArray(arr []int, k int) {
	if k+1 >= len(arr) {
		return
	}

	left, right := 0, len(arr)-1

	reverseArray(arr, left, right)
	// 7, 6, 5, 4, 3, 2, 1
	// left = 0, right = len(arr) - 1 - (k + 1)
	right = len(arr) - k - 2
	reverseArray(arr, left, right)
	// 5, 6, 7, 4, 3, 2, 1
	left, right = right+1, len(arr)-1
	reverseArray(arr, left, right)
}
