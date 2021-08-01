package main

import "fmt"

func main() {
	v1 := []int{1, 2, 3}
	v2 := [3]int{4, 5, 6}
	v3 := map[string]int{"a": 1, "b": 2, "c": 3}
	v4 := "a宿ق"

	fmt.Println("Array:")
	for index, value := range v1 {
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}

	fmt.Println("Slice:")
	for index, value := range v2 {
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}

	fmt.Println("Map:")
	for key, value := range v3 {
		fmt.Printf("Key: %s, value: %d\n", key, value)
	}

	fmt.Println("String:")
	for index, r := range v4 {
		fmt.Printf("Index: %d, rune: %d\n", index, r)
	}
}
