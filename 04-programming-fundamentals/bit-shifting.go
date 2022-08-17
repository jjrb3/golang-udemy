package main

import "fmt"

const(
	_ = iota
	kb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {
	a := 4
	b := a << 1
	fmt.Printf("%d\t\t%b\n", a, a)
	fmt.Printf("%d\t\t%b\n", b, b)

	fmt.Printf("%d\t\t%b\n",kb, kb)
	fmt.Printf("%d\t\t%b\n",gb, gb)
	fmt.Printf("%d\t\t%b\n",tb, tb)
}
