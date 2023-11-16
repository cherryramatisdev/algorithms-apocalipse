package queue

type Node[T any] struct {
	Value T
	// Next is optional
	Next *Node[T]
}

type Queue[T any] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Length: 0,
	}
}

func (q *Queue[T]) Enqueue(item T) {
	node := &Node[T]{Value: item}
	q.Length = q.Length + 1

	if q.tail == nil {
		q.head = node
		q.tail = q.head
	}

	q.tail.Next = node
	q.tail = node
}

func (q *Queue[T]) Deque() T {
	if q.head == nil {
		var zero_value T

		return zero_value
	}

	q.Length = q.Length - 1

	head := q.head
	q.head = head.Next

	return head.Value
}

func (q *Queue[T]) Peek() T {
	return q.head.Value
}
