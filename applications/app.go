package applications

import (
	"encoding/json"
	"fmt"
)

/*
JSON (javascript object notation)
*/

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Estado        string
	Contabancaria float64
}

func DocJson() {
	pessoaTest := Pessoa{
		Nome:          "Gabriela",
		Sobrenome:     "Lacerda",
		Idade:         21,
		Estado:        "Rio de Janeiro",
		Contabancaria: 10.098,
	}
	pessoaTest_2 := Pessoa{
		Nome:          "Julie",
		Sobrenome:     "Texeira",
		Idade:         65,
		Estado:        "Bahia",
		Contabancaria: 2.098,
	}
	pessoaTest_3 := Pessoa{
		Nome:          "Tarsício",
		Sobrenome:     "Almeida",
		Idade:         90,
		Estado:        "Rio Grande do Sul",
		Contabancaria: 898,
	}
	pt, err := json.Marshal(pessoaTest)
	fmt.Println(string(pt))
	if err != nil {
		fmt.Println(err)
	}
	pt2, err := json.Marshal(pessoaTest_2)
	fmt.Println(string(pt2))
	if err != nil {
		fmt.Println(err)
	}
	pt3, err := json.Marshal(pessoaTest_3)
	fmt.Println(string(pt3))
	if err != nil {
		fmt.Println(err)
	}
}
