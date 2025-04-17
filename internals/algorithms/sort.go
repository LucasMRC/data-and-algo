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

func quick_sort(input []int) []int {
	qs(&input, 0, len(input) - 1)

	return input
}

func qs(i *[]int, lo, hi int) {
	if (lo >= hi) {
		return
	}
	pivot := partition(i, lo, hi)
	qs(i, lo, pivot - 1)
	qs(i, pivot + 1, hi)
}

func partition(i *[]int, lo, hi int) int {
	pivot := (*i)[hi]
	idx := lo - 1

	for j := lo; j < hi; j++ {
		if ((*i)[j] <= pivot) {
			idx++
			tmp := (*i)[j]
			(*i)[j] = (*i)[idx]
			(*i)[idx] = tmp
		}
	}

	idx++
	(*i)[hi] = (*i)[idx]
	(*i)[idx] = pivot

	return idx
}
