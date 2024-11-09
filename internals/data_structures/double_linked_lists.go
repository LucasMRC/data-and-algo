package data_structures

type double_linked_node[T any] struct {
	Value T
	Next  *double_linked_node[T]
	Prev  *double_linked_node[T]
}

type DoubleLinkedList[T any] struct {
	length int
	head   *double_linked_node[T]
	tail   *double_linked_node[T]
}

func (l *DoubleLinkedList[T]) Length() int {
	return l.length
}

func (l *DoubleLinkedList[T]) Append(v T) {
	n := double_linked_node[T]{
		Value: v,
		Prev:  l.tail,
	}
	if l.tail != nil {
		l.tail.Next = &n
	}
	l.tail = &n
	if l.head == nil {
		l.head = &n
	}
	l.length++
}

func (l *DoubleLinkedList[T]) Prepend(v T) {
	n := double_linked_node[T]{
		Value: v,
		Next:  l.head,
	}
	if l.head != nil {
		l.head.Next = &n
	}
	l.head = &n
	if l.tail == nil {
		l.tail = &n
	}
	l.length++
}

func (l *DoubleLinkedList[T]) InsertAt(v T, i int) {
	cn := l.tail
	for j := 0; j < i; j++ {
		cn = cn.Next
	}
	n := double_linked_node[T]{
		Value: v,
		Prev:  cn,
		Next:  cn.Next,
	}
	cn.Next = &n
}

func (l *DoubleLinkedList[T]) Pop() double_linked_node[T] {
	n := l.tail
	l.tail = n.Prev
	n.Prev = nil
	l.length--
	return *n
}

func (l *DoubleLinkedList[T]) Shift() double_linked_node[T] {
	n := l.head
	l.head = n.Next
	n.Next = nil
	l.length--
	return *n
}

func (l *DoubleLinkedList[T]) RemoveAt(i int) double_linked_node[T] {
	cn := l.tail
	for j := 0; j <= i; j++ {
		cn = cn.Next
	}
	cn.Next.Prev = cn.Prev
	cn.Prev.Next = cn.Next
	cn.Prev = nil
	cn.Next = nil
	return *cn
}

func (l *DoubleLinkedList[T]) Get(i int) double_linked_node[T] {
	cn := l.tail
	for j := 0; j <= i; j++ {
		cn = cn.Next
	}
	return *cn
}
