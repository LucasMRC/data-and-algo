package exercises

import (
	"math"
)

func crystall_balls(bools []bool) int {
	rs := int(math.Floor(math.Sqrt(float64(len(bools)))))
	result := -1
	for i := rs - 1; i < len(bools); i += rs {
		if i == 0 && bools[i] {
			return result
		} else if bools[i] {
			step := i - rs + 1
			for step <= i {
				if bools[step] {
					result = step
					break
				}
				step++
			}
			break
		}
	}
	return result
}
