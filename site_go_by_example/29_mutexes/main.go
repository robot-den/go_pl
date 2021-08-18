package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var state = make(map[int]int)
var mutex = &sync.Mutex{}
var readOps uint64
var writeOps uint64

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg := &sync.WaitGroup{}
	done := make(chan int)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go reader(wg, done)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go writer(wg, done)
	}

	time.Sleep(1 * time.Second)
	close(done)
	wg.Wait()

	printReport()
}

func reader(wg *sync.WaitGroup, done <-chan int) {
	defer wg.Done()

	for {
		select {
		case <-done:
			return
		default:
			key := rand.Intn(10)
			mutex.Lock()
			_ = state[key]
			mutex.Unlock()
			atomic.AddUint64(&readOps, 1)
			time.Sleep(time.Millisecond)
		}
	}
}

func writer(wg *sync.WaitGroup, done <-chan int) {
	defer wg.Done()

	for {
		select {
		case <-done:
			return
		default:
			key := rand.Intn(10)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] += val
			mutex.Unlock()
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
		}
	}
}

func printReport() {
	fmt.Println("reads:", readOps)
	fmt.Println("writes:", writeOps)
	fmt.Println("state:", state)
}
