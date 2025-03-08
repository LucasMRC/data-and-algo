package data_structures

import "fmt"

func Run(f string) {
	fmt.Println("============================")

	switch f {
	case "linked":
		TestLinkedList()
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
		TestLinkedList()
		TestDoubleLinkedList()
		TestStack()
		TestQueue()
		TestArrayList()
		TestRingBuffer()
	}
}
