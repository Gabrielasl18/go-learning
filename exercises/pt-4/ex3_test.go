package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseThree(t *testing.T) {
	var buf bytes.Buffer
	expected := `[1 2 3]
[5 6 7 8 9 10]
[2 3 4 5 6 7]
[3 4 5 6 7 8 9]
[3 4 5 6 7 8 9]` + "\n" + "[]int"

	exerciseThree(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseThree(out io.Writer) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Fprintln(out, s[:3])
	fmt.Fprintln(out, s[4:])
	fmt.Fprintln(out, s[1:7])
	fmt.Fprintln(out, s[2:9])
	fmt.Fprintln(out, s[2:len(s)-1])
	fmt.Fprintf(out, "%T", s)
}
