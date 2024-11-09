package data_structures

import "errors"

type IQueue[T any] interface {
	Enqueue(T)
	Deque() T
	Peek() T
}

type Queue[T any] struct {
	Length int
	head   *linked_node[T]
	tail   *linked_node[T]
}

func (q *Queue[T]) Enqueue(v T) {
	n := linked_node[T]{
		Value: v,
	}
	if q.Length == 0 {
		q.head = &n
	}
	q.tail.Next = &n
	q.tail = &n
	q.Length++
}

func (q *Queue[T]) Deque() (T, error) {
	if q.Length == 0 {
		return *new(T), errors.New("Invalid operation on empty queue")
	}
	n := q.head
	q.Length--
	if q.Length == 0 {
		q.tail = nil
		q.head = nil
	} else {
		q.head = q.head.Next
		n.Next = nil
	}
	return n.Value, nil
}

func (q *Queue[T]) Peek() T {
	return q.head.Value
}
