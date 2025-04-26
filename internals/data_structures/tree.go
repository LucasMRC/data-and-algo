package data_structures

import "slices"

type Binary_Node[T comparable] struct {
	Value	T
	Left	*Binary_Node[T]
	Right	*Binary_Node[T]
}

func tree_search[T comparable](b Binary_Node[T], walk func(Binary_Node[T], *[]T)) []T {
	var v []T
	walk(b, &v)
	return v
}

func (b *Binary_Node[T]) PreOrderSearch() []T {
	return tree_search(*b, walk_pre_order)
}

func (b *Binary_Node[T]) InOrderSearch() []T {
	return tree_search(*b, walk_in_order)
}

func (b *Binary_Node[T]) PostOrderSearch() []T {
	return tree_search(*b, walk_post_order)
}

func walk_pre_order[T comparable](n Binary_Node[T], path *[]T) {
	var empty T
	if n.Value == empty { 
		return
	}
	if !slices.Contains(*path, n.Value) {
		*path = append(*path, n.Value)
	}
	if !slices.Contains(*path, (*n.Left).Value) {
		walk_pre_order(*n.Left, path)
	}
	if !slices.Contains(*path, (*n.Right).Value) {
		walk_pre_order(*n.Right, path)
	}
}

func walk_in_order[T comparable](n Binary_Node[T], path *[]T) {
	var empty T
	if n.Value == empty { 
		return
	}
	if !slices.Contains(*path, (*n.Left).Value) {
		walk_in_order(*n.Left, path)
	}
	if !slices.Contains(*path, n.Value) {
		*path = append(*path, n.Value)
	}
	if !slices.Contains(*path, (*n.Right).Value) {
		walk_in_order(*n.Right, path)
	}
}

func walk_post_order[T comparable](n Binary_Node[T], path *[]T) {
	var empty T
	if n.Value == empty { 
		return
	}
	if !slices.Contains(*path, (*n.Left).Value) {
		walk_post_order(*n.Left, path)
	}
	if !slices.Contains(*path, (*n.Right).Value) {
		walk_post_order(*n.Right, path)
	}
	if !slices.Contains(*path, n.Value) {
		*path = append(*path, n.Value)
	}
}
