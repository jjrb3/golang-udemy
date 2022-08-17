package main

import "fmt"

func main() {
	defer foo() // Execute the last
	bar()
}

func foo()  {
	fmt.Println("This is food")
}

func bar()  {
	fmt.Println("This is bad")
}
