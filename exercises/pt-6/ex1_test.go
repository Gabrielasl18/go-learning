package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseOne(t *testing.T) {
	var buf bytes.Buffer
	expected := "9\n"

	exerciseOne(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseOne(out io.Writer) {
	fmt.Fprintln(out, returnInt())
	fmt.Println(returnIntString())
}

func returnInt() int {
	return 9
}

func returnIntString() (int, string) {
	return 10, "oi"
}
