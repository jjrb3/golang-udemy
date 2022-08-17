package main

import "fmt"

func main() {
	c := make(chan int)

	// Send.
	go send(c)
	// Receive.
	receive(c)

	fmt.Println("Finished")
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
