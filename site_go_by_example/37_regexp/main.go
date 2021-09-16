package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	fmt.Println("str", r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	fmt.Println("str with sub:", r.FindStringSubmatch("peach punch"))
	fmt.Println("idx with sub:", r.FindStringSubmatchIndex("peach punch"))
	fmt.Println("all str:", r.FindAllString("peach punch pinch", -1))
	fmt.Println("all ids with sub:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println("byte:", r.Match([]byte("peach punch pinch")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
