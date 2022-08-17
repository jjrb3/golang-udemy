package main

import "fmt"

type money float32

var m money

func main() {
	m = 1_000

	fmt.Print(m)
	fmt.Printf("%T", m)
}
