package exercises

import "fmt"

func ExerciseOne() {
	s := [5]int{1, 2, 3, 4, 5}
	for i, value := range s {
		fmt.Println(i, value)
	}
	fmt.Printf("%T", s)
}
