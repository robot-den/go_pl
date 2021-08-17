package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	var num uint64
	wg := &sync.WaitGroup{}

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go worker(&num, wg)
	}

	wg.Wait()
	log.Println(num)
}

func worker(num *uint64, wg *sync.WaitGroup) {
	defer wg.Done()

	for c := 0; c < 1000; c++ {
		atomic.AddUint64(num, 1)
	}
}
