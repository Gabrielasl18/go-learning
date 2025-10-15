package codewars

// input -> codewars.SortMyString("CodeWars")
func SortMyString(s string) string {
	var strOdd string
	var strEven string
	for i, value := range s {
		if i%2 == 0 || i == 0 {
			strEven += string(value)
		} else {
			strOdd += string(value)
		}
	}
	strEven += " " + strOdd

	return strEven
}
