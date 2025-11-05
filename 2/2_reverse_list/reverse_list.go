package reverselist

import (
	pkg "algorithms/2/pkg"
)

/*
	Развернуть односвязный список

	Необходимо написать функцию,
	которая принимает на вход односвязный список и разворачивает его.
*/

func ReverseList(ll *pkg.LinkedList) {
	if ll == nil {
		return
	}

	var prev *pkg.Node
	current := ll.GetHead()

	for current != nil {
		next := current.Next()
		current.SetNext(prev)

		prev = current
		current = next
	}

	ll.SetHead(prev)
}
