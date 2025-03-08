package data_structures

import "fmt"

func Run(f string) {
	fmt.Println("============================")

	switch f {
	case "singly":
		TestSinglyLinkedList()
	case "double":
		TestDoubleLinkedList()
	case "stack":
		TestStack()
	case "queue":
		TestQueue()
	case "arraylist":
		TestArrayList()
	case "ring":
		TestRingBuffer()
	case "all":
		TestSinglyLinkedList()
		TestDoubleLinkedList()
		TestStack()
		TestQueue()
		TestArrayList()
		TestRingBuffer()
	}
}
