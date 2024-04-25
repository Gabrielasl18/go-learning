package functionsgo

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestCallback(t *testing.T) {
	var buf bytes.Buffer
	expected := "330\n"

	mainFunctions(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func mainFunctions(out io.Writer) {
	t := onlyPairs(sum, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Fprintln(out, t)
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func onlyPairs(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
