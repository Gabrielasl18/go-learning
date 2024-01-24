package sort

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenaPorPotencia []carro
type ordenaPorConsumo []carro
type ordenaPorLucroEmpresa []carro

func (x ordenaPorPotencia) Len() int {
	return len(x)
}
func (x ordenaPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia

}
func (x ordenaPorPotencia) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x ordenaPorConsumo) Len() int {
	return len(x)
}
func (x ordenaPorConsumo) Less(i, j int) bool {
	return x[i].consumo > x[j].consumo

}
func (x ordenaPorConsumo) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]

}

func (x ordenaPorLucroEmpresa) Len() int {
	return len(x)
}
func (x ordenaPorLucroEmpresa) Less(i, j int) bool {
	return x[i].consumo < x[j].consumo
}
func (x ordenaPorLucroEmpresa) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func SortPackage_2() {
	carros := []carro{carro{"Chevette", 50, 8},
		carro{"Porsche", 500, 5},
		carro{"Fusca", 20, 30},
	}
	fmt.Println("inicial\n", carros)
	sort.Sort(ordenaPorPotencia(carros))
	fmt.Println("potencia\n", carros)
	sort.Sort(ordenaPorConsumo(carros))
	fmt.Println("consumo\n", carros)
	sort.Sort(ordenaPorLucroEmpresa(carros))
	fmt.Println("lucro\n", carros)
}
