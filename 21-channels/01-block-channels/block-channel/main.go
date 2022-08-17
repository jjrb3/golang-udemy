package main

import "fmt"

func main() {
	// Unbuffered channel.
	ca := make(chan int)

	ca <- 42

	fmt.Println(<-ca)
}
