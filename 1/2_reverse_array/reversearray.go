package reversearray

/*
	Дан массив целых чисел. Необходимо развернуть его.
	Сделать это надо за линейное время без дополнительных аллокаций.
*/

func reverseArray(arr []int) {

	left, right := 0, len(arr)-1

	for right > left {
		arr[right], arr[left] = arr[left], arr[right]

		left++
		right--
	}
}
