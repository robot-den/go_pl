package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// dump string into file
	path := "site_go_by_example/47_writing_files/file.txt"
	d1 := []byte("hello from golang\n!!!")
	err := os.WriteFile(path, d1, 0644)
	check(err)

	// create empty file
	path2 := "site_go_by_example/47_writing_files/file2.txt"
	f, err := os.Create(path2)
	check(err)
	defer f.Close()

	// write bytes to file
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// write string to file
	n3, err := f.WriteString("string\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	// write file to stable storage
	f.Sync()

	// write using buffered writer
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	// write data to stable storage
	w.Flush()
}
