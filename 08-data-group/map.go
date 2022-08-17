package main

import "fmt"

func main() {
	m := map[string]int{
		"Jeremy": 32,
		"Andrea": 36,
	}

	fmt.Println(m)

	fmt.Println(m["Jeremy"])
	fmt.Println(m["Sofi"])

	if v, ok := m["Sofi"]; ok {
		fmt.Println(v)
	}

	m["Jacqueline"] = 69

	fmt.Println("-----")

	for i, v := range m {
		fmt.Printf("%v - %v\n", i, v)
	}

	delete(m, "Jeremy")

	fmt.Println("-----")

	for i, v := range m {
		fmt.Printf("%v - %v\n", i, v)
	}
}
