package strgs

import "fmt"

var e = rune('a')
var c = '˜'

func StringRune() {
	s := "gabriela"

	s_rune := []rune(s)
	v_byte := []byte(s)

	fmt.Println(string(s_rune))
	fmt.Println(v_byte)
	fmt.Println(s_rune)
	fmt.Printf("%c = %U\n", e, e)
	fmt.Printf("%c = %U", c, c)
}
