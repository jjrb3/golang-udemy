package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	// Doing login
	lp := `pass123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(lp))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully login")
}
