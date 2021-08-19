package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var readOps uint64
var writeOps uint64

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg := &sync.WaitGroup{}
	reads := make(chan readOp)
	writes := make(chan writeOp)
	done := make(chan int)

	wgStateKeeper := &sync.WaitGroup{}
	doneStateKeeper := make(chan int)

	wgStateKeeper.Add(1)
	go stateKeeper(reads, writes, doneStateKeeper, wgStateKeeper)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go reader(reads, wg, done)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go writer(writes, wg, done)
	}

	time.Sleep(1 * time.Second)
	close(done)
	wg.Wait()
	close(doneStateKeeper)
	wgStateKeeper.Wait()
}

func stateKeeper(reads <-chan readOp, writes <-chan writeOp, done <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	state := make(map[int]int)
	defer printReport(state)

	for {
		select {
		case read := <-reads:
			read.resp <- state[read.key]
		case write := <-writes:
			state[write.key] = write.value
			write.resp <- true
		case <-done:
			return
		}
	}
}

func reader(reads chan<- readOp, wg *sync.WaitGroup, done <-chan int) {
	defer wg.Done()

	read := readOp{
		key:  0,
		resp: make(chan int),
	}

	for {
		select {
		case <-done:
			return
		default:
			read.key = rand.Intn(10)
			reads <- read
			_ = <-read.resp
			atomic.AddUint64(&readOps, 1)
			time.Sleep(time.Millisecond)
		}
	}
}

func writer(writes chan<- writeOp, wg *sync.WaitGroup, done <-chan int) {
	defer wg.Done()

	write := writeOp{
		key:   0,
		value: 0,
		resp:  make(chan bool),
	}

	for {
		select {
		case <-done:
			return
		default:
			write.key = rand.Intn(10)
			write.value = rand.Intn(100)
			writes <- write
			_ = <-write.resp
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
		}
	}
}

func printReport(state map[int]int) {
	fmt.Println("reads:", readOps)
	fmt.Println("writes:", writeOps)
	fmt.Println("state:", state)
}
