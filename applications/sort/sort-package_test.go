package sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortPackage_1(t *testing.T) {
	carros := []carro{{"Chevette", 50, 8},
		{"Porsche", 500, 5},
		{"Fusca", 20, 30},
	}
	// Ordenando por potência
	sort.Sort(ordenaPorPotencia(carros))
	if !reflect.DeepEqual(carros, []carro{{"Fusca", 20, 30}, {"Chevette", 50, 8}, {"Porsche", 500, 5}}) {
		t.Error("Sorted Failed - Ordena Por Potencia []carro{{Fusca, 20, 30}, {Chevette, 50, 8}, {Porsche, 500, 5}} ")
	}
	// Resetando a slice antes de testar a ordenação por consumo
	carros = []carro{{"Chevette", 50, 8},
		{"Porsche", 500, 5},
		{"Fusca", 20, 30},
	}
	sort.Sort(ordenaPorConsumo(carros))
	if !reflect.DeepEqual(carros, []carro{{"Fusca", 20, 30}, {"Chevette", 50, 8}, {"Porsche", 500, 5}}) {
		t.Error("Sorted Failed - Ordena Por Consumo []carro{{Fusca, 20, 30}, {Chevette, 50, 8}, {Porsche, 500, 5}}")
	}
	// Resetando a slice antes de testar a ordenação por lucro da empresa
	carros = []carro{{"Chevette", 50, 8},
		{"Porsche", 500, 5},
		{"Fusca", 20, 30},
	}
	sort.Sort(ordenaPorLucroEmpresa(carros))
	if !reflect.DeepEqual(carros, []carro{{"Porsche", 500, 5}, {"Chevette", 50, 8}, {"Fusca", 20, 30}}) {
		t.Error("Sorted Failed - Ordena Por Lucro  []carro{{Porsche, 500, 5}, {Chevette, 50, 8}, {Fusca, 20, 30}}")
	}
}

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
