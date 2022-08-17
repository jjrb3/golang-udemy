package main

import "fmt"

var xx int = 42
var yy string = "James Bond"
var zz bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", xx, yy, zz)

	fmt.Println(s)
}
