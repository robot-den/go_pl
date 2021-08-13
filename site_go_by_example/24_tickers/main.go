package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan struct{})
	go tick(ticker, done)

	time.Sleep(5 * time.Second)
	ticker.Stop()
	close(done)

	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func tick(ticker *time.Ticker, done chan struct{}) {
	defer fmt.Println("Goroutine stopped")

	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		case <-done:
			return
		}
	}
}
