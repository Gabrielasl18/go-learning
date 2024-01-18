package exercises

import "fmt"

func ExerciseTwo() {
	slice := []int{1, 9, 2, 3, 1, 1, 1, 1, 1}
	fmt.Println(variadico(slice...))
}

func variadico(i ...int) int {
	total := 0
	for _, value := range i {
		total += value
	}
	return total
}
