package exercises

import "fmt"

func ExerciseOne() {
	fmt.Println(returnInt())
	fmt.Println(returnIntString())
}

func returnInt() int {
	return 9
}

func returnIntString() (int, string) {
	return 10, "oi"
}
