package middlenode

import "algorithms/2/pkg"

/*
	Найти середину списка

	Дан связный список.

	Необходимо найти середину списка.
	Сделать это необходимо за O(n) без дополнительных аллокаций
*/

func MiddleNode(ll *pkg.LinkedList) *pkg.Node {
	if ll.GetHead() == nil {
		return nil
	}
	slow, fast := ll.GetHead(), ll.GetHead()
	for fast != nil {
		next := fast.Next()
		if next == nil {
			break
		}
		slow = slow.Next()
		fast = next.Next()
	}
	return slow
}
