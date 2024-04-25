package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type pessoa_2 struct {
	nome                    string
	sobrenome               string
	saboresFavoritosSorvete []string
}

func TestExerciseTwo(t *testing.T) {
	var buf bytes.Buffer
	expected := `Gabriela
	baunilha
	creme
	chocolate
Poli
	abacaxi
	acerola
	manga
Jessica
	maracuja
	goiaba
	morango` + "\n"

	exerciseTwo(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseTwo(out io.Writer) {

	pessoasMap := map[string]pessoa_2{
		"Lacerda": pessoa_2{
			nome:                    "Gabriela",
			sobrenome:               "Lacerda",
			saboresFavoritosSorvete: []string{"baunilha", "creme", "chocolate"},
		},
		"Ana": pessoa_2{
			nome:                    "Poli",
			sobrenome:               "Ana",
			saboresFavoritosSorvete: []string{"abacaxi", "acerola", "manga"},
		},
		"Ferreira": pessoa_2{
			nome:                    "Jessica",
			sobrenome:               "Ferreira",
			saboresFavoritosSorvete: []string{"maracuja", "goiaba", "morango"},
		},
	}

	for _, value := range pessoasMap {
		fmt.Fprintln(out, value.nome)
		for _, value := range value.saboresFavoritosSorvete {
			fmt.Fprintln(out, "\t", value)
		}
	}
}
