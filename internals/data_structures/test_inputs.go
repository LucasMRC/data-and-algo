package data_structures

import (
	"github.com/LucasMRC/kata-machine-go/internals/test_utils"
)

type node struct {
	value int
}

type testResult struct {
	length int
	head   node
	tail   node
}

var LinkedList_TestCases = []test_utils.ActionTestCase[LinkedList[int], testResult]{
	{
		Name:     "Empty list",
		Action:   func(l *LinkedList[int]) {},
		Expected: testResult{},
	},
	{
		Name: "Append four",
		Action: func(l *LinkedList[int]) {
			l.Append(4)
			l.Append(3)
			l.Append(2)
			l.Append(1)
		},
		Expected: testResult{
			length: 4,
			head: node{
				value: 4,
			},
			tail: node{
				value: 1,
			},
		},
	},
	{
		Name: "Prepend two",
		Action: func(l *LinkedList[int]) {
			l.Prepend(5)
			l.Prepend(6)
		},
		Expected: testResult{
			length: 6,
			head: node{
				value: 6,
			},
			tail: node{
				value: 1,
			},
		},
	},
	{
		Name: "Shift two",
		Action: func(l *LinkedList[int]) {
			l.Shift()
			l.Shift()
		},
		Expected: testResult{
			length: 4,
			head: node{
				value: 4,
			},
			tail: node{
				value: 1,
			},
		},
	},
}

var DoubleLinkedList_TestCases = []test_utils.ActionTestCase[DoubleLinkedList[int], testResult]{
	{
		Name:     "Empty list",
		Action:   func(l *DoubleLinkedList[int]) {},
		Expected: testResult{},
	},
	{
		Name: "Append four",
		Action: func(l *DoubleLinkedList[int]) {
			l.Append(4)
			l.Append(3)
			l.Append(2)
			l.Append(1)
		},
		Expected: testResult{
			length: 4,
			head: node{
				value: 4,
			},
			tail: node{
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
		Expected: testResult{
			length: 6,
			head: node{
				value: 6,
			},
			tail: node{
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
		Expected: testResult{
			length: 3,
			head: node{
				value: 6,
			},
			tail: node{
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
		Expected: testResult{
			length: 1,
			head: node{
				value: 4,
			},
			tail: node{
				value: 4,
			},
		},
	},
}

var Queue_TestCases = []test_utils.ActionTestCase[Queue[int], int]{
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

var Stack_TestCases = []test_utils.ActionTestCase[Stack[int], int]{
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

var ArrayList_TestCases = []test_utils.ActionReturnTestCase[ArrayList[int], []int]{
	{
		Name: "First",
		Action: func(a *ArrayList[int]) []int {
			a.Append(5)
			a.Append(7)
			a.Append(9)

			result1 := a.Get(2)
			result2 := a.RemoveAt(1)
			result3 := a.Length

			return []int{result1, result2, result3}
		},
		Expected: []int{9, 7, 2},
	},
	{
		Name: "Second",
		Action: func(a *ArrayList[int]) []int {
			a.Append(11)

			result1 := a.RemoveAt(1)
			result2 := a.Remove(9)
			result3 := a.RemoveAt(0)
			result4 := a.RemoveAt(0)
			result5 := a.Length

			return []int{result1, result2, result3, result4, result5}
		},
		Expected: []int{9, 0, 5, 11, 0},
	},
	{
		Name: "Third",
		Action: func(a *ArrayList[int]) []int {
			a.Prepend(5)
			a.Prepend(7)
			a.Prepend(9)

			result1 := a.Get(2)
			result2 := a.Get(0)
			result3 := a.Remove(9)
			result4 := a.Length
			result5 := a.Get(0)

			return []int{result1, result2, result3, result4, result5}
		},
		Expected: []int{5, 9, 9, 2, 7},
	},
}

var RingBuffer_TestCases = []test_utils.ActionReturnTestCase[RingBuffer[int], []int]{
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
}
