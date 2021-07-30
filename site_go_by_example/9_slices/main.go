package main

import "fmt"

func main() {
	s1 := make([]int, 3)
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s1, s1, len(s1), cap(s1))
	s1 = append(s1, 8)
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s1, s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	s2[0] = 9
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s2, s2, len(s2), cap(s2))
	s2 = append(s2, 8)
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s2, s2, len(s2), cap(s2))

	s3 := make([]int, len(s2))
	copy(s3, s2)
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s3, s3, len(s3), cap(s3))

	s4 := make([]int, 1)
	copy(s4, s2)
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s4, s4, len(s4), cap(s4))

	s5 := []string{"a", "b", "c", "d"}
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s5, s5, len(s5), cap(s5))
	fmt.Printf("Slice: %v\n", s5[1:3])

	s6 := make([][]int, 3)
	for i, _ := range s6 {
		internalSlice := make([]int, i+1)
		s6[i] = internalSlice
	}

	for _, internalSlice := range s6 {
		fmt.Println(internalSlice)
	}

	s7 := s5[0:2]
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s7, s7, len(s7), cap(s7))
	s5[1] = "x"
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s7, s7, len(s7), cap(s7))

	arr1 := [3]float64{1.1, 2.2, 3.3}
	s8 := arr1[:2]
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s8, s8, len(s8), cap(s8))
	s8[0] = 1.9
	fmt.Println(s8, " - ", arr1)
	s8 = s8[:cap(s8)]
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s8, s8, len(s8), cap(s8))
	s8 = append(s8, 4.4)
	fmt.Printf("Slice: %v, type %T, len %d, cap %d\n", s8, s8, len(s8), cap(s8))
	s8[0] = 1.8
	fmt.Println(arr1)

	var s9 []string
	s9 = append(s9, s5[1:3]...)
	fmt.Println(s9)
}
