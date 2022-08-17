package main

import (
	"fmt"
	"runtime"
)

var x uint8 // Positive values
var y int	// Positive and negative values

func main() {
	x = 255
	y = -14

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)	// OS
	fmt.Println(runtime.GOARCH) // Architecture
}
