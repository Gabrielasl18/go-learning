package variable

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestVariables(t *testing.T) {
	var buf bytes.Buffer
	expected := ` x: 10, float64
y: olá bom dia, string
	
x: 20, float64
x: 30, int` + "\n"

	variables(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func variables(out io.Writer) {
	x := 10.0
	y := "olá bom dia"
	fmt.Fprintf(out, "x: %v, %T\n", x, x)
	fmt.Fprintf(out, "y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Fprintf(out, "x: %v, %T\n", x, x)
	fmt.Fprintf(out, "x: %v, %T\n", z, z)
}
