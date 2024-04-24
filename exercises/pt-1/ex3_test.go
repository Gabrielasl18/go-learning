package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// expected := 42      James Bond      true

func TestExerciseThree(t *testing.T) {
	var buf bytes.Buffer

	exerciseThree(&buf)

	expected := "42\t" + "James Bond\t" + "true"

	if expected != buf.String() {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

var r int = 42
var u string = "James Bond"
var i bool = true

func exerciseThree(out io.Writer) {
	s := fmt.Sprintf("%v\t%v\t%v", r, u, i)
	fmt.Fprintf(out, "%s", s)
}
