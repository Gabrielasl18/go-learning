package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseTwo(t *testing.T) {
	var buf bytes.Buffer
	expected := `0 1
1 2
2 3
3 4
4 5
5 6
6 7
7 8
8 9
9 10
[]int`

	exerciseTwo(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseTwo(out io.Writer) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, value := range s {
		fmt.Fprintln(out, i, value)
	}
	fmt.Fprintf(out, "%T", s)
}
