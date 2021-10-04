package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Raw args:")
	for i, arg := range os.Args {
		fmt.Printf("  arg #%d: `%s`\n", i, arg)
	}

	fmt.Println("Flags:")
	stringFlag := flag.String("str", "default", "a string")
	numberFlag := flag.Int("int", 42, "a number")
	boolFlag := flag.Bool("bool", false, "a boolean")
	var some string
	flag.StringVar(&some, "strvar", "something default", "some string")

	flag.Parse()

	fmt.Println("str:", *stringFlag)
	fmt.Println("int:", *numberFlag)
	fmt.Println("bool:", *boolFlag)
	fmt.Println("some:", some)
	fmt.Println("tail:", flag.Args())
}
