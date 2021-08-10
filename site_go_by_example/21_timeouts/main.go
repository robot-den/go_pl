package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	c := make(chan string, 1)

	go func(c chan<- string) {
		delay := rand.Intn(1000)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		c <- strconv.Itoa(delay)
	}(c)

	select {
	case result := <-c:
		fmt.Println("Result is taken:", result)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Canceled by timeout...")
	}
}
