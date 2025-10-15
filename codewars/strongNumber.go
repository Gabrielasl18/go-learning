package codewars

import (
	"strconv"
	"strings"
)

// input -> codewars.Strong(145)
func Strong(n int) string {
	sliceOfNumbers := strings.Split(strconv.Itoa(n), "")
	var numberFinal int
	for i := 0; i < len(sliceOfNumbers); i++ {
		number, _ := strconv.Atoi(sliceOfNumbers[i])
		mult := 1
		for i := number; i > 0; i-- {
			mult *= i
		}
		numberFinal += mult
	}

	if numberFinal != n {
		return "Not Strong !!"
	}

	return "STRONG!!!!"
}
