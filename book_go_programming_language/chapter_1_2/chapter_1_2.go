// Prints command arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

var sep = " "

func main() {
	fmt.Println(os.Args[0])
	args := os.Args[1:]
	// 1
	var s string
	for i := 0; i < len(args); i++ {
		s += args[i] + sep
	}
	fmt.Println(s)
	// 2
	s = ""
	for _, arg := range args {
		s += arg + sep
	}
	fmt.Println(s)
	// 3
	s = strings.Join(args, sep)
	fmt.Println(s)
	// 4
	fmt.Println(args[1:])
	// 5
	for i, arg := range args {
		fmt.Println(i, arg)
	}
}
