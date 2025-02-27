package data_structures

type IQueue[T any] interface {
	Enqueue(T)
	Deque() double_linked_node[T]
	Peek() T
}

type Queue[T any] struct {
	DoubleLinkedList[T]
}

func (q *Queue[T]) Enqueue(v T) {
	q.Append(v)
}

func (q *Queue[T]) Deque() double_linked_node[T] {
	return q.Pop()
}

func (q *Queue[T]) Peek() T {
	return q.head.Value
}
