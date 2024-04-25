package slice

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestSlice(t *testing.T) {
	var buf bytes.Buffer
	expected := `[1 2 3 4 5]
[1 2 3 4 5 6]
0 1
1 2
2 3
3 4
4 5
0123412345[chocolate abacaxi uva]
[chocolate abacaxi maça queijo mussarela]` + "\n"
	sliceTest(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

func sliceTest(out io.Writer) {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Fprintln(out, slice)
	slice2 := append(slice, 6)
	fmt.Fprintln(out, slice2)
	for indice, valor := range slice {
		fmt.Fprintln(out, indice, valor)
	}
	for indice2, _ := range slice {
		fmt.Fprint(out, indice2)
	}
	for _, valor := range slice {
		fmt.Fprint(out, valor)
	}
	sabores := []string{"chocolate", "abacaxi", "uva", "maça", "queijo", "mussarela"}
	fatia := sabores[:3]
	fmt.Fprintln(out, fatia)

	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Fprintln(out, sabores)

}
