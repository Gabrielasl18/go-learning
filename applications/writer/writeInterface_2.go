package write

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// WriteBuff codifica um objeto do tipo user em JSON e o escreve em um buffer de bytes.
func WriteBuff() {
	// Cria um novo buffer de bytes (um espaço de armazenamento temporário em memória)
	buf := new(bytes.Buffer)

	// Cria um objeto do tipo user
	u := user{
		Name: "bob",
		Age:  20,
	}

	// Usa o encoder JSON para codificar o objeto user e escrever no buffer de bytes
	// O método Encode do json.NewEncoder escreve no buffer
	err := json.NewEncoder(buf).Encode(u)
	if err != nil {
		panic(err)
	}

	// Imprime o conteúdo do buffer de bytes (representação JSON do objeto user)
	fmt.Print(buf)
}
