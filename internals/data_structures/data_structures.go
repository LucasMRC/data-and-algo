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
	// case "queu":
	// TestQueu()
	case "all":
		TestLinkedList()
		TestDoubleLinkedList()
		TestStack()
	}
}
