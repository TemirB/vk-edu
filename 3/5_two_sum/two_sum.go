package twosum

/*
	Дан не отсортированный массив целых чисел и некоторое число target.
	Необходимо написать функцию, которая найдет два таких элемента в массиве,
	сумма которых будет равна target.

	Один элемент можно использовать лишь один раз.
	В случае если два таких элемента не нашлось, возвращаем пустой массив.
*/

func twoSum(data []int, target int) [2]int {
	var ans [2]int
	dictionary := make(map[int]int)

	for i, v := range data {
		dictionary[v] = i
	}

	for i, v := range data {
		diff := target - v
		if j, ok := dictionary[diff]; ok && i != j {
			ans[0] = v
			ans[1] = data[j]
			return ans
		}
	}

	return ans
}
