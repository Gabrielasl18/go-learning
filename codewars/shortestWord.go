package codewars

import (
	"strings"
)

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
