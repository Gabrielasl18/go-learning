package exercises

import "fmt"

func ExerciseFive() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:])
	y := []int{}
	y = append(y, x[:3]...)
	fmt.Println(y)
	y = append(y, x[6:]...)
	fmt.Println(y)
}
