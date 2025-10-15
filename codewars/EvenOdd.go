package codewars

// input -> codewars.EvenOrOdd(11)
func EvenOrOdd(number int) string {
	numberString := ""
	if number%2 == 0 {
		numberString = "Even"
	} else {
		numberString = "Odd"
	}
	return numberString
}
