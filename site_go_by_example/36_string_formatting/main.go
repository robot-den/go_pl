package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	fmt.Printf("struct: %v\n", point{1, 2})
	fmt.Printf("struct: %+v\n", point{1, 2})
	fmt.Printf("struct: %#v\n", point{1, 2})
	fmt.Printf("struct: %T\n", point{1, 2})

	fmt.Printf("boolean: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("int in binary: %b\n", 123)
	fmt.Printf("int in hex: %x\n", 123)
	fmt.Printf("char from code: %c\n", 33)

	fmt.Printf("float: %f\n", 123.33)
	fmt.Printf("float scientific: %e\n", 123.33)
	fmt.Printf("float scientific: %E\n", 123.33)

	fmt.Printf("string: %s\n", "\"love\"")
	fmt.Printf("string as in source: %q\n", "\"love\"")
	fmt.Printf("string in hex bytes: %x\n", "\"love\"")

	fmt.Printf("pointer: %p\n", &point{1, 2})

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.3f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
