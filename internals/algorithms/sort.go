package algorithms

func bubble_sort(input []int) []int {
	for i := range len(input) - 1 {
		for j := range len(input[i:]) - 1 {
			if input[j] > input[j+1] {
				temp := input[j]
				input[j] = input[j+1]
				input[j+1] = temp
			}
		}
	}
	return input
}
