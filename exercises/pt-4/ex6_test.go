package exercises

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExerciseSix(t *testing.T) {
	var buf bytes.Buffer
	expected := `[Acre Alagoas Bahia Ceará Piauí Rio de Janeiro Rondônia São Paulo Tocantins Distrito Federal Sergipe Rio Grande do Sul]
0 Acre
1 Alagoas
2 Bahia
3 Ceará
4 Piauí
5 Rio de Janeiro
6 Rondônia
7 São Paulo
8 Tocantins
9 Distrito Federal
10 Sergipe
11 Rio Grande do Sul
12 -- 12` + "\n"

	exerciseSix(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func exerciseSix(out io.Writer) {

	estadosBrasil := make([]string, 10, 11)
	estadosBrasil = []string{"Acre", "Alagoas", "Bahia", "Ceará", "Piauí", "Rio de Janeiro", "Rondônia", "São Paulo", "Tocantins", "Distrito Federal", "Sergipe", "Rio Grande do Sul"}

	fmt.Fprintln(out, estadosBrasil)

	for indice, value := range estadosBrasil {
		fmt.Fprintln(out, indice, value)
	}

	fmt.Fprintln(out, len(estadosBrasil), "--", cap(estadosBrasil))
}
