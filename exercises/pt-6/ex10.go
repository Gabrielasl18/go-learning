package exercises

import "fmt"

func adder() func(...int) int {
	sum := 0
	return func(x ...int) int {
		for _, value := range x {
			sum += value
		}
		return sum
	}
}

func ExerciseTen() {
	test := adder()(8, 9, 9)
	test_2 := adder()(10, 2, 3)
	fmt.Println(test)
	fmt.Println(test_2)
}
