package exercises

import "fmt"

func ExerciseSix() {

	estadosBrasil := make([]string, 10, 11)
	estadosBrasil = []string{"Acre", "Alagoas", "Bahia", "Ceará", "Piauí", "Rio de Janeiro", "Rondônia", "São Paulo", "Tocantins", "Distrito Federal", "Sergipe", "Rio Grande do Sul"}

	fmt.Println(estadosBrasil)

	for indice, value := range estadosBrasil {
		fmt.Println(indice, value)
	}

	fmt.Println(len(estadosBrasil), "--", cap(estadosBrasil))
}
