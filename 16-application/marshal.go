package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
	Age uint
}

func main() {
	p1 := Person{
		FirstName: "Jeremy",
		LastName: "Reyes",
		Age: 32,
	}

	p2 := Person{
		FirstName: "Aleja",
		LastName: "Barragan",
		Age: 27,
	}

	people := []Person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
