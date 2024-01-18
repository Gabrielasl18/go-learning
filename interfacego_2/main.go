package main

import "fmt"

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

func selected(a Animal) {
	fmt.Println(a)
	fmt.Println(a.MakeAnimalSpeak())
	fmt.Println(a.particulars())
}

func main() {
	d := dog{paws: 4, esteem: true}
	b := bird{wings: 2, esteem: false}

	selected(d)
	selected(b)
	// funçoes sendo invocadas e vao retornar mensagens diferentes, pois os tipos sao diferentes, apesar de terem a mesma instância
}
