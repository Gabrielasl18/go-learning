package codewars

func EvenOrOdd(number int) string {
	numberString := ""
	if number%2 == 0 {
		numberString = "Even"
	} else {
		numberString = "Odd"
	}
	return numberString
}
