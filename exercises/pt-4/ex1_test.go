package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseOne(t *testing.T) {
	var buf bytes.Buffer
	expected := `0 1
1 2
2 3
3 4
4 5` + "\n" + "[5]int"

	exerciseOne(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseOne(out io.Writer) {
	s := [5]int{1, 2, 3, 4, 5}
	for i, value := range s {
		fmt.Fprintln(out, i, value)
	}
	fmt.Fprintf(out, "%T", s)
}
