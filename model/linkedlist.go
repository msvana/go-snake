package model

type Node[T any] struct {
	Value T
	Next  int
	Prev  int
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value, -1, -1}
}

type LinkedList[T any] struct {
	Nodes     []Node[T]
	HeadIndex int
	TailIndex int
}

func NewLinkedList[T any](size uint) *LinkedList[T] {
	nodes := make([]Node[T], 0, size)
	return &LinkedList[T]{nodes, -1, -1}
}

func (ll *LinkedList[T]) Add(node Node[T]) {
	ll.Nodes = append(ll.Nodes, node)
	if ll.HeadIndex == -1 {
		ll.HeadIndex = 0
		ll.TailIndex = 0
	} else {
		ll.Nodes[ll.TailIndex].Next = len(ll.Nodes) - 1
		ll.Nodes[len(ll.Nodes)-1].Prev = ll.TailIndex
		ll.TailIndex = len(ll.Nodes) - 1
	}
}

func (ll *LinkedList[T]) Head() *Node[T] {
	return &ll.Nodes[ll.HeadIndex]
}

func (ll *LinkedList[T]) Tail() *Node[T] {
	return &ll.Nodes[ll.TailIndex]
}
