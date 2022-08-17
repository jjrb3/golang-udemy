package main

import "fmt"

func main() {
	x := []int{1, 2, 3,4}
	fmt.Println(x)

	y := append(x[:2], x[3:]...)
	fmt.Println(y)
}
