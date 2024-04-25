package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseNine(t *testing.T) {
	var buf bytes.Buffer
	expected := "20\n"
	exerciseNine(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseNine(out io.Writer) {
	o := test_2(afe)
	fmt.Fprintln(out, o)
}
func afe() int {
	return 20
}

func test_2(t func() int) int {
	return t()
}
