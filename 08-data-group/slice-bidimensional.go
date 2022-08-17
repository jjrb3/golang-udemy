package main

import "fmt"

func main() {
	jr := []string{"Jeremy", "Jose", "Reyes", "Barrios"}
	fmt.Println(jr)
	ar := []string{"Andrea", "Carolina", "Reyes", "Barrios"}
	fmt.Println(ar)

	f := [][]string{jr, ar}

	fmt.Println(f)
}
