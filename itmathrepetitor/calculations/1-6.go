package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}

func ex1() {
	fmt.Println("Ex 1:")
	fmt.Println("Silence is golden")
}

func ex2() {
	fmt.Println("\nEx 2:")
	now := time.Now()
	fmt.Println(now.Weekday())
	fmt.Println(now.Month())
	fmt.Println("Denis")
}

func ex3() {
	fmt.Println("\nEx 3:")
	//for i := 1; i <= 5; i++ {
	//	msg := ""
	//	for j := 0; j < i; j++ {
	//		msg += "0"
	//	}
	//	fmt.Println(msg)
	//}
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("%0*s", i, "0")
		fmt.Println(msg)
	}
}

func ex4() {
	rows := 5
	cols := 8
	letter := "A"
	fmt.Println("\nEx 4:")
	for i := 0; i < rows; i++ {
		msg := strings.Repeat(letter, cols)
		fmt.Println(msg)
	}
}

func ex5() {
	fmt.Println("\nEx 5:")
	fmt.Println("*     *     *")
	fmt.Println(" *   * *   *")
	fmt.Println("  * *   * *")
	fmt.Println("   *     *")
}

func ex6() {
	fmt.Println("\nEx 6:")
	fmt.Println(1 + 2 - 4)
}
