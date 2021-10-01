package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	originDir, err := os.Getwd()
	check(err)

	path := "site_go_by_example/50_directories/www"
	os.RemoveAll(path)

	err = os.Mkdir(path, 0755)
	check(err)
	defer func() {
		os.Chdir(originDir)
		os.RemoveAll(path)
	}()

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	createEmptyFile(filepath.Join(path, "file1.txt"))

	err = os.MkdirAll(filepath.Join(path, "my", "directory"), 0755)
	check(err)

	createEmptyFile(filepath.Join(path, "my", "file2.txt"))
	createEmptyFile(filepath.Join(path, "my", "file3.txt"))
	createEmptyFile(filepath.Join(path, "my", "directory", "file4.txt"))

	fmt.Println("Visiting subdirs")
	err = filepath.Walk(path, visit)
	check(err)

	err = os.Chdir(filepath.Join(path, "my"))
	check(err)
	listFiles(".")

	err = os.Chdir("../..")
	check(err)
	listFiles(".")
}

func visit(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		listFiles(path)
	}
	return nil
}

func listFiles(path string) {
	c, err := os.ReadDir(path)
	check(err)
	fmt.Println("List of children in", path, ":")
	fmt.Printf("  |%-10s|%-5s|\n", "Filename", "Dir?")
	fmt.Println("  |__________|_____|")
	for _, entry := range c {
		fmt.Printf("  |%-10s|%-5v|\n", entry.Name(), entry.IsDir())
	}
}
