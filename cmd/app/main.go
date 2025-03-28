package main

import (
	"fmt"
	"os"
	"slices"

	algo "github.com/LucasMRC/kata-machine-go/internals/algorithms"
	data "github.com/LucasMRC/kata-machine-go/internals/data_structures"
	ex "github.com/LucasMRC/kata-machine-go/internals/exercises"
)

func main() {
	argCount := len(os.Args)

	var what, run string
	if argCount == 1 {
		what = "all"
		run = "all"
	} else if argCount == 2 {
		if !slices.Contains([]string{"algo", "data", "all", "exer"}, os.Args[1]) {
			fmt.Printf(`
Invalid argument '%s'.

Valid options:
	- %s (algorithms)
	- %s (data structures)
	- %s (exercises)

`,
				os.Args[1], "algo", "data", "exer")
			os.Exit(1)
		}
		what = os.Args[1]
		run = "all"
	} else if argCount == 3 {
		switch os.Args[1] {
		case "algo":
			if !slices.Contains([]string{"linear", "binary", "bubble", "all"}, os.Args[2]) {
				fmt.Printf(`
Invalid argument '%s'.

Valid options:
	- %s (Linear)
	- %s (Binary)
	- %s (Bubble sort)

`,
					os.Args[2], "linear", "binary", "bubble")
				os.Exit(1)
			}
		case "data":
			if !slices.Contains([]string{"singly", "double", "stack", "queue", "arraylist", "all", "ring"}, os.Args[2]) {
				fmt.Printf(`
Invalid argument '%s'.

Valid options:
	- %s (Singly Linked List)
	- %s (Double Linked List)
	- %s (Stack)
	- %s (Queue)
	- %s (ArrayList),
	- %s (RingBuffer)

`,
					os.Args[2], "singly", "double", "stack", "queue", "arraylist", "ring")
				os.Exit(1)
			}
		case "exer":
			if !slices.Contains([]string{"crystall", "maze", "all"}, os.Args[2]) {
				fmt.Printf(`
Invalid argument '%s'.

Valid options:
	- %s (Crystall Balls)
	- %s (Maze solver)

`,
					os.Args[2], "crystall", "maze")
				os.Exit(1)
			}
		}
		what = os.Args[1]
		run = os.Args[2]
	}

	switch what {
	case "algo":
		algo.Run(run)
	case "data":
		data.Run(run)
	case "exer":
		ex.Run(run)
	case "all":
		algo.Run(run)
		data.Run(run)
		ex.Run(run)
	}
}
