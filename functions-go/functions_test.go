package functionsgo

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

//func (receiver) name(parameters) (returns){code}

// funcoes sao definidas com PARAMETROS e sao chamadas com ARGUMENTOS

// Tudo em GO é *pass by value*

// É possível ter funcao de múltiplos retornos

// Quando tiver um parametro variadico, o variadico tem q ser o ultimo na hora de retornar (s string, x...int)

func TestMainSoma(t *testing.T) {
	var buf bytes.Buffer
	expected := `33 6 bom dia
72
53
94` + "\n"

	mainSoma(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func mainSoma(out io.Writer) {
	total, quantos, oi := soma(11, 4, 3, 4, 5, 6)
	fmt.Fprintln(out, total, quantos, oi)
	deferFunction(out)
}

func soma(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "bom dia"
}

func deferFunction(out io.Writer) {
	defer fmt.Println(out, "091")
	fmt.Fprintln(out, "72")
	fmt.Fprintln(out, "53")
	fmt.Fprintln(out, "94")
}
