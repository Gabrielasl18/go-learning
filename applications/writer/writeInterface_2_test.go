package write

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestWriteInterface(t *testing.T) {
	u := user{
		Name: "bob",
		Age:  20,
	}

	userBytesBuffer := WriteBuff(u)
	u_json := `{"name":"bob","age":20}`

	// Remove espaços em branco e caracteres de quebra de linha das strings
	userBytesBufferStr := strings.ReplaceAll(userBytesBuffer.String(), " ", "")
	userBytesBufferStr = strings.ReplaceAll(userBytesBufferStr, "\n", "")

	u_json = strings.ReplaceAll(u_json, " ", "")
	u_json = strings.ReplaceAll(u_json, "\n", "")

	if userBytesBufferStr != u_json {
		t.Error("failed write buffer", userBytesBufferStr, u_json)
	}
}

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// WriteBuff codifica um objeto do tipo user em JSON e o escreve em um buffer de bytes.
func WriteBuff(u user) *bytes.Buffer {
	// Cria um novo buffer de bytes (um espaço de armazenamento temporário em memória)
	buf := new(bytes.Buffer)

	// Usa o encoder JSON para codificar o objeto user e escrever no buffer de bytes
	// O método Encode do json.NewEncoder escreve no buffer
	err := json.NewEncoder(buf).Encode(u)
	if err != nil {
		panic(err)
	}
	return buf
}
