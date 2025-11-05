package ispalindrome

import "algorithms/2/pkg"

/*
	Является ли слово палиндромом?

	Напишите функцию, которая принимает на вход строку и
	возвращает true, если она является палиндромом*.
	В противном случае верните false.

	*слово или текст, одинаково читающиеся в обоих направлениях.
*/

func IsPalindromeOnStack(s string) bool {
	stack := pkg.NewStack(nil)
	for _, r := range s {
		stack.Push(r)
	}

	for _, r := range s {
		if val, ok := stack.Pop(); ok && r != val {
			return false
		}
	}

	return true
}

func ispalindromeInDeque(s string) bool {
	d := pkg.NewDeque()

	for _, r := range s {
		d.PushBack(r)
	}

	for !d.Empty() {
		if d.Head() == d.Tail() {
			return true
		}

		left, _ := d.PopFront()
		right, _ := d.PopBack()

		if left != right {
			return false
		}
	}

	return true
}
