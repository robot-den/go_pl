package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("First cycle:", i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println("Second cycle:", j)
	}

	for {
		fmt.Println("Third cycle: endless")
		break
	}

	for j := 0; j < 3; j++ {
		if j%2 != 0 {
			continue
		}
		fmt.Println("Fourth cycle:", j)
	}

	for {
		fmt.Println("Useless cycle")
		return
	}
}
