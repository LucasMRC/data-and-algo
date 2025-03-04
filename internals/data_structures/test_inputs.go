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
