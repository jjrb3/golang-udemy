package main

import "fmt"

type parents struct {
	fistName string
	lastName string
}

type family struct {
	fistName string
	lastName string
	parents  map[int]parents
}

func main() {
	f1 := family{
		fistName: "Jeremy",
		lastName: "Reyes",
		parents: map[int]parents{
			0: {
				fistName: "Jacqueline",
				lastName: "Barrios",
			},
			1: {
				fistName: "Freddy",
				lastName: "Reyes",
			},
		},
	}

	f2 := family{
		fistName: "Helen",
		lastName: "Mejia",
		parents: map[int]parents{
			0: {
				fistName: "Andrea",
				lastName: "Reyes",
			},
			1: {
				fistName: "Javier",
				lastName: "Mejia",
			},
		},
	}

	fmt.Print(f1, f2.parents[0])

	// Anonymous struct.
	as := struct {
		name string
		age  int
	}{
		name: "Aleja",
		age:  27,
	}
	fmt.Println(as)
}
