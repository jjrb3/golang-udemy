// Package example
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r := strings.NewReader("Jeremy Reyes B.")

	io.Copy(f, r)

	log.Println("log Println") // Log but not stop.
	log.Fatalln("log Fatalln") // Log and stop with exit estatus.
	panic("panic")             // Stop completely and call the file called.
}
