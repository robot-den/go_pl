package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if 1 == 1 {
		fmt.Println("WOW! 1 is equal 1!")
	}

	if rand.Intn(100) < 50 {
		fmt.Println("less than 50")
	} else {
		fmt.Println("greater than 50")
	}

	if x := rand.Intn(10); x < 5 {
		fmt.Println("less than 5:", x)
	} else if x == 8 {
		fmt.Println("equal to 8:", x)
	} else {
		fmt.Println("greater than 5, but not equal to 8:", x)
	}
}
