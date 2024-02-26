package gourotines

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func RoutinesGO() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)

	go func1()
	go func2()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
	}
	wg.Done()
}
