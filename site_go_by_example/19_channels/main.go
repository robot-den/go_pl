package main

import (
	"fmt"
)

func main() {
	unbuffered()
	buffered()
	synchronization()
}

func unbuffered() {
	messages := make(chan string)

	go func() {
		messages <- "unbuffered"
	}()

	fmt.Println(<-messages)
}

func buffered() {
	messages := make(chan string, 2)

	messages <- "unbuffered 1"
	messages <- "unbuffered 2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func synchronization() {
	stage1 := make(chan string, 2)
	stage2 := make(chan string, 2)
	done := make(chan string)

	go generator(stage1)
	go worker1(stage1, stage2)
	go worker2(stage2, done)

	for n := range done {
		fmt.Println("Processed: ", n)
	}
}

func generator(out chan<- string) {
	samples := []string{"one", "two", "three", "four", "five"}
	for _, sample := range samples {
		out <- sample
	}
	close(out)
}

func worker1(in <-chan string, out chan<- string) {
	for text := range in {
		out <- fmt.Sprintf("text: %s, len: %d", text, len(text))
	}
	close(out)
}

func worker2(in <-chan string, out chan<- string) {
	for text := range in {
		out <- fmt.Sprintf("%s. worker2 to done", text)
	}
	close(out)
}
