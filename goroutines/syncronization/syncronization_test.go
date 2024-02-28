package gourotines

import (
	"fmt"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	done := make(chan bool, 1)
	go check(done)
	if <-done {
		go check2()
		time.Sleep(time.Second)
	} else {
		t.Error("It not work")
	}
}

func check(done chan bool) {
	fmt.Println("Welcome to...")
	time.Sleep(time.Second)
	fmt.Println("TutorialsPoint")
	done <- true
}

func check2() {
	fmt.Println("Learn Go from here!")
}
