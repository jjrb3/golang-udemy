package main

import (
	"encoding/json"
	"fmt"
)

type p struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	s := `[{"FirstName":"Jeremy","LastName":"Reyes","Age":32},{"FirstName":"Aleja","LastName":"Barragan","Age":27}]`
	bs := []byte(s) // Convert byte slice

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []p

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Println("\nPerson number ", i+1)
		fmt.Println(v.FirstName, v.LastName, v.Age)
	}
}
