package flowcontrol

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestSwitchcASE(t *testing.T) {
	var buf bytes.Buffer
	expected := `um nome qualquer é joana
sempre que um nome qualquer for joana, a giovana também aparece` + "\n"

	switchCase(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func switchCase(out io.Writer) {
	umnomequalquer := "joana"

	switch umnomequalquer {
	case "gabi":
		fmt.Fprintln(out, "um nome qualquer é gabi")
	case "joao":
		fmt.Fprintln(out, "um nome qualquer é joao")
	case "joana":
		fmt.Fprintln(out, "um nome qualquer é joana")
		fallthrough
	case "giovana":
		fmt.Fprintln(out, "sempre que um nome qualquer for joana, a giovana também aparece")
	default:
		fmt.Fprintln(out, "não tem nenhum nome qualquer")
	}
}
