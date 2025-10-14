package moveevennumbers

/*
	Дан не отсортированный массив целых чисел

	Необходимо перенести в начало массива все четные числа, сохраняя их очередность.

	То есть если 8 стояла после 2, то после переноса в начало, он по-прежнему должна стоять после 2.
	Пример ввода: [3, 2, 4, 1, 11, 8, 9]
	Пример вывода: 2 4 8 1 11 3 9
*/

func moveEvenNumbers(numbers []int) {
	even := 0

	for i := range numbers {
		if numbers[i]%2 == 0 {
			numbers[even], numbers[i] = numbers[i], numbers[even]
			even++
		}
	}
}
