package shellsort

func ShellSort(a []int) {
	n := len(a)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			tmp := a[i]
			j := i
			for j >= gap && a[j-gap] > tmp {
				a[j] = a[j-gap]
				j -= gap
			}
			a[j] = tmp
		}
	}
}
