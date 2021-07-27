package main

import "fmt"

var global = "global"

func main() {
	fmt.Println(global)
	var a string
	fmt.Printf("String a: %s\n", a)
	a = "filled with some text"
	fmt.Printf("String a: %s\n", a)
	var b, c = 5, 6
	fmt.Printf("Ints b and c: %d, %d\n", b, c)
	d := 3.333
	fmt.Printf("Float d: %g\n", d)
}
