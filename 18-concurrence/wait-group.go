package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	if runtime.NumCPU() > 2 {
		wg.Add(2) 		 // Quantity Goroutines.
	} else {
		wg.Add(runtime.NumCPU()) // Quantity Goroutines.
	}

	runtime.GOMAXPROCS(4) // Select procesor

	go foo()
	go bar()

	wg.Wait() // Wait that goroutines finish.
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}

	wg.Done() // Notify that 1 goroutine finished.
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done() // Notify that 1 goroutine finished.
}
