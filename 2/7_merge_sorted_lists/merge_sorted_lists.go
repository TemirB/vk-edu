package mergesortedlists

import "algorithms/2/pkg"

/*
	Слияние двух отсортированных списков

	Написать функцию, которая принимает на вход два отсортированных односвязных списка
	и объединяет их в один отсортированный список.
	При этом затраты по памяти должны быть O(1)

	Входные данные: list1 = [3, 6, 8], list2 = [4, 7, 9, 11]
	Выходные: [3,4,6,7,8,9,11]
*/

func MergeSortedList(l1, l2 *pkg.LinkedList) *pkg.LinkedList {
	var h1, h2 *pkg.Node
	if l1 != nil {
		h1 = l1.GetHead()
	}
	if l2 != nil {
		h2 = l2.GetHead()
	}

	head := MergeSorted(h1, h2)
	return pkg.NewList(head)
}

func MergeSorted(a, b *pkg.Node) *pkg.Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	dummy := pkg.NewNode(0, nil)
	tail := dummy

	for a != nil && b != nil {
		if a.Data() <= b.Data() {
			nextA := a.Next()
			tail.SetNext(a)
			tail = a
			a = nextA
		} else {
			nextB := b.Next()
			tail.SetNext(b)
			tail = b
			b = nextB
		}
	}

	if a != nil {
		tail.SetNext(a)
	} else {
		tail.SetNext(b)
	}

	return dummy.Next()
}
