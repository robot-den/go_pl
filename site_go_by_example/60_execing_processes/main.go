package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	args := []string{"ls", "-lah"}
	env := os.Environ()

	err = syscall.Exec(binary, args, env)

	fmt.Println("this string won't be printed because syscall replaces current process with ls")
	if err != nil {
		panic(err)
	}
}
