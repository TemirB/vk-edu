package pkg

// RuneDequeNode — двусвязный узел для дека.
type RuneDequeNode struct {
	data rune
	prev *RuneDequeNode
	next *RuneDequeNode
}

func NewRuneDequeNode(d rune) *RuneDequeNode {
	return &RuneDequeNode{
		data: d,
	}
}

func (n *RuneDequeNode) Data() rune {
	return n.data
}

func (n *RuneDequeNode) Next() *RuneDequeNode {
	return n.next
}

func (n *RuneDequeNode) Prev() *RuneDequeNode {
	return n.prev
}

func (n *RuneDequeNode) SetNext(next *RuneDequeNode) {
	n.next = next
}

func (n *RuneDequeNode) SetPrev(prev *RuneDequeNode) {
	n.prev = prev
}

// Deque — двусторонняя очередь.
type Deque struct {
	head *RuneDequeNode
	tail *RuneDequeNode
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) Head() *RuneDequeNode {
	return d.head
}

func (d *Deque) Tail() *RuneDequeNode {
	return d.tail
}

func (d *Deque) PushFront(v rune) {
	newNode := NewRuneDequeNode(v)

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	newNode.SetNext(d.head)
	d.head.SetPrev(newNode)
	d.head = newNode
}

func (d *Deque) PushBack(v rune) {
	newNode := NewRuneDequeNode(v)

	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	newNode.SetPrev(d.tail)
	d.tail.SetNext(newNode)
	d.tail = newNode
}

func (d *Deque) PopFront() (rune, bool) {
	if d.head == nil {
		return 0, false
	}

	val := d.head.Data()
	d.head = d.head.Next()

	if d.head == nil {
		d.tail = nil
	} else {
		d.head.SetPrev(nil)
	}

	return val, true
}

func (d *Deque) PopBack() (rune, bool) {
	if d.tail == nil {
		return 0, false
	}

	val := d.tail.Data()
	d.tail = d.tail.Prev()

	if d.tail == nil {
		d.head = nil
	} else {
		d.tail.SetNext(nil)
	}

	return val, true
}

func (d *Deque) Empty() bool {
	return d.head == nil
}
