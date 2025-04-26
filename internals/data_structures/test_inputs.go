package data_structures

import (
	"github.com/LucasMRC/kata-machine-go/internals/test_utils"
)

type test_node struct {
	value int
}

type test_result struct {
	length int
	head   test_node
	tail   test_node
}

var SinglyLinkedList_TestCases = []test_utils.ChainedActionTestCase[SinglyLinkedList[int], test_result]{
	{
		Name:     "Empty list",
		Action:   func(l *SinglyLinkedList[int]) {},
		Expected: test_result{},
	},
	{
		Name: "Append four",
		Action: func(l *SinglyLinkedList[int]) {
			l.Append(4)
			l.Append(3)
			l.Append(2)
			l.Append(1)
		},
		Expected: test_result{
			length: 4,
			head: test_node{
				value: 4,
			},
			tail: test_node{
				value: 1,
			},
		},
	},
	{
		Name: "Prepend two",
		Action: func(l *SinglyLinkedList[int]) {
			l.Prepend(5)
			l.Prepend(6)
		},
		Expected: test_result{
			length: 6,
			head: test_node{
				value: 6,
			},
			tail: test_node{
				value: 1,
			},
		},
	},
	{
		Name: "Shift two",
		Action: func(l *SinglyLinkedList[int]) {
			l.Shift()
			l.Shift()
		},
		Expected: test_result{
			length: 4,
			head: test_node{
				value: 4,
			},
			tail: test_node{
				value: 1,
			},
		},
	},
}

var DoubleLinkedList_TestCases = []test_utils.ChainedActionTestCase[DoubleLinkedList[int], test_result]{
	{
		Name:     "Empty list",
		Action:   func(l *DoubleLinkedList[int]) {},
		Expected: test_result{},
	},
	{
		Name: "Append four",
		Action: func(l *DoubleLinkedList[int]) {
			l.Append(4)
			l.Append(3)
			l.Append(2)
			l.Append(1)
		},
		Expected: test_result{
			length: 4,
			head: test_node{
				value: 4,
			},
			tail: test_node{
				value: 1,
			},
		},
	},
	{
		Name: "Prepend two",
		Action: func(l *DoubleLinkedList[int]) {
			l.Prepend(5)
			l.Prepend(6)
		},
		Expected: test_result{
			length: 6,
			head: test_node{
				value: 6,
			},
			tail: test_node{
				value: 1,
			},
		},
	},
	{
		Name: "Pop three",
		Action: func(l *DoubleLinkedList[int]) {
			l.Pop()
			l.Pop()
			l.Pop()
		},
		Expected: test_result{
			length: 3,
			head: test_node{
				value: 6,
			},
			tail: test_node{
				value: 4,
			},
		},
	},
	{
		Name: "Shift two",
		Action: func(l *DoubleLinkedList[int]) {
			l.Shift()
			l.Shift()
		},
		Expected: test_result{
			length: 1,
			head: test_node{
				value: 4,
			},
			tail: test_node{
				value: 4,
			},
		},
	},
}

var Queue_TestCases = []test_utils.ChainedActionTestCase[Queue[int], int]{
	{
		Name: "Enqueue four",
		Action: func(l *Queue[int]) {
			l.Enqueue(4)
			l.Enqueue(3)
			l.Enqueue(2)
			l.Enqueue(1)
		},
		Expected: 4,
	},
	{
		Name: "Deque two",
		Action: func(l *Queue[int]) {
			l.Deque()
			l.Deque()
		},
		Expected: 2,
	},
	{
		Name: "Enqueue three",
		Action: func(l *Queue[int]) {
			l.Enqueue(3)
			l.Enqueue(4)
			l.Enqueue(1)
		},
		Expected: 2,
	},
	{
		Name: "Deque one",
		Action: func(l *Queue[int]) {
			l.Deque()
		},
		Expected: 1,
	},
}

