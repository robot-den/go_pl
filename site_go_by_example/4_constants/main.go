package main

import (
	"fmt"
	"math"
)

const firstConst = 10

func main() {
	fmt.Printf("const value %v of type %T\n", firstConst, firstConst)
	fmt.Println(math.Sin(firstConst)) // firstConst is untyped, so we can pass it as float64 without conversion

	const secondConst = 'x'
	fmt.Printf("const value %v of type %T\n", secondConst, secondConst)
	fmt.Println(math.Sin(secondConst)) // secondConst is untyped, so we can pass it as float64 without conversion
}
