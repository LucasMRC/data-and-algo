package data_structures

import "errors"

type linked_node[T any] struct {
	Value T
	Next  *linked_node[T]
}

type LinkedList[T any] struct {
	length int
	head   *linked_node[T]
	tail   *linked_node[T]
}

func (l *LinkedList[T]) Length() int {
	return l.length
}

func (l *LinkedList[T]) Append(v T) {
	n := linked_node[T]{
		Value: v,
	}
	if l.Length() == 0 {
		l.head = &n
	} else {
		l.tail.Next = &n
	}
	l.tail = &n
	l.length++
}

func (l *LinkedList[T]) Prepend(v T) {
	n := linked_node[T]{
		Value: v,
		Next:  l.head,
	}
	if l.Length() == 0 {
		l.tail = &n
	}
	l.head = &n
	l.length++
}

func (l *LinkedList[T]) Shift() (linked_node[T], error) {
	if l.Length() == 0 {
		return linked_node[T]{}, errors.New("Cannot pop an empty list")
	}
	n := *l.head
	l.head = n.Next
	l.length--
	return n, nil
}
