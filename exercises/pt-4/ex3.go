package exercises

import "fmt"

func ExerciseThree() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(s[:3])
	fmt.Println(s[4:])
	fmt.Println(s[1:7])
	fmt.Println(s[2:9])
	fmt.Println(s[2 : len(s)-1])
	fmt.Printf("%T", s)
}
