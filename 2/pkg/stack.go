package pkg

type RuneNode struct {
	data rune
	next *RuneNode
}

func NewRuneNode(d rune, n *RuneNode) *RuneNode {
	return &RuneNode{
		data: d,
		next: n,
	}
}

func (rn *RuneNode) Data() rune {
	return rn.data
}

func (rn *RuneNode) SetData(v rune) {
	rn.data = v
}

func (rn *RuneNode) Next() *RuneNode {
	return rn.next
}

func (rn *RuneNode) SetNext(n *RuneNode) {
	rn.next = n
}

// Stack — стек на односвязном списке.
// Работает по принципу LIFO: push в начало, pop с начала.
type Stack struct {
	top *RuneNode
}

func NewStack(n *RuneNode) *Stack {
	return &Stack{
		top: n,
	}
}

func (s *Stack) Top() *RuneNode {
	return s.top
}

func (s *Stack) SetTop(rn *RuneNode) {
	s.top = rn
}

// Push кладёт элемент на вершину стека.
func (s *Stack) Push(v rune) {
	newTop := NewRuneNode(v, s.top)
	s.SetTop(newTop)
}

// Pop снимает элемент с вершины стека.
// Второй возвращаемый параметр говорит, был ли элемент.
func (s *Stack) Pop() (rune, bool) {
	if s.Top() == nil {
		return 0, false
	}
	val := s.Top().Data()
	s.top = s.Top().Next()
	return val, true
}

// Peek возвращает значение на вершине, но не удаляет.
func (s *Stack) Peek() (rune, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.Top().Data(), true
}

// Empty сообщает, пуст ли стек.
func (s *Stack) Empty() bool {
	return s.top == nil
}
