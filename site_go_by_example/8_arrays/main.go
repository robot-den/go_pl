package main

import "fmt"

func main() {
	var a1 [3]int
	fmt.Println(a1)
	a1[1] = 8
	fmt.Println(len(a1), a1, a1[2])

	a2 := [5]int{1, 2, 4}
	fmt.Println(a2)

	var a3 [3][3]int
	a3[1][1] = 8
	for _, arr := range a3 {
		fmt.Println(arr)
	}
}
