package codewars

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func Digitize(n int) []int {
	numFinal := []int{}
	byteString, err := json.Marshal(n)
	if err != nil {
		fmt.Println(err)
	}
	for i := len(byteString) - 1; i >= 0; i-- {
		t, err := strconv.Atoi(string(byteString[i]))
		if err != nil {
			fmt.Println(err)
		}
		numFinal = append(numFinal, t)
	}
	return numFinal
}
