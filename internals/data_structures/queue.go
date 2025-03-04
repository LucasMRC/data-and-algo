package data_structures

type IQueue[T any] interface {
	Enqueue(T)
	Deque() linked_node[T]
	Peek() T
}

type Queue[T any] struct {
	head   *linked_node[T]
	tail   *linked_node[T]
	length int
}

func (q *Queue[T]) Enqueue(v T) {
	n := linked_node[T]{
		Value: v,
	}
	if q.length == 0 {
		q.head = &n
	} else {
		q.tail.Next = &n
	}
	q.tail = &n
	q.length++
}

func (q *Queue[T]) Deque() linked_node[T] {
	n := q.head
	var head *linked_node[T]
	if q.length > 1 {
		head = q.head.Next
	}
	q.head = head
	return *n
}

func (q *Queue[T]) Peek() T {
	return q.head.Value
}
