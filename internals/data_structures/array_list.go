package data_structures

type ArrayList[T comparable] struct {
	Length   int
	capacity int
	content  []T
}

func (a *ArrayList[T]) Append(v T) {
	if a.Length == a.capacity {
		*a = allocateMemory(a)
	}
	var x T
	a.content = append(a.content, x)
	a.content[a.Length] = v
	a.Length++
}

func (a *ArrayList[T]) Prepend(v T) {
	if a.Length == a.capacity {
		*a = allocateMemory(a)
	}
	a.Length++ // we do this earlier as we'll use the updated Length in the loop
	prev := a.content[0]
	var x T
	a.content = append(a.content, x)
	for i := 1; i < a.Length; i++ {
		temp := a.content[i]
		a.content[i] = prev
		prev = temp
	}
	a.content[0] = v
}

func (a *ArrayList[T]) InsertAt(i int, v T) {
	if a.Length == a.capacity {
		*a = allocateMemory(a)
	}
	a.Length++ // we do this earlier as we'll use the updated Length in the loop
	prev := a.content[i]
	var x T
	a.content = append(a.content, x)
	for j := i + 1; j < a.Length; j++ {
		temp := a.content[i]
		a.content[j] = prev
		prev = temp
	}
	a.content[i] = v
}

func (a *ArrayList[T]) Get(i int) T {
	var v T
	if i < a.Length {
		v = a.content[i]
	}
	return v
}

func (a *ArrayList[T]) RemoveAt(i int) T {
	var v T
	if i < a.Length {
		v = a.content[i]
		for j := i; j+1 < len(a.content); j++ {
			a.content[j] = a.content[j+1]
		}
		a.Length--
	}
	return v
}

func (a *ArrayList[T]) Remove(v T) T {
	var idx int
	found := false
	for i := range a.Length {
		curr := a.content[i]
		if v == curr {
			found = true
			idx = i
			break
		}
	}
	var x T
	if found {
		x = a.content[idx]
		for j := idx; j+1 < len(a.content); j++ {
			a.content[j] = a.content[j+1]
		}
		a.Length--
	}
	return x
}

// ---

func allocateMemory[T comparable](a *ArrayList[T]) ArrayList[T] {
	n := ArrayList[T]{
		Length:   a.Length,
		capacity: a.capacity * 2,
		content:  a.content,
	}
	return n
}
