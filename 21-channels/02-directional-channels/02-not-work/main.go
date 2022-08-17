package main

import "fmt"

func main() {
	// Buffered channel.
	ca := make(chan <- int, 2) //  <- Significa solo enviar información. (send-only)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Println("---")
	fmt.Printf("%T\n", ca)
}
