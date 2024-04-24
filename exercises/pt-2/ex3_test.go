package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseThree(t *testing.T) {
	var buf bytes.Buffer

	exerciseThree(&buf)

	expected := "gabriela" + "90, int\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

const x = "gabriela"
const y int = 90

func exerciseThree(out io.Writer) {
	fmt.Fprint(out, x)
	fmt.Fprintf(out, "%d, %T\n", y, y)
}