var Stack_TestCases = []test_utils.ChainedActionTestCase[Stack[int], int]{
	{
		Name: "Push four",
		Action: func(l *Stack[int]) {
			l.Push(4)
			l.Push(3)
			l.Push(2)
			l.Push(1)
		},
		Expected: 1,
	},
	{
		Name: "Pop two",
		Action: func(l *Stack[int]) {
			l.Pop()
			l.Pop()
		},
		Expected: 3,
	},
	{
		Name: "Push three",
		Action: func(l *Stack[int]) {
			l.Push(3)
			l.Push(4)
			l.Push(1)
		},
		Expected: 1,
	},
	{
		Name: "Pop one",
		Action: func(l *Stack[int]) {
			l.Pop()
		},
		Expected: 4,
	},
}

var ArrayList_TestCases = []test_utils.ChainedActionReturnTestCase[ArrayList[int], []int]{
	{
		Name: "First",
		Action: func(a *ArrayList[int]) []int {
			a.Append(5)
			a.Append(7)
			a.Append(9)

			r1 := a.Get(2)
			r2 := a.RemoveAt(1)
			r3 := a.Length

			return []int{r1, r2, r3}
		},
		Expected: []int{9, 7, 2},
	},
	{
		Name: "Second",
		Action: func(a *ArrayList[int]) []int {
			a.Append(11)

			r1 := a.RemoveAt(1)
			r2 := a.Remove(9)
			r3 := a.RemoveAt(0)
			r4 := a.RemoveAt(0)
			r5 := a.Length

			return []int{r1, r2, r3, r4, r5}
		},
		Expected: []int{9, 0, 5, 11, 0},
	},
	{
		Name: "Third",
		Action: func(a *ArrayList[int]) []int {
			a.Prepend(5)
			a.Prepend(7)
			a.Prepend(9)

			r1 := a.Get(2)
			r2 := a.Get(0)
			r3 := a.Remove(9)
			r4 := a.Length
			r5 := a.Get(0)

			return []int{r1, r2, r3, r4, r5}
		},
		Expected: []int{5, 9, 9, 2, 7},
	},
}

var RingBuffer_TestCases = []test_utils.ChainedActionReturnTestCase[RingBuffer[int], []int]{
	{
		Name: "First",
		Action: func(r *RingBuffer[int]) []int {
			r.Append(5)

			r1 := r.Pop()
			r2 := r.Pop()

			return []int{r1, r2}
		},
		Expected: []int{5, 0},
	},
	{
		Name: "Second",
		Action: func(r *RingBuffer[int]) []int {
			r.Append(42)
			r.Append(9)

			r1 := r.Pop()
			r2 := r.Pop()
			r3 := r.Pop()

			return []int{r1, r2, r3}
		},
		Expected: []int{42, 9, 0},
	},
	{
		Name: "Third",
		Action: func(r *RingBuffer[int]) []int {
			r.Append(42)
			r.Append(9)
			r.Append(12)

			r1 := r.Get(2)
			r2 := r.Get(1)
			r3 := r.Get(0)

			return []int{r1, r2, r3}
		},
		Expected: []int{12, 9, 42},
	},
	{
		Name: "Fourth",
		Action: func(r *RingBuffer[int]) []int {
			r.Append(10)
			r.Append(73)
			r.Append(25)
			r.Append(44)

			r1 := r.Length
			r2 := r.capacity
			r3 := r.Get(r.head)
			r4 := r.Get(r.tail - 1)

			return []int{r1, r2, r3, r4}
		},
		Expected: []int{7, 10, 42, 44},
	},
	{
		Name: "Fifth",
		Action: func(r *RingBuffer[int]) []int {
			r.Append(14)
			r.Append(3)
			r.Append(83)
			r.Append(1)
			r.Append(35)
			r.Append(57)

			r1 := r.Pop()
			r2 := r.Pop()
			r3 := r.head
			r4 := r.tail

			return []int{r1, r2, r3, r4}
		},
		Expected: []int{42, 9, 2, 13},
	},
}

