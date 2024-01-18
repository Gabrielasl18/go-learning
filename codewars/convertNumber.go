package codewars

import (
	"fmt"
	"strconv"
	"strings"
)

func Digitize(n int) []int {
	numberArray := strings.Split(strconv.Itoa(n), "")
	number := []int{}
	numberStr := ""
	nn := []int{}
	nn = append(nn, n)
	numberFinal := []int{}
	for i := len(numberArray); i > 0; i-- {
		n, _ := strconv.Atoi(numberArray[i-1])
		number = append(number, n)
		numberStr += strconv.Itoa(n)
	}
	numberStr = strings.Replace(numberStr, "", ",", len(numberStr))
	numberStr = strings.TrimLeft(numberStr, ",")
	fmt.Println(numberStr)
	fmt.Println(nn)
	return numberFinal
}
