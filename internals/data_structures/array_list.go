package data_structures

type ArrayList[T any] struct {
	Length   int
	capacity int
	content  []T
}

func (a *ArrayList[T]) Push(v T) {
	if a.Length == a.capacity {
		n := ArrayList[T]{
			Length:   a.Length,
			capacity: a.capacity * 2,
			content:  a.content,
		}
		a = &n
	}
	a.content[a.Length] = v
	a.Length++
}

func (a *ArrayList[T]) Pop() T {
	var v T
	if a.Length > 0 {
		var n T
		a.Length--
		v = a.content[a.Length]
		a.content[a.Length] = n
	}
	return v
}

func (a *ArrayList[T]) Get(i int) T {
	return a.content[i]
}
