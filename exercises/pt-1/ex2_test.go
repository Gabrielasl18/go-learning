package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

var a int
var b string
var c bool

/*
	 expected := 0 +
				"	" +
				false
*/

func TestExerciseTwo(t *testing.T) {
	var buf bytes.Buffer
	exerciseTwo(&buf)

	expected := "0" + "\n" + "\nfalse\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseTwo(out io.Writer) {
	fmt.Fprintf(out, "%v\n%v\n%v\n", a, b, c)
}
