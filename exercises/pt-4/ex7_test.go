package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseSeven(t *testing.T) {
	var buf bytes.Buffer
	expected := `[Gabriela Lacerda Skate]
[Joana Pinheiro Surfar]
[Luciana Silva Tocar viola]` + "\n"

	exerciseSeven(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}

}

func exerciseSeven(out io.Writer) {
	slice := [][]string{
		[]string{
			"Gabriela",
			"Lacerda",
			"Skate",
		},
		[]string{
			"Joana",
			"Pinheiro",
			"Surfar",
		},
		[]string{
			"Luciana",
			"Silva",
			"Tocar viola",
		},
	}

	for _, value := range slice {
		fmt.Fprintln(out, value)
	}
}
