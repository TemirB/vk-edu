package merge2sortarray

/*
	Дано два отсортированных массива.

	Необходимо написать функцию которая объединит эти два массива в один отсортированный.
*/

// Решение 1
func mergeWithNewSlice(arr1, arr2 []int) []int {
	out := make([]int, 0)
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			out = append(out, arr1[i])
			i++
		} else {
			out = append(out, arr2[j])
			j++
		}
	}

	out = append(out, arr1[i:]...)
	out = append(out, arr2[j:]...)

	return out
}

// Решение 2 (без доп аллокации)
func merge(arr1, arr2 []int) {
	writeIdx := len(arr1) - 1
	idx1 := len(arr1) - len(arr2) - 1
	idx2 := len(arr2) - 1

	for idx2 >= 0 {
		if idx1 >= 0 && arr1[idx1] > arr2[idx2] {
			arr1[writeIdx] = arr1[idx1]
			idx1--
		} else {
			arr1[writeIdx] = arr2[idx2]
			idx2--
		}
		writeIdx--
	}
}
