package data_structures

type linked_node[T any] struct {
	Value T
	Next  *linked_node[T]
}

type SinglyLinkedList[T any] struct {
	length int
	head   *linked_node[T]
	tail   *linked_node[T]
}

func (l *SinglyLinkedList[T]) Length() int {
	return l.length
}

func (l *SinglyLinkedList[T]) Append(v T) {
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

func (l *SinglyLinkedList[T]) Prepend(v T) {
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

func (l *SinglyLinkedList[T]) Shift() linked_node[T] {
	n := *l.head
	l.head = n.Next
	l.length--
	return n
}
