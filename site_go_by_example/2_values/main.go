package main

import "fmt"

func main() {
	a := "1" + "2"
	fmt.Printf("%s is an %T, go %#v\n", a, a, a)
	b := 1 + 2
	fmt.Printf("%d is an %T, go %#v\n", b, b, b)
	c := 1.0 / 2.0
	fmt.Printf("%0.2f is an %T, go %#v\n", c, c, c)
	d := true
	fmt.Printf("%t is an %T, go %#v\n", d, d, d)
	e := []int{1, 2}
	fmt.Printf("%v is an %T, go %#v, zero element address: %p\n", e, e, e, e)
	f := &[]string{"1", "2"}
	fmt.Printf("%p (%x) is an %T, go %#v\n", f, f, f, f)
	g := struct{}{}
	fmt.Printf("%v is an %T, go %#v\n", g, g, g)
}
