package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseSeven(t *testing.T) {
	var buf bytes.Buffer
	expected := "2\nola\n"
	exerciseSeven(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseSeven(out io.Writer) {
	d := func(x int, y int) int {
		return x / y
	}
	c := func() {
		fmt.Fprintln(out, "ola")
	}
	fmt.Fprintln(out, d(10, 5))
	c()
}
