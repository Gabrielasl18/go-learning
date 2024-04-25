package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseOne(t *testing.T) {
	var buf bytes.Buffer
	expected := "0\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n"

	exerciseOne(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseOne(out io.Writer) {
	for i := 0; i <= 10; i++ {
		fmt.Fprintln(out, i)
	}
}
