package main

import "fmt"

func main() {
	panicWithRecover()
	justPanic("main()")
}

func panicWithRecover() {
	defer catchPanic("panicWithRecover()")

	justPanic("panicWithRecover()")
}

func justPanic(caller string) {
	panic(fmt.Sprintf("just panic from %s", caller))
}

func catchPanic(caller string) {
	if r := recover(); r != nil {
		fmt.Printf("catch panic from %s\n", caller)
	}
}
