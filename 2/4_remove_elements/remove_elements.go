package removeelements

import "algorithms/2/pkg"

/*
	Удалить элемент из односвязного списка

	Необходимо написать функцию, которая принимает на вход односвязный список и некоторое целое число val.
	Необходимо удалить узел из списка, значение которого равно val.
*/

func RemoveElement(ll *pkg.LinkedList, val int) {
	dummy := pkg.NewNode(0, ll.GetHead())
	prev := dummy
	current := ll.GetHead()

	for current != nil {
		if current.Data() == val {
			prev.SetNext(current.Next())
		} else {
			prev = current
		}

		current = current.Next()
	}

	ll.SetHead(dummy.Next())
}
