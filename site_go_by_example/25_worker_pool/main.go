package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const workersAmount = 3

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	jobs := make(chan int)
	results := make(chan string)
	wg := &sync.WaitGroup{}

	go func(results <-chan string) {
		for result := range results {
			log.Println(result)
		}
	}(results)

	for i := 0; i < workersAmount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, wg)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	close(results)
	log.Println("everything is done")
}

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		log.Printf("worker #%d catch job #%d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		results <- fmt.Sprintf("result for #%d", job)
	}
}

func Kek() {
	log.Println("kek")
}
