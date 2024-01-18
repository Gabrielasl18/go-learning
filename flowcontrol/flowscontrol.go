package flowcontrol

import "fmt"

func FlowControl() {

	for X := 33; X <= 122; X++ {
		fmt.Printf("%d, %#x, %#U\n", X, X, X)
	}

}
