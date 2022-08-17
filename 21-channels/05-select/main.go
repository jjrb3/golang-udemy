package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	// Send.
	go send(par, impar, salir)

	// Receive.
	receive(par, impar, salir)

	fmt.Println("Finished")
}

func send(p, i, s chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}

	s <- 0
}

func receive(p, i, s <-chan int) {
	for {
		select {
		case v:= <-p:
			fmt.Println("From Par channel", v)
		case v:= <-i:
			fmt.Println("From Impar channel", v)
		case v:= <-s:
			fmt.Println("From Salir channel", v)
			return
		}
	}
}
