package applications

import (
	"encoding/json"
	"fmt"
)

type Informacoes struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"` // essa tag de json é para explicar o que vai vir no rótulo do outro json, mas que no meu programa pode mudar
	Estado        string  `json:"Estado"`
	Contabancaria float64 `json:"Contabancaria"`
}

func DocJsonUnmarshal() {
	t := []byte(`{"Nome":"Gabriela","Sobrenome":"Lacerda","Idade":21,"Estado":"Rio de Janeiro","Contabancaria":10.098}`)
	t2 := []byte(`{"Nome":"Julie","Sobrenome":"Texeira","Idade":65,"Estado":"Bahia","Contabancaria":2.098}`)
	t3 := []byte(`{"Nome":"Tarsício","Sobrenome":"Almeida","Idade":90,"Estado":"Rio Grande do Sul","Contabancaria":898}`)

	var julieInfo Informacoes
	var gabeInfo Informacoes
	var tarcInfo Informacoes

	err := json.Unmarshal(t, &gabeInfo)
	if err != nil {
		fmt.Println(err)
	}
	err2 := json.Unmarshal(t2, &julieInfo)
	if err2 != nil {
		fmt.Println(err2)
	}
	err3 := json.Unmarshal(t3, &tarcInfo)
	if err != nil {
		fmt.Println(err3)
	}
	fmt.Println(julieInfo)
	fmt.Println(gabeInfo)
	fmt.Println(tarcInfo)
}
