package data_structures

type Array[T any] struct {
	Length   int
	Capacity int
	content  []T
}

func (a *Array[T]) Write(i int, v T) {
	if i > a.Capacity {
		return
	}
	a.content[i] = v
}

func (a *Array[T]) Delete(i int) T {
	var v T
	if i < a.Length {
		var n T
		v = a.content[i]
		a.content[i] = n
	}
	return v
}

func (a *Array[T]) Get(i int) T {
	var v T
	if i < a.Length {
		v = a.content[i]
	}
	return v
}
