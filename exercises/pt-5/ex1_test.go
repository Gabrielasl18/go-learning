package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseOne(t *testing.T) {
	var buf bytes.Buffer
	expected := `baunilha - creme - chocolate - 
abacaxi - acerola - manga - 
maracuja - goiaba - morango - `

	exerciseOne(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type pessoa struct {
	nome                    string
	sobrenome               string
	saboresFavoritosSorvete []string
}

func exerciseOne(out io.Writer) {
	humano := pessoa{
		nome:                    "Gabriela",
		sobrenome:               "Lacerda",
		saboresFavoritosSorvete: []string{"baunilha", "creme", "chocolate"},
	}
	humano2 := pessoa{
		nome:                    "Poli",
		sobrenome:               "Ana",
		saboresFavoritosSorvete: []string{"abacaxi", "acerola", "manga"},
	}
	humano3 := pessoa{
		nome:                    "Jessica",
		sobrenome:               "Ferreira",
		saboresFavoritosSorvete: []string{"maracuja", "goiaba", "morango"},
	}
	for _, value := range humano.saboresFavoritosSorvete {
		fmt.Fprint(out, value+" - ")
	}
	fmt.Fprint(out, "\n")
	for _, value := range humano2.saboresFavoritosSorvete {
		fmt.Fprint(out, value+" - ")
	}
	fmt.Fprint(out, "\n")
	for _, value := range humano3.saboresFavoritosSorvete {
		fmt.Fprint(out, value+" - ")
	}
}
