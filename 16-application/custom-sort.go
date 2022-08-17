package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func main() {
	p1 := persona{"Eduar", 32}
	p2 := persona{"María", 25}
	p3 := persona{"Carolina", 56}
	p4 := persona{"Adriana", 60}

	personas := []persona{p1, p2, p3, p4}

	fmt.Println(personas)
}
