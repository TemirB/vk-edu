package pkg

/*
	Служебный пакет, тут хранятся реализации необхлджимых структур, для решения задач
*/

type Node struct {
	data int
	next *Node
}

func NewNode(data int, next *Node) *Node {
	return &Node{
		data: data,
		next: next,
	}
}

func (n *Node) SetData(data int) {
	if n == nil {
		return
	}
	n.data = data
}

func (n *Node) Data() int {
	return n.data
}

func (n *Node) SetNext(nextNode *Node) {
	n.next = nextNode
}

func (n *Node) Next() *Node {
	return n.next
}

type LinkedList struct {
	head *Node
}

func NewList(head *Node) *LinkedList {
	return &LinkedList{
		head: head,
	}
}

func (ll *LinkedList) GetHead() *Node {
	return ll.head
}

func (ll *LinkedList) SetHead(newHead *Node) {
	ll.head = newHead
}

// Эта функция создает лист по набору values
func BuildList(values ...int) *LinkedList {
	ll := &LinkedList{}
	var prev *Node

	for _, v := range values {
		n := NewNode(v, nil)

		if ll.GetHead() == nil {
			ll.SetHead(n)
		} else {
			prev.SetNext(n)
		}
		prev = n
	}

	return ll
}

// Превращает лист, в слайс
// Значения в котором = значения нод, в начале стоит head, в конце tail
func ToSlice(head *Node) []int {
	out := make([]int, 0)
	for p := head; p != nil; p = p.next {
		out = append(out, p.data)
	}
	return out
}

// Вспомогательная функция, позволяющая сравнить 2 листа
func (ll *LinkedList) Compare(other *LinkedList) bool {
	if ll == nil || other == nil {
		return ll == other
	}
	a, b := ll.head, other.head
	for a != nil && b != nil {
		if a.data != b.data {
			return false
		}
		a = a.Next()
		b = b.Next()
	}
	return a == nil && b == nil
}

func Equal(a, b *Node) bool {
	for a != nil && b != nil {
		if a.data != b.data {
			return false
		}
		a, b = a.next, b.next
	}
	return a == nil && b == nil
}
