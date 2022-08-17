package main

import "fmt"

type Number int

var N Number

func main() {
	fmt.Println(N)
	fmt.Printf("El tipo de X es: %T\n", N)

	N = 42
	fmt.Println(N)
}
