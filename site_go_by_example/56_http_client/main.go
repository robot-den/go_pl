package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response http status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	counter := NewCounter()

	for scanner.Scan() {
		fmt.Printf("line #%-3d => %s\n", counter(), firstN(scanner.Text(), 20))
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}

func NewCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
