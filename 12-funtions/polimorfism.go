package main

import "fmt"

type user struct {
	firstName string
	lastName  string
}

type agentSecret struct {
	person user
	lpm    bool
}

type human interface {
	show()
}

func main() {
	sc1 := agentSecret{
		person: user{
			firstName: "Jeremy",
			lastName: "Reyes",
		},
		lpm:    true,
	}

	sc2 := agentSecret{
		person: user{
			firstName: "Helen",
			lastName: "Mejia",
		},
		lpm:    false,
	}

	p := user{
		firstName: "Jaqueline",
		lastName: "Barrios",
	}

	sc1.show()
	sc2.show()
	sc1.person.show()

	barr(sc1)
	barr(sc2)
	barr(p)
}

func barr(h human)  {
	switch h.(type) { // Check type.
	case user:
		fmt.Println("Is a person ", h.(user).lastName)
		break
	case agentSecret:
		fmt.Println("Is a secrect agent ", h.(agentSecret).person.firstName)
		break
	}

	fmt.Println("Pass to function bar", h)
}

// func ("Receptor or method") "function name"() ("returns")
func (sa agentSecret) show() {
	fmt.Println("Hello, I'm", sa.person.firstName, sa.person.lastName)
}

func (u user) show() {
	fmt.Println("The person is ", u.lastName)
}

