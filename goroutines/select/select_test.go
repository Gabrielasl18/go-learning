package gourotines

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go f1(c1)
	go f2(c2)

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	}
}

func f1(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "1"
}

func f2(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "2"
}
