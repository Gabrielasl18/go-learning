package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestInterface_go_2(t *testing.T) {
	var buf bytes.Buffer
	expected := `{4 true}
au-au
playful, cute, funny
{2 false}
piu-piu
small, yellow` + "\n"

	interface_go_2(&buf)

	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected:\n%s\nActual:\n%s\n", expected, buf.String())
	}
}

type Animal interface {
	MakeAnimalSpeak() string
	particulars() string
}

/*
A interface animal fornece uma abstração sobre os detalhes específicos dos tipos concretos (dog e bird).
Isso permite que você trate diferentes tipos de animais de maneira uniforme.
*/
type dog struct {
	paws   int
	esteem bool
}

type bird struct {
	wings  int
	esteem bool
}

/*
Uma instância de dog (d) está implementando o método MakeAnimalSpeak() da interface animal.
Portanto, automaticamente, ela é considerada uma instância de Animal
*/
func (d dog) MakeAnimalSpeak() string {
	return "au-au"
}
func (d dog) particulars() string {
	return "playful, cute, funny"
}

func (b bird) MakeAnimalSpeak() string {
	return "piu-piu"
}
func (b bird) particulars() string {
	return "small, yellow"
}

func selected(a Animal, out io.Writer) {
	fmt.Fprintln(out, a)
	fmt.Fprintln(out, a.MakeAnimalSpeak())
	fmt.Fprintln(out, a.particulars())
}

func interface_go_2(out io.Writer) {
	d := dog{paws: 4, esteem: true}
	b := bird{wings: 2, esteem: false}

	selected(d, out)
	selected(b, out)
	// funçoes sendo invocadas e vao retornar mensagens diferentes, pois os tipos sao diferentes, apesar de terem a mesma instância
}
