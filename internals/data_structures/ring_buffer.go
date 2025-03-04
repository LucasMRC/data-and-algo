package data_structures

type RingBuffer[T any] struct {
	capacity int
	head     int
	tail     int
	content  []T
}

func (a *RingBuffer[T]) Push(v T) {
	if a.head == a.tail {
		n := RingBuffer[T]{
			capacity: a.capacity * 2,
			head:     a.head,
			tail:     a.tail,
			content:  a.content,
		}
		a = &n
	}
	a.content[a.tail] = v
	a.tail++
}

func (a *RingBuffer[T]) Pop() T {
	var v T
	if a.tail == 0 {
		var n T
		v = a.content[a.tail]
		a.content[a.tail] = n
		a.tail--
	}
	return v
}

func (a *RingBuffer[T]) Shift(v T) {
	if a.head == a.tail {
		n := RingBuffer[T]{
			capacity: a.capacity * 2,
			head:     a.head,
			tail:     a.tail,
			content:  a.content,
		}
		a = &n
	}
	a.head--
	a.content[a.head] = v
}

func (a *RingBuffer[T]) Unshift() T {
	var v T
	if a.tail > 0 {
		var n T
		v = a.content[a.head]
		a.content[a.head] = n
		a.head++
	}
	return v
}

func (a *RingBuffer[T]) Get(i int) T {
	var v T
	if i < a.tail%a.capacity {
		v = a.content[a.head+i]
	}
	return v
}
