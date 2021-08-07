package main

import (
	"fmt"
	"time"
)

func f(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go f("goroutine")

	go func(name string) {
		for i := 0; i < 3; i++ {
			fmt.Println(name, ":", i)
			time.Sleep(500 * time.Millisecond)
		}
	}("anonymous")

	f("main")

	time.Sleep(time.Second)
}
