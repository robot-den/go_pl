package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("a", "b", "c.txt")
	fmt.Println(p)
	fmt.Println(filepath.Join("a//", "b/", "c"))
	fmt.Println(filepath.Join("a/./b", "c"))

	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.Base(p))
	fmt.Println(filepath.Split(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	ext := filepath.Ext(p)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(p, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
