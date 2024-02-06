package arrays

import "fmt"

var x [5]int
var y [6]int

func ArraysData() {
	x[0] = 1
	x[4] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(len(x))
}
