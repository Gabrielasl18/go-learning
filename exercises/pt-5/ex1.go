package exercises

import "fmt"

type pessoa struct {
	nome                    string
	sobrenome               string
	saboresFavoritosSorvete []string
}

func ExerciseOne() {
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
		fmt.Print(value + " - ")
	}
	fmt.Print("\n")
	for _, value := range humano2.saboresFavoritosSorvete {
		fmt.Print(value + " - ")
	}
	fmt.Print("\n")
	for _, value := range humano3.saboresFavoritosSorvete {
		fmt.Print(value + " - ")
	}
}
