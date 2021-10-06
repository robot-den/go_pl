package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "is set")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println("\nFirst 10 environment variables:")

	var pairs [][]string
	var maxKeyLen int
	for _, e := range os.Environ()[:10] {
		pair := strings.SplitN(e, "=", 2)
		pairs = append(pairs, pair)
		if maxKeyLen < len(pair[0]) {
			maxKeyLen = len(pair[0])
		}
	}

	for _, pair := range pairs {
		fmt.Printf("%*s%s\n", -maxKeyLen-2, pair[0], pair[1])
	}

}
