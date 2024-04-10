package flowcontrol

import (
	"fmt"
)

func SwitchCase() {
	umnomequalquer := "joana"

	switch umnomequalquer {
	case "gabi":
		fmt.Println("um nome qualquer é gabi")
	case "joao":
		fmt.Println("um nome qualquer é joao")
	case "joana":
		fmt.Println("um nome qualquer é joana")
		fallthrough
	case "giovana":
		fmt.Println("sempre que um nome qualquer for joana, a giovana também aparece")
	default:
		fmt.Println("não tem nenhum nome qualquer")
	}
}
