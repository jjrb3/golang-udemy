package main

import "fmt"

type numberTwo int

var n2 numberTwo
var y2 int

func main() {
	fmt.Println(n2)
	fmt.Printf("El tipo de X es: %T\n", n2)

	n2 = 42
	fmt.Println(n2)

	y2 = int(n2)
	fmt.Println(y2)
	fmt.Printf("%T", y2)
}
