package slice

import "fmt"

func SliceTest() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice2 := append(slice, 6)
	fmt.Println(slice2)
	for indice, valor := range slice {
		fmt.Println(indice, valor)
	}
	for indice2, _ := range slice {
		fmt.Print(indice2)
	}
	for _, valor := range slice {
		fmt.Print(valor)
	}
	sabores := []string{"chocolate", "abacaxi", "uva", "maça", "queijo", "mussarela"}
	fatia := sabores[:3]
	fmt.Println(fatia)

	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)

}
