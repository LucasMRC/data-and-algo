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
	} else {
		if !slices.Contains([]string{"algo", "data", "all", "ex"}, os.Args[1]) {
			fmt.Printf("Invalid argument.\nValid options:\n\t- %s (algorithms)\n\t- %s (data structures)\n\t- %s (exercises)\n", "algo", "data", "ex")
			os.Exit(1)
		}
	}
	if argCount == 2 {
		what = os.Args[1]
		run = "all"
	} else {
		switch os.Args[1] {
		case "algo":
			if !slices.Contains([]string{"linear", "binary", "bubble", "all"}, os.Args[2]) {
				fmt.Printf("Invalid argument.\nValid options:\n\t- %s (Linear)\n\t- %s (Binary)\n\t- %s (Bubble search)\n", "linear", "binary", "bubble")
				os.Exit(1)
			}
		case "data":
			if !slices.Contains([]string{"linked", "double", "stack", "all"}, os.Args[2]) {
				fmt.Printf("Invalid argument.\nValid options:\n\t- %s (Linked List)\n\t- %s (Double Linked List)\n\t- %s (Stack)\n", "linked", "double", "stack")
				os.Exit(1)
			}
		case "ex":
			if !slices.Contains([]string{"crystall", "all"}, os.Args[2]) {
				fmt.Printf("Invalid argument.\nValid options:\n\t- %s (Crystall Balls)\n", "crystall")
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
	case "ex":
		ex.Run(run)
	}
}
