package merge2sortarray

/*
	Дано два отсортированных массива.

	Необходимо написать функцию которая объединит эти два массива в один отсортированный.

	p.s.: По условию не ясно какой вариант решать, буду решать тот что без аллокации
	(с учетом того, что 1 массивчик проиннициализирован 0, там где )
*/

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
