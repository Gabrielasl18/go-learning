package datastructure

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	r := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7)
	if r != 7 {
		t.Error("Message not found")
	}
}

func binarySearch(array []int, number int) int {
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
			return number
		}
	}

	return 0
}
