package pointers

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestInterface_go_2(t *testing.T) {
	var buf bytes.Buffer
	expected := `creature = shark
pointer = 0xc00002c220
*pointer = shark
*pointer new = gabriela shark
creature = gabriela shark` + "\n"

	pointerTest(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func pointerTest(out io.Writer) {
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Fprintln(out, "creature =", creature)
	fmt.Fprintln(out, "pointer =", pointer)
	fmt.Fprintln(out, "*pointer =", *pointer)
	*pointer = "gabriela shark"
	fmt.Fprintln(out, "*pointer new =", *pointer)
	fmt.Fprintln(out, "creature =", creature)

}
