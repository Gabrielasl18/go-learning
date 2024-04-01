package codewars

import (
	"math"
)

func FindNextSquare(n int64) int64 {
	i := n + 1
	boolean := false

	if math.Floor(math.Sqrt(float64(n))) == math.Sqrt(float64(n)) {
		for !boolean {
			if math.Sqrt(float64(i)) == math.Floor(math.Sqrt(float64(i))) {
				boolean = true
				return int64(i)
			}
			i++
		}
	}
	return -1
}
