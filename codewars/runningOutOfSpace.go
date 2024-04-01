package codewars

import (
	"strings"
)

func Spacey(arr []string) []string {
	var stringFinal []string
	for i := 1; i <= len(arr); i++ {
		f := strings.Join(arr[:i], "")
		stringFinal = append(stringFinal, f)
	}
	return stringFinal
}
