package main

import "fmt"

var z = 41

func main() {
	x := 2 + 4
	y := "James Bond"
	fmt.Println(x)
	fmt.Println(y)

	x = 50
	fmt.Println(x)
	fmt.Println(z)

	number()
}

func number()  {
	fmt.Println(z)
}
