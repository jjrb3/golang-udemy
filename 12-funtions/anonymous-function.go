package main

import "fmt"

func main() {
	example()

	func() {
		fmt.Println("Execute anonymous example")
	}()

	func(x int) {
		fmt.Println("The parameter is", x)
	}(42)

	// Function by expression.
	f := func() {
		fmt.Println("My first expresion")
	}

	f()

	fmt.Println(anotherFunction()())
}

func example() {
	fmt.Println("Execute example")
}

func anotherFunction() func() int {
	return func() int {
		return 1990
	}
}