package exercises

import "fmt"

func ExerciseSeven() {
	d := func(x int, y int) int {
		return x / y
	}
	c := func() {
		fmt.Println("ola")
	}
	fmt.Println(d(10, 5))
	c()
}
