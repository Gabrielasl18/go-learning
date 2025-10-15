package codewars

import (
	"fmt"
)

// input -> codewars.Disemvowel("This website is for losers LOL!")
func Disemvowel(comment string) string {
	format := ""
	stringFormated := ""
	for i := 0; i < len(comment); i++ {
		if string(comment[i]) != "a" && string(comment[i]) != "A" && string(comment[i]) != "e" && string(comment[i]) != "E" && string(comment[i]) != "i" &&
			string(comment[i]) != "I" && string(comment[i]) != "O" && string(comment[i]) != "o" && string(comment[i]) != "u" && string(comment[i]) != "U" {
			format = string(comment[i])
			stringFormated = stringFormated + fmt.Sprint(format)
		}
	}
	return stringFormated
}
