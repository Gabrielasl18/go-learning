package write

import (
	"fmt"
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	WriteFile()
	data, err := os.ReadFile("example.txt")
	check(err)
	expected := "Olá, esse é um arquivo escrito por: Gabriela Lacerda"
	if string(data) != expected {
		t.Errorf("Erro: esperado '%s', mas obtido '%s'", expected, string(data))
	} else {
		t.Logf("Sucesso: conteúdo do arquivo: '%s'", string(data))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFile() {
	file, err := os.Create("example.txt")
	check(err)
	defer file.Close()
	message := []byte("Olá, esse é um arquivo escrito por: Gabriela Lacerda")

	bytesWritten, err := file.Write(message)
	check(err)
	fmt.Printf("Numero de bytes escritos: %d\n", bytesWritten)
}
