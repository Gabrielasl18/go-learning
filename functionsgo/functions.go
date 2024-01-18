package functionsgo

import "fmt"

//func (receiver) name(parameters) (returns){code}

// funcoes sao definidas com PARAMETROS e sao chamadas com ARGUMENTOS

// Tudo em GO é *pass by value*

// É possível ter funcao de múltiplos retornos

// Quando tiver um parametro variadico, o variadico tem q ser o ultimo na hora de retornar (s string, x...int)

func MainSoma() {
	// total, quantos, oi := soma(11, 4, 3, 4, 5, 6)
	// fmt.Println(total, quantos, oi)
	deferFunction()
}

func soma(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "bom dia"
}

func deferFunction() {
	defer fmt.Println("091")
	fmt.Println("72")
	fmt.Println("53")
	fmt.Println("94")
}
