package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseTwo(t *testing.T) {
	var buf bytes.Buffer
	expected := `{gabriela}
{lacerda}` + "\n"

	exerciseTwo(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type Pessoa struct {
	nome string
}

func mudeMe(p *Pessoa) {
	p.nome = "lacerda"
}

func exerciseTwo(out io.Writer) {
	p := Pessoa{nome: "gabriela"}
	fmt.Fprintln(out, p)
	mudeMe(&p)
	fmt.Fprintln(out, p)
}
