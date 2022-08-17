package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type form interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s form) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
}
