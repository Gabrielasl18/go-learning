package strgs

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestString(t *testing.T) {
	var buf bytes.Buffer
	expected := `gabriela
[103 97 98 114 105 101 108 97]
[103 97 98 114 105 101 108 97]
a = U+0061
˜ = U+02DC` + "\n"
	stringRune(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

var e = rune('a')
var c = '˜'

func stringRune(out io.Writer) {
	s := "gabriela"

	s_rune := []rune(s)
	v_byte := []byte(s)

	fmt.Fprintln(out, string(s_rune))
	fmt.Fprintln(out, v_byte)
	fmt.Fprintln(out, s_rune)
	fmt.Fprintf(out, "%c = %U\n", e, e)
	fmt.Fprintf(out, "%c = %U", c, c)
}
