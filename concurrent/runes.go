package concurrent

import (
	"fmt"
	"runtime"
	"sync"
)

func RunesTest() {
	count := 0
	goroutines := 10

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			v := count
			//yield
			runtime.Gosched()
			v++
			count = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
