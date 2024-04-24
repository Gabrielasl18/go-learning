package applications

import (
	"encoding/json"
	"testing"
)

/*
JSON (javascript object notation)
*/

func TestDocJson(t *testing.T) {
	pessoaTest, pessoaTest_2, pessoaTest_3 := docJson()
	expected := `{"Nome":"Gabriela","Sobrenome":"Lacerda","Idade":21,"Estado":"Rio de Janeiro","Contabancaria":10.098}`
	expected_2 := `{"Nome":"Julie","Sobrenome":"Texeira","Idade":65,"Estado":"Bahia","Contabancaria":2.098}`
	expected_3 := `{"Nome":"Tarsício","Sobrenome":"Almeida","Idade":90,"Estado":"Rio Grande do Sul","Contabancaria":898}`

	pt, err := json.Marshal(pessoaTest)
	check(err)
	pt2, err := json.Marshal(pessoaTest_2)
	check(err)
	pt3, err := json.Marshal(pessoaTest_3)
	check(err)
	if expected != string(pt) || expected_2 != string(pt2) || expected_3 != string(pt3) {
		t.Log("error")
	}

}

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Estado        string
	Contabancaria float64
}

func docJson() (Pessoa, Pessoa, Pessoa) {
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
	return pessoaTest, pessoaTest_2, pessoaTest_3
}
