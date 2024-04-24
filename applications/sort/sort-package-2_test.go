package sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortPackage(t *testing.T) {
	s_sorted := []int{0, 1, 2, 4, 5, 6, 6, 7, 7, 8, 9}
	s2_sorted := []string{"bbobora", "banana", "cachorro", "ervilha", "gato", "laranja", "macarrao"}

	s, s2 := sortPackage()
	if !reflect.DeepEqual(s2, s2_sorted) || !reflect.DeepEqual(s, s_sorted) {
		t.Error("Sorted Failed")
	}
}

func sortPackage() ([]int, []string) {
	s := []string{"abobora", "banana", "macarrao", "laranja", "ervilha", "gato", "cachorro"}
	s2 := []int{5, 7, 8, 9, 0, 7, 6, 6, 4, 2, 1}

	// ordenando as strings
	sort.Strings(s)

	// ordenando os ints
	sort.Ints(s2)
	return s2, s
}
