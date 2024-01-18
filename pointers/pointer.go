package pointers

import "fmt"

func PointerTest() {
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)
	fmt.Println("*pointer =", *pointer)
	*pointer = "gabriela shark"
	fmt.Println("*pointer new =", *pointer)
	fmt.Println("creature =", creature)

}
