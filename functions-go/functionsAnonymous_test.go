package functionsgo

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestAnonymous(t *testing.T) {
	var buf bytes.Buffer
	expected := "3\n"

	anonymous(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func anonymous(out io.Writer) {

	soma := func(x int, y int) int {
		return x + y
	}(1, 2)
	fmt.Fprintln(out, soma)

}
