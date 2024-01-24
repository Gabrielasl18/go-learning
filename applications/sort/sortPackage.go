package sort

import (
	"fmt"
	"sort"
)

func SortPackage() {
	s := []string{"abobora", "banana", "macarrao", "laranja", "ervilha", "gato", "cachorro"}
	s2 := []int{5, 7, 8, 9, 0, 7, 6, 6, 4, 2, 1}

	// ordenando as strings
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)

	// ordenando os ints
	fmt.Println(s2)
	sort.Ints(s2)
	fmt.Println(s2)
}
