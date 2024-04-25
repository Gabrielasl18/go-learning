package datastructure

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arrSorted := []int{7, 12, 23, 34, 67, 87, 90, 203}
	r := selectSort([]int{34, 67, 87, 23, 12, 7, 90, 203})
	if reflect.DeepEqual(r, arrSorted) {
		t.Log("ok!")
	} else {
		t.Error("error")
	}
}

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		reserva := arr[i]
		arr[i] = arr[index]
		arr[index] = reserva
	}
	return arr
}
