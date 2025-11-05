package detectcycle

import (
	"algorithms/2/pkg"
)

/*
	Проверить, является ли список циклическим

	Дан односвязный список.
	Необходимо проверить, является ли этот список циклическим.

	Циклическим (кольцевым) списком называется список,
	у которого последний узел ссылается на один из предыдущих узлов.
*/

func HasCycle(list *pkg.LinkedList) bool {
	if list == nil {
		return false
	}
	dict := make(map[*pkg.Node]struct{})
	current := list.GetHead()
	for current.Next() != nil {
		if _, exist := dict[current]; exist {
			return true
		}

		dict[current] = struct{}{}
		current = current.Next()
	}

	return false
}
