package arrays

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestInterface_go_2(t *testing.T) {
	var buf bytes.Buffer
	expected := "1 0\n5" + "\n"

	arraysData(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

var x [5]int
var y [6]int

func arraysData(out io.Writer) {
	x[0] = 1
	x[4] = 10
	fmt.Fprintln(out, x[0], x[1])
	fmt.Fprintln(out, len(x))
}
