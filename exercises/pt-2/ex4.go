package exercises

import "fmt"

func ExerciseFour() {
	a := 77
	fmt.Printf("%d, %b, %#x\n", a, a, a)
	g := a << 1
	fmt.Printf("%d, %b, %#x\n", g, g, g)
}
