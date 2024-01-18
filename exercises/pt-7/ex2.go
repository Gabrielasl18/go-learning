package exercises

import "fmt"

type Pessoa struct {
	nome string
}

func mudeMe(p *Pessoa) {
	p.nome = "lacerda"
}

func ExerciseTwo() {
	p := Pessoa{nome: "gabriela"}
	fmt.Println(p)
	mudeMe(&p)
	fmt.Println(p)
}
