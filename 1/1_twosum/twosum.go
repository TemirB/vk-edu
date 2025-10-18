package twosum

/*
	Дан отсортированный по возрастанию массив целых чисел
и некоторое число target.
	Необходимо найти два числа в массиве, которые в сумме дают
заданное значение target, и вернуть их индексы.
*/

func twosum(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1
	for right > left {
		if arr[left]+arr[right] == target {
			return left, right
		} else if arr[left]+arr[right] > target {
			right--
		} else {
			left++
		}
	}
	return -1, -1
}
