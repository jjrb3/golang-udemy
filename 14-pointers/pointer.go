package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 43
	fmt.Println(*b)

	fmt.Println("-----------------")

	x := 0
	fmt.Println("X Before", x)
	fmt.Println("X Before", &x)

	foo(&x)
	fmt.Println("X After", x)
	fmt.Println("X After", &x)
}

func foo(y *int) {
	fmt.Println("Y Before", y)
	fmt.Println("Y Before", *y)
	*y = 42
	fmt.Println("Y After", y)
	fmt.Println("Y After", *y)
}
