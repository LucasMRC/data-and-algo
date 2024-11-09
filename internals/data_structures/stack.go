package data_structures

import "errors"

type IStack[T any] interface {
	Push(T)
	Pop() T
	Peek() T
}

type Stack[T any] struct {
	Length int
	head   *linked_node[T]
}

func (q *Stack[T]) Push(v T) {
	n := linked_node[T]{
		Value: v,
		Next:  q.head,
	}
	q.head = &n
	q.Length++
}

func (q *Stack[T]) Pop() (T, error) {
	if q.Length == 0 {
		return *new(T), errors.New("Invalid operation on empty queue")
	}
	n := q.head
	q.Length--
	if q.Length == 0 {
		q.head = nil
	} else {
		q.head = q.head.Next
		n.Next = nil
	}
	return n.Value, nil
}

func (q *Stack[T]) Peek() T {
	return q.head.Value
}
