package main

import "fmt"

func main() {
	runPointersExample(1)
	runPointersExample("string")
	runPointersExample(make(map[string]int))
	runPointersExample(make([]int, 0, 3))
}

func runPointersExample(i interface{}) {
	switch t := i.(type) {
	case int:
		runIntExample(t)
	case string:
		runStringExample(t)
	case map[string]int:
		runMapExample(t)
	case []int:
		runSliceExample(t)
	default:
		fmt.Println("Type is not supported")
	}
}

func runIntExample(t int) {
	fmt.Println("Initial:", t)
	changeInt(t)
	fmt.Println("By value:", t)
	changeIntP(&t)
	fmt.Println("By pointer:", t)
}

func changeInt(t int) {
	t++
}

func changeIntP(t *int) {
	*t++
}

func runStringExample(t string) {
	fmt.Println("Initial:", t)
	changeString(t)
	fmt.Println("By value:", t)
	changeStringP(&t)
	fmt.Println("By pointer:", t)
}

func changeString(t string) {
	t = t + " by value"
}

func changeStringP(t *string) {
	*t += " by pointer"
}

func runMapExample(t map[string]int) {
	fmt.Println("Initial:", t)
	changeMap(t)
	fmt.Println("By value:", t)
	changeMapP(&t)
	fmt.Println("By pointer:", t)
}

func changeMap(t map[string]int) {
	t["byValue"] = 10
}

func changeMapP(t *map[string]int) {
	(*t)["byPointer"] = 20
}

func runSliceExample(t []int) {
	fmt.Println("Initial:", t)
	changeSlice(t)
	fmt.Println("By value:", t)
	changeSliceP(&t)
	fmt.Println("By pointer:", t)
}

func changeSlice(t []int) {
	t = append(t, 1)
}

func changeSliceP(t *[]int) {
	*t = append(*t, 2)
}
