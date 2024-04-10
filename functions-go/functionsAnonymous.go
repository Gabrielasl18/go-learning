package functionsgo

import "fmt"

func Anonymous() {

	soma := func(x int, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(soma)

}
