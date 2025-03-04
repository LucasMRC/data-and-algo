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
		Name: "Appended 4",
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
		Name: "Prepended 2",
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
		Name: "Shifted two",
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
		Name: "Appended four",
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
		Name: "Prepended two",
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
		Name: "Poped three",
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
		Name: "Shifted two",
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
