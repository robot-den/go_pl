package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	doneExample()
	takeValuesAfterCloseExample()
}

func doneExample() {
	jobs := make(chan string, 5)
	done := make(chan struct{})

	go producer(jobs)
	go worker(jobs, done)

	fmt.Println("Wait for goroutines")
	<-done
}

func producer(out chan<- string) {
	for i := 0; i < 10; i++ {
		out <- strconv.Itoa(i)
	}
	close(out)
}

func worker(in <-chan string, done chan struct{}) {
	for job := range in {
		fmt.Printf("Process job #%s\n", job)
		time.Sleep(100 * time.Millisecond)
	}
	close(done)
}

func takeValuesAfterCloseExample() {
	c := make(chan string, 2)
	c <- "first"
	c <- "second"
	close(c)

	for str := range c {
		fmt.Println(str)
	}
}
