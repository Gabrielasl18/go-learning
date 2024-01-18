package codewars

import (
	"fmt"
	"strconv"
	"strings"
)

var menor int
var maior int

func HighAndLow(in string) string {
	res1 := strings.Split(in, " ")
	maior, _ = strconv.Atoi(string(res1[0]))
	menor, _ = strconv.Atoi(string(res1[0]))
	for i := 0; i < len(res1); i++ {
		number, _ := strconv.Atoi(string(res1[i]))

		fmt.Println(number)
		if number > maior {
			maior = number
		} else if number < menor {
			menor = number
		}
	}
	final := fmt.Sprint(strconv.Itoa(maior), " ", strconv.Itoa(menor))
	return final
}
