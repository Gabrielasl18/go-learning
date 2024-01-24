package codewars

import (
	"strings"
)

func ToCamelCase(s string) string {
	oc := "_-"
	stringT := s
	finalString := ""
	for _, value := range oc {
		old := string(value)
		stringT = strings.Replace(stringT, old, " ", -1)
	}
	string_test := strings.Split(stringT, " ")
	for i := 0; i < len(string_test); i++ {
		if i == 0 {
			finalString += string_test[i]

		} else {
			finalString += strings.Title(string_test[i])

		}
	}
	return finalString
}
