package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseFour(t *testing.T) {
	var buf bytes.Buffer
	expected := `[42 43 44 45 46 47 48 49 50 51]
[42 43 44 45 46 47 48 49 50 51 52]
[42 43 44 45 46 47 48 49 50 51 52 53 54 55]
[42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60]` + "\n"

	exerciseFour(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseFour(out io.Writer) {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{56, 57, 58, 59, 60}
	fmt.Fprintln(out, x)
	x = append(x, 52)
	fmt.Fprintln(out, x)
	x = append(x, 53, 54, 55)
	fmt.Fprintln(out, x)
	x = append(x, y...)
	fmt.Fprintln(out, x)
}
