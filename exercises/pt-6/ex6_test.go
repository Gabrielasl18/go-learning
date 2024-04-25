package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseSix(t *testing.T) {
	var buf bytes.Buffer
	expected := "1000\n"
	exerciseSix(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseSix(out io.Writer) {
	x := 10
	test := func(number int) int {
		return 100 * number
	}(x)
	fmt.Fprintln(out, test)

}
