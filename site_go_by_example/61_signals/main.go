package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan int)
	wg := &sync.WaitGroup{}

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go handleSignal(sigs, done)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go waiter(done, wg, i)
	}

	wg.Wait()
	fmt.Println("main: exiting")
}

func waiter(done <-chan int, wg *sync.WaitGroup, n int) {
	defer wg.Done()

	fmt.Printf("goroutine #%d: started\n", n)
	<-done
	fmt.Printf("goroutine #%d: done\n", n)
}

func handleSignal(sigs <-chan os.Signal, done chan int) {
	fmt.Println("signal handler: listen")
	sig := <-sigs
	fmt.Printf("\nsignal handler: receives signal: %v\n", sig)
	close(done)
}
