package exercises

import "fmt"

type gabe int

var x gabe

func ExerciseFour() {

	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Println(x)
}
