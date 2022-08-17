package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v)
	fmt.Println(ok)

	v, ok = <-c
	fmt.Println(v)
	fmt.Println(ok)

	fmt.Println("Finalizando.")
}
