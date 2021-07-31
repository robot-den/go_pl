package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m2 := map[string]int{"b": 3, "c": 4}

	fmt.Println("m1:", m1, "m2:", m2)

	m1["a"] = 1
	m2["d"] = 5

	fmt.Println("m1:", m1, "m2:", m2)
	delete(m2, "b")
	fmt.Println("m2:", m2, len(m2))
	value, ok := m2["c"]
	fmt.Println(value, ok)

	for key, value := range m2 {
		fmt.Println(key, ":", value)
	}
}
