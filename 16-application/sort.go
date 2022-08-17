package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 2, 6, 1}
	xs := []string{"Jeremy", "Aleja", "Andrea", "Helen", "Jacky"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
