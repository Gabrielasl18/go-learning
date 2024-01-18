package exercises

import "fmt"

func ExerciseThree() {
	defer fmt.Println("primeira coisa que vai ser apresentado por ultimo por conta do defer")
	fmt.Println(10)
	fmt.Println(5)
	fmt.Println(0)
	fmt.Println("ultima coisa que vai ser apresentado primeiro")
}
