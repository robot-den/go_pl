package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	timers := make(chan *time.Timer, 10)
	producer(timers)

	n := rand.Intn(10)
	fmt.Printf("will sleep for %d seconds\n", n)
	time.Sleep(time.Duration(n) * time.Second)

	for timer := range timers {
		ok := timer.Stop()
		if ok {
			fmt.Println("Timer closed")
		}
	}
	fmt.Println("Sleep at the end")
	time.Sleep(1 * time.Second)
}

func producer(timers chan<- *time.Timer) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			n := rand.Intn(10)
			timer := time.NewTimer(time.Duration(n) * time.Second)
			timers <- timer
			wg.Done()
			<-timer.C
			fmt.Println("Timer that waited for", n, "seconds")
		}()
	}
	wg.Wait()
	close(timers)
}
