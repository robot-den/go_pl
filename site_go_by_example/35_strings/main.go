package main

import (
	"log"
	s "strings"
)

func main() {
	str := "Hello, bro"
	log.Println(s.Contains(str, "ll"))
	log.Println(s.Count(str, "l"))
	log.Println(s.HasPrefix(str, "He"))
	log.Println(s.HasSuffix(str, "ro"))
	log.Println(s.Index(str, ","))
	log.Println(s.Join([]string{"1", "2", "3"}, "-"))
	log.Println(s.Repeat(str, 3))
	log.Println(s.Replace(str, "o", "0", -1))
	log.Println(s.Replace(str, "o", "0", 1))
	log.Println(s.Split(str, " "))
	log.Println(s.ToLower(str))
	log.Println(s.ToUpper(str))

	log.Println(len(str))
	log.Println(str[1])
}
