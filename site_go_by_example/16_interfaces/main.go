package main

import "fmt"

type mortal interface {
	die()
}

type human struct {}

func (h *human) die()  {
	fmt.Println("Finally!")
}

type hope struct {}

func (h *hope) die() {
	fmt.Println("Can't believe it...")
}

func kill(m mortal) {
	m.die()
}

func main() {
	kill(&human{})
	kill(&hope{})
}