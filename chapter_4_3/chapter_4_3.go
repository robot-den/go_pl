package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	var m map[string]int
	m2 := map[string]int{}
	m3 := make(map[string]int)

	fmt.Println(m, m == nil, len(m), m["try"])
	fmt.Println(m2, m2 == nil, len(m2), m2["try"])
	fmt.Println(m3, m3 == nil, len(m3), m3["try"])

	v, ok := m["try"]
	fmt.Println(v, ok)
	// dedup()
	charCount()
}

func charCount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charCount: %v", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("Counts: total different runes count is %d\n", len(counts))
	for rune, count := range counts {
		fmt.Printf("%s - %d\n", string(rune), count)
	}

	fmt.Printf("utflen: total different sizes is %d\n", len(utflen))
	for len, count := range utflen {
		fmt.Printf("%d - %d\n", len, count)
	}

	if invalid > 0 {
		fmt.Printf("Invalid runes count is %d", invalid)
	}
}

func dedup() {
	seen := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v", err)
		os.Exit(1)
	}
}
