package exercises

import "fmt"

func ExerciseSeven() {
	slice := [][]string{
		[]string{
			"Gabriela",
			"Lacerda",
			"Skate",
		},
		[]string{
			"Joana",
			"Pinheiro",
			"Surfar",
		},
		[]string{
			"Luciana",
			"Silva",
			"Tocar viola",
		},
	}

	for _, value := range slice {
		fmt.Println(value)
	}
}
