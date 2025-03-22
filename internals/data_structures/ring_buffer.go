package data_structures

type RingBuffer[T comparable] struct {
	capacity int
	head     int
	tail     int
	content  []T
	Length   int
}

func (r *RingBuffer[T]) Initialize(cap int) {
	r.capacity = cap
	r.content = make([]T, cap)
	r.head = len(r.content) / 5
	r.tail = r.head
}

func (a *RingBuffer[T]) Append(v T) {
	tail := a.tail % a.capacity
	if a.head == tail && a.Length > 0 {
		content := make([]T, a.capacity*2)
		for i := range a.Length {
			v := a.content[(a.head+i)%a.capacity]
			if v != *new(T) {
				content[i] = v
			}
		}
		n := RingBuffer[T]{
			capacity: a.capacity * 2,
			head:     0,
			tail:     a.Length,
			Length:   a.Length,
			content:  content,
		}
		*a = n
	}
	a.content[a.tail%a.capacity] = v
	a.tail++
	a.Length++
}

func (a *RingBuffer[T]) Pop() T {
	var v T
	if a.Length > 0 {
		var temp T
		v = a.content[a.head]
		a.content[a.head] = temp
		a.head++
		a.Length--
	}
	if a.head == a.capacity {
		a.head = a.head % a.capacity
		a.tail = a.tail % a.capacity
	}
	return v
}

func (a *RingBuffer[T]) Get(i int) T {
	var v T
	idx := (a.head + i) % a.capacity
	if idx < a.tail {
		v = a.content[idx]
	}
	return v
}
