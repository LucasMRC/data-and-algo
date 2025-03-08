package data_structures

import (
	"fmt"

	"github.com/LucasMRC/kata-machine-go/internals/test_utils"
)

var test_results = test_utils.TestResults

func TestLinkedList() {
	fmt.Println("LinkedList: running tests.")
	fmt.Println("----------------------------")
	l := LinkedList[int]{}
	var rl string
	var rh string
	var rt string
	for _, test := range LinkedList_TestCases {
		fmt.Printf("\n%s:\n", test.Name)
		test.Action(&l)
		rl = test_utils.ToEqual(l.Length(), test.Expected.length)
		if l.Length() == 0 {
			rh = test_utils.ToEqual(l.head == nil, true)
			rt = test_utils.ToEqual(l.tail == nil, true)
		} else {
			rh = test_utils.ToEqual(l.head.Value, test.Expected.head.value)
			rt = test_utils.ToEqual(l.tail.Value, test.Expected.tail.value)
		}
		fmt.Printf("\tLength: %v\n", rl)
		if rl == test_results.FAILED {
			fmt.Printf("\t  expected %v, result %v\n", test.Expected.length, rl)
		}
		fmt.Printf("\tHead: %v\n", rh)
		if rh == test_results.FAILED {
			fmt.Printf("\t  expected value %v, result %v\n", test.Expected.head.value, l.head)
		}
		fmt.Printf("\tTail: %v\n", rt)
		if rt == test_results.FAILED {
			fmt.Printf("\t  expected value %v, result %v\n", test.Expected.tail.value, l.tail)
		}
	}
	fmt.Println("============================")
}

func TestDoubleLinkedList() {
	fmt.Println("DoubleLinkedList: running tests.")
	fmt.Println("----------------------------")
	l := DoubleLinkedList[int]{}
	var rl string
	var rh string
	var rt string
	for _, test := range DoubleLinkedList_TestCases {
		fmt.Printf("\n%s:\n", test.Name)
		test.Action(&l)
		rl = test_utils.ToEqual(l.Length(), test.Expected.length)
		if l.Length() == 0 {
			rh = test_utils.ToEqual(l.head == nil, true)
			rt = test_utils.ToEqual(l.tail == nil, true)
		} else {
			rh = test_utils.ToEqual(l.head.Value, test.Expected.head.value)
			rt = test_utils.ToEqual(l.tail.Value, test.Expected.tail.value)
		}
		fmt.Printf("\tLength: %v\n", rl)
		if rl == test_results.FAILED {
			fmt.Printf("\t  expected %v, result %v\n", test.Expected.length, rl)
		}
		fmt.Printf("\tHead: %v\n", rh)
		if rh == test_results.FAILED {
			fmt.Printf("\t  expected value %v, result %v\n", test.Expected.head.value, l.head.Value)
		}
		fmt.Printf("\tTail: %v\n", rt)
		if rt == test_results.FAILED {
			fmt.Printf("\t  expected value %v, result %v\n", test.Expected.tail.value, l.tail.Value)
		}
	}
	fmt.Println("============================")
}

func TestQueue() {
	fmt.Println("Queue: running tests.")
	fmt.Println("----------------------------")
	q := Queue[int]{}
	for _, test := range Queue_TestCases {
		test.Action(&q)
		result := q.Peek()
		test_status := test_utils.ToEqual(result, test.Expected)
		fmt.Printf("%s: %v\n", test.Name, test_status)
		if test_status == test_results.FAILED {
			fmt.Printf("\texpected %v, result %v\n", test.Expected, result)
		}
	}
	fmt.Println("============================")
}

func TestStack() {
	fmt.Println("Stack: running tests.")
	fmt.Println("----------------------------")
	q := Stack[int]{}
	for _, test := range Stack_TestCases {
		test.Action(&q)
		result := q.Peek()
		test_status := test_utils.ToEqual(result, test.Expected)
		fmt.Printf("%s: %v\n", test.Name, test_status)
		if test_status == test_results.FAILED {
			fmt.Printf("\texpected %v, result %v\n", test.Expected, result)
		}
	}
	fmt.Println("============================")
}

func TestArrayList() {
	fmt.Println("ArrayList: running tests.")
	fmt.Println("----------------------------")
	q := ArrayList[int]{}
	for _, test := range ArrayList_TestCases {
		result := test.Action(&q)
		test_status := test_utils.SlicesToEqual(result, test.Expected)
		fmt.Printf("%s: %v\n", test.Name, test_status)
		if test_status == test_results.FAILED {
			fmt.Printf("\texpected %v, result %v\n", test.Expected, result)
		}
	}
	fmt.Println("============================")
}

func TestRingBuffer() {
	fmt.Println("RingBuffer: running tests.")
	fmt.Println("----------------------------")
	q := RingBuffer[int]{}
	q.Initialize(5)
	for _, test := range RingBuffer_TestCases {
		result := test.Action(&q)
		test_status := test_utils.SlicesToEqual(result, test.Expected)
		fmt.Printf("%s: %v\n", test.Name, test_status)
		if test_status == test_results.FAILED {
			fmt.Printf("\texpected %v, result %v\n", test.Expected, result)
		}
	}
	fmt.Println("============================")
}
