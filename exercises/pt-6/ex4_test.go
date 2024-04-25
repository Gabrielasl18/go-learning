package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseFour(t *testing.T) {
	var buf bytes.Buffer
	expected := "gabriela lacerda\n"
	exerciseFour(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p Pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func exerciseFour(out io.Writer) {
	pessoa1 := Pessoa{
		nome:      "gabriela",
		sobrenome: "lacerda",
		idade:     21,
	}
	fmt.Fprintln(out, pessoa1.nomeCompleto())
}
