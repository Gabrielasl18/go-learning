package codewars

import (
	"testing"
)

// melhoria: usar strings.ContainsRune(vowels, c) e colocar os vogais dentro da variável vowels (vowels := AEIOUaeiou) e strings.Builder para concatenar
func TestSortVowels(t *testing.T) {
	expected := "L|\n|O\nr|\n|E\nm|\n |\n|I\np|\ns|\n|U\nm|\n |\nd|\n|O\nl|\n|O\nr|\n |\ns|\n|I\nt|\n |\n|A\nm|\n|E\nt|"
	r := sortVowels("LOrEm IpsUm dOlOr sIt AmEt")

	if r != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, r)
	}

}

func sortVowels(s string) string {
	var newString string

	for i, value := range s {
		if string(value) == "a" || string(value) == "A" || string(value) == "e" || string(value) == "E" ||
			string(value) == "i" || string(value) == "I" || string(value) == "O" || string(value) == "o" || string(value) == "u" || string(value) == "U" {
			newString += "|" + string(value)
		} else {
			newString += string(value) + "|"
		}
		if i != len(s)-1 {
			newString += "\n"
		}
	}
	return newString
}
