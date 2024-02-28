package gourotines

import (
	"fmt"
	"testing"
)

func TestChannels(t *testing.T) {
	message := make(chan string)

	go InsertMessages(message)

	msg := <-message
	fmt.Println(msg)
	if msg == "" {
		t.Error("Message not found")
	}
}

func InsertMessages(message chan string) {
	message <- "opaaa"
}
