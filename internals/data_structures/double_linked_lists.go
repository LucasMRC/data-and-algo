package data_structures

type double_linked_node[T comparable] struct {
	Value T
	Next  *double_linked_node[T]
	Prev  *double_linked_node[T]
}

type DoubleLinkedList[T comparable] struct {
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
		l.head.Prev = &n
	}
	l.head = &n
	if l.tail == nil {
		l.tail = &n
	}
	l.length++
}

func (l *DoubleLinkedList[T]) InsertAt(v T, i int) {
	cn := l.tail
	for range i {
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
	for range i {
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
	for range i {
		cn = cn.Next
	}
	return *cn
}

func (l *DoubleLinkedList[T]) Remove(n T) T {
	var tmp T
	c := l.head
	for i := range l.Length() {
		if c.Value == n {
			tmp = l.RemoveAt(i).Value
		} else {
			c = c.Next
		}
	}
	return tmp
} 
