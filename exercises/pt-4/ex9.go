package exercises

import "fmt"

func ExerciseNine() {
	m := map[string][]string{
		"gabriela_lacerda": []string{
			"skate", "ser legal",
		},
		"luciana_silva": []string{
			"tocar viola", "ser legal",
		},
		"joana_pinheiro": []string{
			"surfar", "ser chata",
		},
	}
	m["leonardo_papel"] = []string{
		"correr", "nadar", "patinar",
	}

	delete(m, "luciana_silva")

	for key, value := range m {
		fmt.Println(key)
		for i, h := range value {
			fmt.Println("\t", i, "-", h)
		}
	}
}
