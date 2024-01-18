package exercises

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p Pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func ExerciseFour() {
	pessoa1 := Pessoa{
		nome:      "gabriela",
		sobrenome: "lacerda",
		idade:     21,
	}
	fmt.Println(pessoa1.nomeCompleto())
}
