package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseThree(t *testing.T) {
	var buf bytes.Buffer
	expected := `10
5
0
ultima coisa que vai ser apresentado primeiro
primeira coisa que vai ser apresentado por ultimo por conta do defer` + "\n"

	exerciseThree(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseThree(out io.Writer) {
	defer fmt.Fprintln(out, "primeira coisa que vai ser apresentado por ultimo por conta do defer")
	fmt.Fprintln(out, 10)
	fmt.Fprintln(out, 5)
	fmt.Fprintln(out, 0)
	fmt.Fprintln(out, "ultima coisa que vai ser apresentado primeiro")
}
