package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	printDate()
	callGrep()
	callLs()
}

func callLs() {
	lsCmd := exec.Command("bash", "-c", "ls -lah")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("> ls -lah\n%s\n", string(lsOut))
}

func callGrep() {
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, err := io.ReadAll(grepOut)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(grepBytes))
	grepCmd.Wait()
}

func printDate() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("> date\n%v\n", string(dateOut))
}
