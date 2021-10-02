package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "gofile")
	check(err)
	defer os.Remove(f.Name())

	fmt.Println("temp file:", f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp("", "godir")
	check(err)
	defer os.RemoveAll(dname)

	fmt.Println("temp dir:", dname)

	fname := filepath.Join(dname, "somefile")
	err = os.WriteFile(fname, []byte{1, 2, 3, 4}, 0666)

	check(err)
}
