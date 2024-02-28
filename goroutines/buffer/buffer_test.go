package gourotines

import (
	"fmt"
	"testing"
)

func TestBuffering(t *testing.T) {
	//cria um canal com buffer de capacidade 2
	ch := make(chan string, 2)

	go bufferGoRoutines(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	if <-ch == "" {
		t.Error("message not found")
	}
}
func bufferGoRoutines(ch chan string) {
	ch <- ""
	ch <- ""
}
