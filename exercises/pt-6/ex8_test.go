package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseEight(t *testing.T) {
	var buf bytes.Buffer
	expected := "6\n"
	exerciseEight(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseEight(out io.Writer) {
	testing := test()(2, 3)
	fmt.Fprintln(out, testing)
}

func test() func(int, int) int {
	return func(x int, y int) int {
		return x * y
	}
}
