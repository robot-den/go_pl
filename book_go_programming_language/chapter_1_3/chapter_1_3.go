// Depending of args parameter:
// lines_count - print words that present more than one time in files
// lines_dup - print words that present in several files
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("Provide work type: `lines_count` or `lines_dup`")
	}
	checkType := os.Args[1]
	paths := os.Args[2:]
	if checkType == "lines_count" {
		countLines(paths)
	} else if checkType == "lines_dup" {
		showDup(paths)
	} else {
		panic("Available check types are `lines_count` or `lines_dup`")
	}
}

func countLines(paths []string) {
	counts := make(map[string]int)

	if len(paths) == 0 {
		collectLinesCount(os.Stdin, counts)
	} else {
		readByLines(paths, counts)
		// readByGulps(paths, counts)
	}

	printLines(counts)
}

func showDup(paths []string) {
	counts := make(map[string]map[string]bool)

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed %v\n", err)
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			wordPaths := counts[scanner.Text()]
			if wordPaths != nil {
				wordPaths[path] = true
			} else {
				counts[scanner.Text()] = make(map[string]bool)
				counts[scanner.Text()][path] = true
			}
		}

		file.Close()
	}

	for word, paths := range counts {
		var keys []string
		for key, _ := range paths {
			keys = append(keys, key)
		}
		if len(keys) > 1 {
			fmt.Printf("Word `%s` is present in: %v\n", word, keys)
		}
	}
}

func readByGulps(paths []string, counts map[string]int) {
	for _, path := range paths {
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed %v\n", err)
			continue
		}

		text := string(bytes)
		lines := strings.Split(text, "\n")

		for _, line := range lines {
			if line == "" {
				continue
			}
			counts[line]++
		}
	}
}

func readByLines(paths []string, counts map[string]int) {
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed %v\n", err)
			continue
		}
		collectLinesCount(file, counts)
		file.Close()
	}
}

func collectLinesCount(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}

func printLines(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
