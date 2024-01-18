package exercises

import "fmt"

func ExerciseTwo() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, value := range s {
		fmt.Println(i, value)
	}
	fmt.Printf("%T", s)
}
