package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseTen(t *testing.T) {
	var buf bytes.Buffer
	expected := "26\n15\n"
	exerciseTen(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func adder() func(...int) int {
	sum := 0
	return func(x ...int) int {
		for _, value := range x {
			sum += value
		}
		return sum
	}
}

func exerciseTen(out io.Writer) {
	test := adder()(8, 9, 9)
	test_2 := adder()(10, 2, 3)
	fmt.Fprintln(out, test)
	fmt.Fprintln(out, test_2)
}
