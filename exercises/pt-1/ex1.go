package exercises

import "fmt"

func ExerciseOne() {
	a := 42
	b := "James Bond"
	c := true

	fmt.Printf("%d, %s, %t\n", a, b, c)
	fmt.Printf("%d, %T\n", a, a)
	fmt.Printf("%t, %T\n", c, c)
	fmt.Printf("%s, %T\n", b, b)
}
