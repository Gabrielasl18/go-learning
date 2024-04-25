package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseFive(t *testing.T) {
	var buf bytes.Buffer
	expected := `[42 43 44 45 46 47 48 49 50 51]
[42 43 44]
[42 43 44 48 49 50 51]` + "\n"

	exerciseFive(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseFive(out io.Writer) {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Fprintln(out, x[:])
	y := []int{}
	y = append(y, x[:3]...)
	fmt.Fprintln(out, y)
	y = append(y, x[6:]...)
	fmt.Fprintln(out, y)
}
