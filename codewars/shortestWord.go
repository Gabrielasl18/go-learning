package codewars

import (
	"strings"
)

// input -> FindShort("Meu nome é Gabriela Lacerda")
func FindShort(s string) int {
	t := strings.Split(s, " ")
	lengthLess := len(t[0])
	for _, value := range t {
		if len(value) < lengthLess {
			lengthLess = len(value)
		}
	}
	return lengthLess
}
