package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseFive(t *testing.T) {
	var buf bytes.Buffer

	exerciseFive(&buf)

	expected :=
		"0 exercises.gab\n" +
			"int\n"

	if expected != buf.String() {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}

}

type gab int

var h int
var p gab

func exerciseFive(out io.Writer) {

	fmt.Fprintf(out, "%v %T\n", p, p)
	p = 42
	h = int(p)
	fmt.Fprintf(out, "%T\n", h)
}
