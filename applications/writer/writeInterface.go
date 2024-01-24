package write

import (
	"fmt"
	"os"
)

func WriteFile() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo", err)
		return
	}
	defer file.Close()
	message := []byte("Olá, esse é um arquivo escrito por: Gabriela Lacerda")

	bytesWritten, err := file.Write(message)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo", err)
		return
	}
	fmt.Printf("Numero de bytes escritos: %d\n", bytesWritten)
}
