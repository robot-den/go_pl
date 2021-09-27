package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := "site_go_by_example/46_reading_files/file.txt"
	// read entire file into memory
	data, err := os.ReadFile(path)
	check(err)
	fmt.Println(string((data)))

	// open file to read as we want
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	// read first n bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: '%s'\n", n1, string(b1))

	// read at some point
	o2, err := f.Seek(5, 0)
	check(err)
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes from offset %d: '%s'\n", n2, o2, string(b2))

	// read using io package
	o3, err := f.Seek(10, 0)
	check(err)
	b3 := make([]byte, 4)
	n3, err := io.ReadAtLeast(f, b3, 4)
	check(err)
	fmt.Printf("%d bytes from offset %d: '%s'\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	// read using bufio
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(4)
	check(err)
	fmt.Printf("4 bytes: '%s'\n", string(b4))
}
