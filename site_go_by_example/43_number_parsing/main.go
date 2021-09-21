package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Printf("%T: %.3f\n", f, f)

	i, _ := strconv.ParseInt("234", 0, 64)
	fmt.Printf("%T: %d\n", i, i)

	i2, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Printf("%T: %d\n", i2, i2)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Printf("%T: %d\n", u, u)

	k, _ := strconv.Atoi("135")
	fmt.Printf("%T: %d\n", k, k)

	_, e := strconv.Atoi("non-number")
	fmt.Println(e)
}
