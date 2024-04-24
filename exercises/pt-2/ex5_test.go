package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

/*
expected := dasd
                                dsafmasdka
                sdasda98027kjds
        fdf


*/

func TestExerciseFive(t *testing.T) {
	var buf bytes.Buffer

	exerciseFive(&buf)

	expected := `dasd
				dsafmasdka	
		sdasda98027kjds
	fdf
	`

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseFive(out io.Writer) {
	x := `dasd
				dsafmasdka	
		sdasda98027kjds
	fdf
	`
	fmt.Fprint(out, x)
}
