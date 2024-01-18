package exercises

import "fmt"

type pessoa_2 struct {
	nome                    string
	sobrenome               string
	saboresFavoritosSorvete []string
}

func ExerciseTwo() {

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
		fmt.Println(value.nome)
		for _, value := range value.saboresFavoritosSorvete {
			fmt.Println("\t", value)
		}
	}
}
