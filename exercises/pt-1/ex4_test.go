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

	expected := "0 exercises.gabe\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type gabe int

var x gabe

func exerciseFour(out io.Writer) {
	fmt.Fprintf(out, "%v %T\n", x, x)
	x = 42
}
