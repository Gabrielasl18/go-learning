package exercises

import "fmt"

func ExerciseSix() {
	x := 10
	test := func(number int) int {
		return 100 * number
	}(x)
	fmt.Println(test)

}
