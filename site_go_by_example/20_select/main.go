package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func(c chan<- int) {
		for i := 0; i < 3; i++ {
			c <- i
		}
		close(c)
	}(c1)

	go func(c chan<- string) {
		for _, letter := range strings.Split("golang", "") {
			c <- letter
		}
		close(c)
	}(c2)

	for i := 0; i < 6; i++ {
		select {
		case n := <-c1:
			fmt.Println("Channel 1:", n)
		case s := <-c2:
			fmt.Println("Channel 2:", s)
		default:
			fmt.Println("Default")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
