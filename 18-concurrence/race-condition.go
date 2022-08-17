package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()

			v := counter
			v++

			runtime.Gosched() // Yield

			counter = v

			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter: ", counter)
}
