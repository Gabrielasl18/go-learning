// concurrent1
// Make the tests pass!

// I AM NOT DONE
package concurrent

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

func TestPrinter(t *testing.T) {
	/* Cria um novo buffer de bytes "[]byte" com alguns métodos para torná-lo utilizável com outros componentes
	vc pode escrever nele, ler a partir dele. é um arquivo de memória
	*/
	var buf bytes.Buffer
	print(&buf)

	//string retorna o conteúdo da parte não lida do buffer como uma string. Se o Buffer for um ponteiro nulo, ele retornará "<nil>".
	out := buf.String()

	for i := 0; i < 3; i++ {
		want := fmt.Sprintf("Hello from goroutine %d!", i)
		//verificar se a string que out se encontra em want
		if !bytes.Contains([]byte(out), []byte(want)) {
			t.Errorf("Output did not contain expected string. Wanted: %q, Got: %q", want, out)
		}
	}
}

func print(buf *bytes.Buffer) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	goroutines := 3

	for i := 0; i < goroutines; i++ {
		//adiciona uma nova goroutine ao contador do waitgroup toda vez que o loop roda
		wg.Add(1)

		//func de goroutine que passa o contado do loop (i)
		go func(i int) {
			//vai ser o ultimo comando executado pela func para dizer que finalizou a goroutine, mesmo que falhe, o contador do waitgroup vai decrementar uma goroutine
			defer wg.Done()

			//trava o Mutex para evitar acessos concorrentes ao buffer
			mu.Lock()
			fmt.Fprintf(buf, "Hello from goroutine %d!\n", i)

			//libera o Mutex para evitar acessos concorrentes ao buffer
			mu.Unlock()
		}(i)
	}

	//trava o fluxo da goroutine atual (funcao print) até que o contador do waitgroup zere
	wg.Wait()
}
