package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseTwo(t *testing.T) {
	var buf bytes.Buffer
	expected := `65
	U+0041 'A'
	U+0041 'A'
	U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'` + "\n"

	exerciseTwo(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseTwo(out io.Writer) {
	for i := 65; i <= 66; i++ {
		fmt.Fprintln(out, i)
		for j := 0; j <= 2; j++ {
			fmt.Fprintf(out, "\t%#U\n", i)
		}
	}
}
