package goroutine

import (
	"testing"
	"time"
)

func TestTimeouts2(t *testing.T) {
	select {
	//isso imprimirá “timeout” depois que o tempo passou.
	case <-time.After(1 * time.Second):
		t.Errorf("timeout")
	}
}
