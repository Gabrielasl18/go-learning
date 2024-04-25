package strct

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestStructGo(t *testing.T) {
	var buf bytes.Buffer
	expected := `{João da Silva false}
{Joana Maia true}` + "\n"
	structsGo(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func structsGo(out io.Writer) {
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}
	cliente2 := cliente{
		nome:      "Joana",
		sobrenome: "Maia",
		fumante:   true,
	}
	fmt.Fprintln(out, cliente1)
	fmt.Fprintln(out, cliente2)
}
