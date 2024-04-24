package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseFour(t *testing.T) {
	var buf bytes.Buffer

	exerciseFour(&buf)

	expected := "77, 1001101, 0x4d\n" +
		"154, 10011010, 0x9a\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseFour(out io.Writer) {
	a := 77
	fmt.Fprintf(out, "%d, %b, %#x\n", a, a, a)
	g := a << 1
	fmt.Fprintf(out, "%d, %b, %#x\n", g, g, g)
}
