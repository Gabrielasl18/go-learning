package codewars

import (
	"fmt"
)

// input -> codewars.SeriesOfNth((5))
func SeriesOfNth(n int) string {
	x := 1.0
	t := 0.0
	for i := 1; i <= n; i++ {
		if i > 1 {
			t += float64(1 / (x + 3))
			x = x + 3
		} else {
			t += 1
		}
	}
	strFinal := fmt.Sprintf("%.2f", float64(t))
	return strFinal
}
