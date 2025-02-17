package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseOne(t *testing.T) {
	var buf bytes.Buffer

	exerciseOne(&buf)

	expected := "42, James Bond, true\n" +
		"42, int\n" +
		"true, bool\n" +
		"James Bond, string\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseOne(out io.Writer) {
	a := 42
	b := "James Bond"
	c := true

	fmt.Fprintf(out, "%d, %s, %t\n", a, b, c)
	fmt.Fprintf(out, "%d, %T\n", a, a)
	fmt.Fprintf(out, "%t, %T\n", c, c)
	fmt.Fprintf(out, "%s, %T\n", b, b)
}
