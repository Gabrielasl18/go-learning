package exercises

import "fmt"

type gab int

var h int
var p gab

func ExerciseFive() {

	fmt.Printf("%v %T\n", p, p)
	p = 42
	fmt.Println(p)
	h = int(p)
	fmt.Println(h)
	fmt.Printf("%T\n", h)
}
