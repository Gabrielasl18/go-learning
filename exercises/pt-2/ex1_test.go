package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseOne(t *testing.T) {
	var buf bytes.Buffer

	exerciseOne(&buf)

	expected := "100\n" + "1100100\n" + "0x64"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseOne(out io.Writer) {
	x := 100
	fmt.Fprintf(out, "%d\n", x)
	fmt.Fprintf(out, "%b\n", x)
	fmt.Fprintf(out, "%#x", x)
}
