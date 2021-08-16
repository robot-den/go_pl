package main

import (
	"log"
	"time"
)

func main() {
	Limiter()
	LimiterWithBurst()
}

func Limiter() {
	log.Println("simple limiter")

	requests := make(chan int, 3)
	go requestsGenerator(requests)

	worker(requests, time.Tick(200*time.Millisecond))
}

func LimiterWithBurst() {
	log.Println("bursty limiter")

	requests := make(chan int, 3)
	go requestsGenerator(requests)

	worker(requests, burstyLimiter())
}

func requestsGenerator(requests chan<- int) {
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
}

func worker(requests <-chan int, limiter <-chan time.Time) {
	for request := range requests {
		t := <-limiter
		log.Printf("request #%d processed: %v", request, t.Format("2006-01-02T15:04:05.999"))
	}
}

func burstyLimiter() <-chan time.Time {
	burstyLimiter := make(chan time.Time, 3)

	// warm up limiter
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	return burstyLimiter
}
