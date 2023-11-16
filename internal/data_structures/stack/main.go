package stack

import "math"

type Node[T any] struct {
	Value T
	Prev  *Node[T]
}

type Stack[T any] struct {
	Length int
	head   *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	node := &Node[T]{Value: item}

	s.Length = s.Length + 1

	if s.head == nil {
		s.head = node
		return
	}

	node.Prev = s.head
	s.head = node
}

func (s *Stack[T]) Pop() T {
	s.Length = int(math.Max(0, float64(s.Length-1)))

	if s.Length == 0 {
		head := s.head
		s.head = nil

		return head.Value
	}

	head := s.head
	s.head = head.Prev

	return head.Value
}

func (s *Stack[T]) Peek() T {
	return s.head.Value
}
