package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseSix(t *testing.T) {
	var buf bytes.Buffer

	exerciseSix(&buf)

	expected := "2024 2025 2026 2027"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

const (
	_ = iota + 2023
	c0
	c1
	c2
	c3
)

func exerciseSix(out io.Writer) {
	fmt.Fprint(out, c0, c1, c2, c3)
}