func InitializeRingBuffer() RingBuffer[int] {
	r := RingBuffer[int]{
		capacity: 5,
		content: make([]int, 5),
		head: 0,
		tail: 0,
	}
	return r
}

var tree_one = Binary_Node[int]{
	Value: 20,
	Right: &Binary_Node[int]{
		Value: 50,
		Right: &Binary_Node[int]{
			Value: 100,
			Right: &Binary_Node[int]{},
			Left: &Binary_Node[int]{},
		},
		Left: &Binary_Node[int]{
			Value: 30,
			Right: &Binary_Node[int]{
				Value: 45,
				Right: &Binary_Node[int]{},
				Left: &Binary_Node[int]{},
			},
			Left: &Binary_Node[int]{
				Value: 29,
				Right: &Binary_Node[int]{},
				Left: &Binary_Node[int]{},
			},
		},
	},
	Left: &Binary_Node[int]{
		Value: 10,
		Right: &Binary_Node[int]{
			Value: 15,
			Right: &Binary_Node[int]{},
			Left: &Binary_Node[int]{},
		},
		Left: &Binary_Node[int]{
			Value: 5,
			Right: &Binary_Node[int]{
				Value: 7,
				Right: &Binary_Node[int]{},
				Left: &Binary_Node[int]{},
			},
			Left: &Binary_Node[int]{},
		},
	},
}

var tree_two = Binary_Node[int]{
	Value: 20,
	Right: &Binary_Node[int]{
		Value: 50,
		Right: &Binary_Node[int]{},
		Left: &Binary_Node[int]{
			Value: 30,
			Right: &Binary_Node[int]{
				Value: 45,
				Right: &Binary_Node[int]{
					Value: 49,
					Right: &Binary_Node[int]{},
					Left: &Binary_Node[int]{},
				},
				Left: &Binary_Node[int]{},
			},
			Left: &Binary_Node[int]{
				Value: 29,
				Right: &Binary_Node[int]{},
				Left: &Binary_Node[int]{
					Value: 21,
					Right: &Binary_Node[int]{},
					Left: &Binary_Node[int]{},
				},
			},
		},
	},
	Left: &Binary_Node[int]{
		Value: 10,
		Right: &Binary_Node[int]{
			Value: 15,
			Right: &Binary_Node[int]{},
			Left: &Binary_Node[int]{},
		},
		Left: &Binary_Node[int]{
			Value: 5,
			Right: &Binary_Node[int]{
				Value: 7,
				Right: &Binary_Node[int]{},
				Left: &Binary_Node[int]{},
			},
			Left: &Binary_Node[int]{},
		},
	},
}

var Tree_TestCases = []test_utils.ActionReturnTestCase[Binary_Node[int], []int]{
	{
		Name: "PreOrderSearch 1",
		Action: func() []int {
			return tree_one.PreOrderSearch()
		},
		Expected: []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100},
	},
	{
		Name: "PreOrderSearch 2",
		Action: func() []int {
			return tree_two.PreOrderSearch()
		},
		Expected: []int{20, 10, 5, 7, 15, 50, 30, 29, 21, 45, 49},
	},
	{
		Name: "InOrderSearch 1",
		Action: func() []int {
			return tree_one.InOrderSearch()
		},
		Expected: []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100},
	},
	{
		Name: "InOrderSearch 2",
		Action: func() []int {
			return tree_two.InOrderSearch()
		},
		Expected: []int{5, 7, 10, 15, 20, 21, 29, 30, 45, 49, 50},
	},
	{
		Name: "PostOrderSearch 1",
		Action: func() []int {
			return tree_one.PostOrderSearch()
		},
		Expected: []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20},
	},
	{
		Name: "PostOrderSearch 2",
		Action: func() []int {
			return tree_two.PostOrderSearch()
		},
		Expected: []int{7, 5, 15, 10, 21, 29, 49, 45, 30, 50, 20},
	},
}

