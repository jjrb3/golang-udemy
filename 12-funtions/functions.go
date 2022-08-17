package main

import "fmt"

func main() {
	xi := []int{1,2,3,45}

	fmt.Println(Foo(xi...))
}

func Foo(x ...int) int {
	s := 0

	for _,v := range x {
		s += v
	}

	return s
}
