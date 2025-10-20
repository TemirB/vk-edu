package netherlands

/*
	Дан массив состоящий из нулей, единиц и двоек

	Необходимо его отсортировать за линейное время
*/

func sortNetherlands(arr []int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
			continue
		}

		if arr[mid] == 2 {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
			continue
		}

		mid++
	}
}
