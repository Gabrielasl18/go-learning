package codewars

import (
	"fmt"
	"strings"
)

// input -> codewars.PartList([]string{"I", "wish", "I", "hadn't", "come"})
func PartList(arr []string) string {
	var txtFormat string
	for i := 1; i < len(arr); i++ {
		txt := fmt.Sprintf("(%s,%s)", strings.Join(arr[:i], " "), strings.Join(arr[i:], " "))
		txtFormat += txt
	}
	return txtFormat
}
