package exercises

import "fmt"

func ExerciseEight() {
	testing := test()(2, 3)
	fmt.Println(testing)
}

func test() func(int, int) int {
	return func(x int, y int) int {
		return x * y
	}
}
