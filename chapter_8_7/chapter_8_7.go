package main

import (
  "fmt"
  "time"
  "os"
)

func main() {
  abort := make(chan struct{})
  go func () {
    buffer := make([]byte, 1)
    os.Stdin.Read(buffer)
    fmt.Println("Abort command is accepted!")
    abort <- struct{}{}
  }()
  fmt.Println("Countdown just started! Press <Enter> to abort launch.")
  // tryAfter(abort)
  tryTick(abort)
}

func launch() {
  fmt.Println("Boom!")
}

func tryTick(abort <-chan struct{}) {
  // tick := time.Tick(1 * time.Second) // gorutine leaks if we don't use it later in programm
  ticker := time.NewTicker(1 * time.Second)
  defer ticker.Stop()
  for count := 10; count > 0; count-- {
    select {
    case <-ticker.C:
      fmt.Println(count)
    case <-abort:
      fmt.Println("Launch is terminated!")
      return
    }
  }
  launch()
}

func tryAfter(abort <-chan struct{}) {
  after := time.After(3 * time.Second)
  select {
  case <- after:
    fmt.Println("After works")
    return
  case <- abort:
    fmt.Println("Aborted")
    return
  }
}

// NOTE: non-determ. behaviour of select statement!
  // ch := make(chan int, 2)
  //
  // for i := 0; i < 10; i++ {
  //   select {
  //   case x := <-ch:
  //     fmt.Println(x)
  //   case ch <-i:
  //   }
  // }
