package main

import "fmt"

var a int
var b = "Program"

func main() {
	fmt.Print(a, b)
	fmt.Println(a, b)
	fmt.Printf("\n%v - %v", a, b) // Impresión con el valor original
	fmt.Printf("\n%#v - %#v", a, b) // Impresion del valor
	fmt.Printf("\n%T - %T", a, b) // Imprimir tipo del valor
	fmt.Printf("\n%% - %%", a, b) // Imprimir tipo con valor de porcentaje

	fmt.Printf("\n\n%d - %s", a, b) // Impresión por formato

	s1 := fmt.Sprint("\n\nThe ", b, " say hello world") // Concatenación

	fmt.Print(s1)
	fmt.Printf("\n%T", s1)
}
