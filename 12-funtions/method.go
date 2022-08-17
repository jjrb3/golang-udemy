package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person person
	lpm    bool
}

func main() {
	sc1 := secretAgent{
		person: person{
			firstName: "Jeremy",
			lastName: "Reyes",
		},
		lpm:    true,
	}

	sc2 := secretAgent{
		person: person{
			firstName: "Helen",
			lastName: "Mejia",
		},
		lpm:    false,
	}

	sc1.show()
	sc2.show()
}

// func ("Receptor or method") "function name"() ("returns")
func (sa secretAgent) show() {
	fmt.Println("Hello, I'm", sa.person.firstName, sa.person.lastName)
}

