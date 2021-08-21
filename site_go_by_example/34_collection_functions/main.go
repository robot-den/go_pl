package main

import (
	"log"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func main() {
	strs := []string{"apple", "orange", "plum", "watermalon"}

	log.Println(Index(strs, "orange"))
	log.Println(Index(strs, "some"))

	log.Println(Include(strs, "plum"))

	log.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "water")
	}))
}
