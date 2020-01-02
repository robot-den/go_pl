package main

import (
  "fmt"
  "time"
  "sync"
)

func main() {
  var wg sync.WaitGroup
  names := []string{
    "one",
    "two2",
    "three",
  }

  sizes := make(chan int)

  for _, name := range names {
    wg.Add(1)

    go func (n string) {
      defer wg.Done()

      fmt.Println("gorutine started")
      size := len(n)
      time.Sleep(time.Duration(size) * time.Second)
      sizes <- size
      fmt.Println("gorutine stopped")
    }(name)
  }

  go func () {
    wg.Wait()
    fmt.Println("Going to close sizes channel")
    close(sizes)
  }()

  var total int
  for size := range sizes {
    if size == 0 {
      fmt.Println("zero in sizes")
    }

    total += size
  }

  fmt.Println(total)
}
