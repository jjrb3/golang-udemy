package main

import "fmt"

func main() {
	var x [5]uint8
	x[3] = 2

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
}
