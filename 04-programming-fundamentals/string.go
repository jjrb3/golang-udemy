package main

import "fmt"

func main() {
	s1 := "Hello World"
	s2 := `This is a break
	line.`

	fmt.Println(s1)
	fmt.Printf("The s1 type is: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("The s2 type is: %T\n", s2)

	fmt.Println("------")

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T", bs)

	fmt.Println("------")

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}

	fmt.Println("\n------")

	for i, v := range s1 {
		fmt.Printf("Index: %d - Value  %#x\n", i, v)
	}

	fmt.Printf("With the verb %q i said print the string %s", "%s", s1)
}
