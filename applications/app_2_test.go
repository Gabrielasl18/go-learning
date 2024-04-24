package applications

import (
	"encoding/json"
	"testing"
)

func TestDocJsonUnmarshal(t *testing.T) {
	r1, r2, r3 := docJsonUnmarshal()

	var result1, result2, result3 Informacoes
	err1 := json.Unmarshal(r1, &result1)
	check(err1)

	err2 := json.Unmarshal(r2, &result2)
	check(err2)

	err3 := json.Unmarshal(r3, &result3)
	check(err3)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Informacoes struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Estado        string  `json:"Estado"`
	Contabancaria float64 `json:"Contabancaria"`
}

func docJsonUnmarshal() ([]byte, []byte, []byte) {
	t := []byte(`{"Nome":"Gabríéla","Sobrenome":"Lacerda","Idade":21,"Estado":"Rio de Janeiro","Contabancaria":10.098}`)
	t2 := []byte(`{"Nome":"Julie","Sobrenome":"Texeira","Idade":65,"Estado":"Bahia","Contabancaria":2.098}`)
	t3 := []byte(`{"Nome":"Tarsício","Sobrenome":"Almeida","Idade":90,"Estado":"Rio Grande do Sul","Contabancaria":898}`)
	return t, t2, t3
}
