package main

import "fmt"

var m int

type money float32

var b money

func main() {
	m = 1_000

	fmt.Print(m)
	fmt.Printf("%T", m)

	b = 1_000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
