package datastructure

import "fmt"

func BinarySearch(array []int, number int) {
	first := 0
	last := len(array)

	for first < last {
		mid := (first + last) / 2
		if number < array[mid] {
			last = mid - 1
		} else if number > array[mid] {
			first = mid + 1
		} else {
			fmt.Printf("achei no indice: %d", mid)
			break
		}
	}
}
