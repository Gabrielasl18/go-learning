package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseTwo(t *testing.T) {
	var buf bytes.Buffer
	expected := "20\n"

	exerciseTwo(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseTwo(out io.Writer) {
	slice := []int{1, 9, 2, 3, 1, 1, 1, 1, 1}
	fmt.Fprintln(out, variadico(slice...))
}

func variadico(i ...int) int {
	total := 0
	for _, value := range i {
		total += value
	}
	return total
}
