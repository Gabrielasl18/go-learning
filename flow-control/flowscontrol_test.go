package flowcontrol

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestFlowControl(t *testing.T) {
	var buf bytes.Buffer
	expected := `33, 0x21, U+0021 '!'
34, 0x22, U+0022 '"'
35, 0x23, U+0023 '#'
36, 0x24, U+0024 '$'` + "\n"

	flowControl(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func flowControl(out io.Writer) {
	for X := 33; X <= 36; X++ {
		fmt.Fprintf(out, "%d, %#x, %#U\n", X, X, X)
	}

}
