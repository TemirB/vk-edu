package sort0and1arrays

/*
Дан массив, содержащий только 0 и 1.

Напишите функцию, которая сортирует его так,
чтобы все нули оказались в начале, а все единички — в конце.
*/
func sort(arr []int) {
	left, right := 0, len(arr)-1

	for right > left {
		if arr[left] == 0 {
			left++
			continue
		}
		if arr[right] == 1 {
			right--
			continue
		}

		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
