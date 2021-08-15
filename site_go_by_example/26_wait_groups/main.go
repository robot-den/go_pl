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
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, wg)
	}

	wg.Wait()
	fmt.Println("all work is done")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("worker #%d: done\n", id)

	ms := rand.Intn(1000)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
