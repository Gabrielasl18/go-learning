package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseTwo(t *testing.T) {
	var buf bytes.Buffer

	exerciseTwo(&buf)

	expected := "false false false false true"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseTwo(out io.Writer) {
	x := 10 <= 8
	z := 7 != 7
	o := 9 >= 10
	p := 4 == 2
	t := 80 > 20
	fmt.Fprint(out, x, z, o, p, t)
}
