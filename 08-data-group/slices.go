package main

import "fmt"

func main() {
	// Composite literal.
	x := []int{2, 3, 4, 5, 6}
	y := []int{21, 31, 41, 51, 61}

	fmt.Println(cap(x))
	fmt.Println(x[0])

	for i, v := range x {
		fmt.Printf("i: %v; v: %v\n", i, v)
	}

	fmt.Println(x[1:4])

	// Add to slice.
	x = append(x, 9, 8, 7)

	for i, v := range x {
		fmt.Printf("i: %v; v: %v\n", i, v)
	}

	fmt.Println("----------")

	x = append(x, y...)

	for i, v := range x {
		fmt.Printf("i: %v; v: %v\n", i, v)
	}

	fmt.Println("----------")

	// Remove.
	x = append(x[:6], x[8:]...)

	for i, v := range x {
		fmt.Printf("i: %v; v: %v\n", i, v)
	}

	// Slice with make.
	y = make([]int, 10, 10)

	y[9] = 200

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
}
