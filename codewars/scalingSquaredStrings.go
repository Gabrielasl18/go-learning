package codewars

import (
	"strings"
)

func Scale(str string, k int, v int) string {
	var newStr3 string
	intermediate := ""
	words := strings.Split(str, "\n")
	for _, word := range words {
		str2 := strings.Split(word, "")
		for _, letter := range str2 {
			for j := 0; j < k; j++ {
				newStr3 += letter
			}
		}
		for i := 0; i < v; i++ {
			intermediate += newStr3 + "\n"
		}
		newStr3 = ""
	}

	return strings.TrimRight(intermediate, "\n")
}
