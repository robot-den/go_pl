package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	sl := arr[:3]
	sl2 := sl[:4]

	fmt.Println(sl, len(sl), cap(sl))
	fmt.Println(sl2, len(sl2), cap(sl2))
	fmt.Println("-------------")

	sl3 := arr[:4]
	sl4 := arr[:5]
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("sl3: %v\n", sl3)
	fmt.Printf("sl4: %v\n", sl4)
	reverse(sl3)
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("sl3: %v\n", sl3)
	fmt.Printf("sl4: %v\n", sl4)
	fmt.Println("-------------")

	sl5 := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(sl5)
	reverse(sl5[:3])
	fmt.Println(sl5)
	reverse(sl5[3:])
	fmt.Println(sl5)
	reverse(sl5)
	fmt.Println(sl5)
	fmt.Println("-------------")

	arr2 := [5]string{"one", "", "three", "", "five"}
	fmt.Println(nonempty(arr2[:]))
	arr3 := [5]string{"one", "", "three", "", "five"}
	fmt.Println(nonempty2(arr3[:]))
}

func nonempty2(s []string) []string {
	out := s[:0]

	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}

	return out
}

func nonempty(s []string) []string {
	i := 0

	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}

	return s[:i]
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
