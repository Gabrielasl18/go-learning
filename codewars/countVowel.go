package codewars

func GetCount(str string) (count int) {
	strFormated := []rune(str)
	arrayVowel := [10]string{"a", "A", "e", "E", "i", "I", "o", "O", "u", "U"}
	for i := 0; i < len(strFormated); i++ {
		for j := 0; j < len(arrayVowel); j++ {
			if string(strFormated[i]) == arrayVowel[j] {
				count++
			}
		}
	}
	return count
}
