package loop

import "fmt"

func Looping() {
	for i := 0; i < 800; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d x %d = %d\n", j, i, j*i)
		}
	}
}
