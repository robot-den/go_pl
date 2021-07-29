package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	switch x {
	case 10, 20:
		fmt.Printf("Ten or Twenty: %d\n", x)
	case 33, 44, 55:
		fmt.Printf("33, 44 or 55: %d\n", x)
	default:
		fmt.Println("Nothing of expected")
	}

	switch {
	case x > 50:
		fmt.Println("Wow")
	default:
		fmt.Println("Meh")
	}

	whoami(3)
	whoami("str")
	whoami(7.5)
	whoami(struct{}{})
}

func whoami(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("I am a string")
	case float64, float32:
		fmt.Println("I am a float")
	case int:
		fmt.Println("I am an int")
	default:
		fmt.Println("I don't know who I am")
	}
}
