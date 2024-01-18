package exercises

import (
	"fmt"
)

var r int = 42
var u string = "James Bond"
var i bool = true

func ExerciseThree() {
	s := fmt.Sprintf("%v\t%v\t%v", r, u, i)
	fmt.Println(s)
}
