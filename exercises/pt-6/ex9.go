package exercises

import "fmt"

func ExerciseNine() {
	o := testing(afe)
	fmt.Println(o)
}
func afe() int {
	return 20
}

func testing(t func() int) int {
	return t()
}
